package rdbadapter


type RdbInterface interface {
	GetCustomerByID(customerID int64) (Customer, error)
}
