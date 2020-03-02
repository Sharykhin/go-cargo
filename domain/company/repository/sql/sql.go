package sql

import "database/sql"

type (
	CompanyRepository struct {
		db *sql.DB
	}
)

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	repo := CompanyRepository{
		db: db,
	}

	return &repo
}
