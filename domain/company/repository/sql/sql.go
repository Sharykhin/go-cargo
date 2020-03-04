package sql

import (
	"Sharykhin/go-cargo/domain/company/repository/aggregate"
	"context"
	"database/sql"
	"fmt"
)

type (
	CompanyRepository struct {
		db *sql.DB
	}
)

func (r CompanyRepository) CreateCompany(ctx context.Context) (*aggregate.CompanyAggregate, error) {

	tx, ok := ctx.Value("sql-transaction-tx").(*sql.Tx)
	if ok {
		fmt.Println("use transacton", tx)
	}

	return nil, nil
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	repo := CompanyRepository{
		db: db,
	}

	return &repo
}
