package customermodel

type ID int64

type Name string

type NameKana *string

type Telephone string

type Email string

type PersonInChargeName string

type PersonInChargeNameKana *string

type PostalCode string

type PrefID int

type Address1 string

type Address2 string

type Address struct {
	PostalCode PostalCode
	PrefID     PrefID
	Address1   Address1
	Address2   Address2
}

type Customer struct {
	ID                     ID
	Name                   Name
	NameKana               NameKana
	Telephone              Telephone
	Email                  Email
	PersonInChargeName     PersonInChargeName
	PersonInChargeNameKana PersonInChargeNameKana
	Address                Address
}

type SearchConditions struct{}
