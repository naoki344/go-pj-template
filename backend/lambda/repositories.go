package main


type GetNoteByIDRepository struct {
    db SqlDb
}

func NewGetNoteByIDRepository(db SqlDb) *GetNoteByIDRepository {
    return &GetNoteByIDRepository{
        db: db,
    }
}
