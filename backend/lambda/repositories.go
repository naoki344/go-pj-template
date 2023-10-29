package main


type GetCustomerByIDRepository struct {
    db SqlDb
}

func NewGetCustomerByIDRepository(db SqlDb) *GetCustomerByIDRepository {
    return &GetCustomerByIDRepository{
        db: db,
    }
}
