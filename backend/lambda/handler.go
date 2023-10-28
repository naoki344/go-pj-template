/*
Package yourpackage does something interesting.
*/
package main

import (
	"cdk-lambda-go/ogen"
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"net/url"
	"os"
	"sync"
)

type noteService struct {
	notes map[int64]ogen.Note
	mux   sync.Mutex
}

type Note struct {
	ID      int64  `bun:"id"`
	Title   string `bun:"title"`
	Content string `bun:"content"`
}

func (n *noteService) CreateNote(_ context.Context, req *ogen.Note) (*ogen.Note, error) {
	n.mux.Lock()
	defer n.mux.Unlock()

	n.notes[req.ID] = *req
	return req, nil
}

func (n *noteService) GetNoteByID(_ context.Context, params ogen.GetNoteByIDParams) (ogen.GetNoteByIDRes, error) {
	// TODO: DIツール,db-migrator
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
	db := bun.NewDB(conn, mysqldialect.New())

	note_db := Note{}
	get_err := db.NewSelect().Model(&note_db).Where("id = 1").Scan(context.Background())
	if get_err != nil {
		panic(get_err)
	}
	Note := ogen.Note{
		ID:      note_db.ID,
		Title:   note_db.Title,
		Content: note_db.Content,
	}

	return &Note, nil
}

func main() {
	service := &noteService{
		notes: map[int64]ogen.Note{},
	}
	s, _ := ogen.NewServer(service)
	// NOTE: https://github.com/awslabs/aws-lambda-go-api-proxy/blob/master/httpadapter/adapter.go#L16
	lambda.Start(httpadapter.New(s).ProxyWithContext)
}
