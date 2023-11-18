package apiport

type CustomerAPIPortInterface interface {
	GetByID(customerID CustomerID) (*Customer, error)
	UpdateByID(customer *Customer) (*Customer, error)
	CreateCustomer(customer *Customer) (*Customer, error)
	SearchCustomer(pageNumber int64, pageSize int64, conditions *SearchConditions) (*CustomerSearchResult, error)
}
