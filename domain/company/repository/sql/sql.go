package sql

import (
	"Sharykhin/go-cargo/domain/company/repository/aggregate"
	"context"
	"database/sql"
)

type (
	CompanyRepository struct {
		db *sql.DB
	}
)

func (r CompanyRepository) CreateCompany(ctx context.Context) (*aggregate.CompanyAggregate, error) {
	return nil, nil
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	repo := CompanyRepository{
		db: db,
	}

	return &repo
}
