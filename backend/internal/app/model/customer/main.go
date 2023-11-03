package customermodel

type CustomerID int64

type Customer struct {
	ID      CustomerID
	Title   string
	Content string
}
