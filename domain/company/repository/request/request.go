package request

type (
	// CreateCompany represents income request to create a new company in repository
	CreateCompany struct {
		UUID    string
		Name    string
		Country string
		State   string
		City    string
		Street  string
		Zip     string
		Number  string
	}
)
