package model

type (
	CompanyID uint64

	Address struct {
		Country string
		State string
		City string
		Street string
		Number string
	}

	Company struct {
		ID CompanyID
		Address Address
	}
)