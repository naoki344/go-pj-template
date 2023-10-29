// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateNote implements createNote operation.
	//
	// メモを作成する.
	//
	// POST /notes
	CreateNote(ctx context.Context, req *Note) (*Note, error)
	// GetNoteByID implements getNoteByID operation.
	//
	// メモを取得する.
	//
	// GET /notes/{noteID}
	GetNoteByID(ctx context.Context, params GetNoteByIDParams) (GetNoteByIDRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
