package main


type GetCustomerByIDService struct {
    getCustomerByIDRepository *GetCustomerByIDRepository
}

func NewGetCustomerByIDService(rdbRepository *GetCustomerByIDRepository) *GetCustomerByIDService {
    return &GetCustomerByIDService{
        getCustomerByIDRepository: rdbRepository,
    }
}
