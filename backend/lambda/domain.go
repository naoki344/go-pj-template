package main


type GetNoteByIDService struct {
    getNoteByIDRDBRepository *GetNoteByIDRepository
}

func NewGetNoteByIDService(rdbRepository *GetNoteByIDRepository) *GetNoteByIDService {
    return &GetNoteByIDService{
        getNoteByIDRDBRepository: rdbRepository,
    }
}
