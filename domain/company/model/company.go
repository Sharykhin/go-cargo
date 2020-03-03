package model

type (
	// UUID would rather to be used across the application and should be used as general identificator
	// for a company
	UUID string

	Address struct {
		Country string
		State   string
		City    string
		Street  string
		Number  string
	}

	Company struct {
		UUID    UUID
		Address Address
	}
)
