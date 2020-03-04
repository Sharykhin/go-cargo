package sql

import (
	"Sharykhin/go-cargo/domain/company/repository/aggregate"
	"context"
	"database/sql"
	"fmt"
	"time"
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

	_, err := tx.Exec("INSERT INTO companies(uuid, name, created_at) values($1, $2, $3)", "123", "name", time.UTC().Now())
	if err != nil {
		return fmt.Errorf("failed to insert a new row into companies table")
	}

	return nil, nil
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	repo := CompanyRepository{
		db: db,
	}

	return &repo
}
