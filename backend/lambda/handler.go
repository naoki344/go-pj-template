/*
Package yourpackage does something interesting.
*/
package main

import (
	"os"
	"fmt"
	"sync"
	"context"
	"database/sql"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
	"cdk-lambda-go/ogen"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)


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
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	DBHOST := os.Getenv("DB_HOST")
	DBPORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")
	fmt.Println("before connection open")
	conn, err := sql.Open(
		"mysql", 
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s",
		USER, PASS, DBHOST, DBPORT, DBNAME, url.PathEscape("Asia/Tokyo")),
	)
	defer conn.Close()
	if err != nil {
		fmt.Println("Fail to connect db" + err.Error())
	}
	// 接続確認
	err = conn.Ping()
	if err != nil {
		fmt.Println("Failed to connect rds " + err.Error())
	} else {
		fmt.Println("Success to connect rds")
	}

	// TODO: DIツール,ormapper,db-migrator

	// DBからレコードを抽出
	rows, err := conn.Query("select id, title, content from notes;")
	if err != nil {
		fmt.Println("Fail to query from db " + err.Error())
	}
	// データを構造体へ変換
	var Notes []ogen.Note
	for rows.Next() {
		var tmpNote ogen.Note
		err := rows.Scan(&tmpNote.ID, &tmpNote.Title, &tmpNote.Content)
		if err != nil {
			fmt.Println("Fail to scan records " + err.Error())
		}
		Notes = append(Notes, tmpNote)
	}
	return &Notes[0], nil
}

func main() {
	service := &noteService{
		notes: map[int64]ogen.Note{},
	}
	s, _ := ogen.NewServer(service)
	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16
	lambda.Start(httpadapter.New(s).ProxyWithContext)
}
