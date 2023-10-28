/*
Package yourpackage does something interesting.
*/
package main

import (
	"log"
	"fmt"
	"sync"
	"context"
	"net/http"
	"cdk-lambda-go/ogen"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)


func handlerWithHttp(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len) // Content-Length と同じサイズの byte 配列を用意
	r.Body.Read(body)         // byte 配列にリクエストボディを読み込む
	log.Printf("request: %s", string(body))
	fmt.Fprintln(w, string(body))
}

type noteService struct {
	notes map[int64]ogen.Note
	mux   sync.Mutex
}

func (n *noteService) CreateNote(_ context.Context, req *ogen.Note) (*ogen.Note, error) {
	n.mux.Lock()
	defer n.mux.Unlock()

	n.notes[req.ID] = *req
	return req, nil
}

func (n *noteService) GetNoteByID(_ context.Context, params ogen.GetNoteByIDParams) (ogen.GetNoteByIDRes, error) {
	n.mux.Lock()
	defer n.mux.Unlock()

	note, ok := n.notes[params.NoteID]
	if !ok {
		return &ogen.GetNoteByIDNotFound{
			Code:    1,
			Message: "メモが見つかりません",
		}, nil
	}
	return &note, nil
}

func main() {
	service := &noteService{
		notes: map[int64]ogen.Note{},
	}
	s, _ := ogen.NewServer(service)

	http.HandleFunc("/notes", handlerWithHttp)
	lambda.Start(httpadapter.New(s).ProxyWithContext)
}
