/*
Package yourpackage does something interesting.
*/
package main

import (
	"cdk-lambda-go/ogen"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"sync"
)


type Note struct {
	ID      int64  `bun:"id"`
	Title   string `bun:"title"`
	Content string `bun:"content"`
}


type noteService struct {
	notes map[int64]ogen.Note
	dbRepository *GetNoteByIDRepository
	mux   sync.Mutex
}

func (n *noteService) CreateNote(_ context.Context, req *ogen.Note) (*ogen.Note, error) {
	n.mux.Lock()
	defer n.mux.Unlock()
	return req, nil
}

func (n *noteService) GetNoteByID(_ context.Context, params ogen.GetNoteByIDParams) (ogen.GetNoteByIDRes, error) {
	db := bun.NewDB(n.dbRepository.db.conn, mysqldialect.New())

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


func NewNoteService(dbRepository *GetNoteByIDRepository) *noteService {
	return &noteService{
		dbRepository: dbRepository,
	}
}
