// Code generated by ogen, DO NOT EDIT.

package ogen

type GetNoteByIDNotFound struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *GetNoteByIDNotFound) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *GetNoteByIDNotFound) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *GetNoteByIDNotFound) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *GetNoteByIDNotFound) SetMessage(val string) {
	s.Message = val
}

func (*GetNoteByIDNotFound) getNoteByIDRes() {}

// Ref: #/components/schemas/Note
type Note struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GetID returns the value of ID.
func (s *Note) GetID() int64 {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Note) GetTitle() string {
	return s.Title
}

// GetContent returns the value of Content.
func (s *Note) GetContent() string {
	return s.Content
}

// SetID sets the value of ID.
func (s *Note) SetID(val int64) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Note) SetTitle(val string) {
	s.Title = val
}

// SetContent sets the value of Content.
func (s *Note) SetContent(val string) {
	s.Content = val
}

func (*Note) getNoteByIDRes() {}
