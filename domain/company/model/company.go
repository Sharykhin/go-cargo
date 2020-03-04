package model

import (
	"log"

	uuid "github.com/nu7hatch/gouuid"
)

type (
	// UUID would rather to be used across the application and should be used as general identificator
	// for a company
	UUID string

	// Address is a value object that is used as attribute of Company model
	Address struct {
		Country string
		State   string
		City    string
		Street  string
		Number  string
	}
	// Company is a represenation of a root domain model
	Company struct {
		UUID    UUID
		Address Address
	}
)

// NewCompnay is factory for company domain model
func NewCompnay(UUID UUID, country, state, city, street, number string) *Company {
	address := Address{
		Country: country,
		State:   state,
		City:    city,
		Street:  street,
		Number:  number,
	}

	company := Company{
		UUID:    UUID,
		Address: address,
	}

	return &company
}

// NewUUID generates a new UUID of version 4
func NewUUID() UUID {
	u, err := uuid.NewV4()
	if err != nil {
		log.Panicf("failed to generate uuid v4: %v", err)
	}

	return UUID(u.String())
}
