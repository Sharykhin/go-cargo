package sql

import (
	"Sharykhin/go-cargo/domain/company"
	"Sharykhin/go-cargo/domain/company/repository/aggregate"
	"Sharykhin/go-cargo/domain/company/repository/request"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type (
	// CompanyRepository is a concrete implementation of company repository interface
	CompanyRepository struct {
		db *sql.DB
	}
)

// CreateCompany inserts a new row into companies table
func (r CompanyRepository) CreateCompany(ctx context.Context, req request.CreateCompany) (*aggregate.CompanyAggregate, error) {

	tx, ok := ctx.Value(company.SQLTransactionTX).(*sql.Tx)
	if ok {
		fmt.Println("use transacton", tx)
	}

	var compnayID uint64
	err := r.db.QueryRow(
		"INSERT INTO companies(uuid, name, created_at) values($1, $2, $3) RETURNING id",
		req.UUID,
		req.Name,
		time.Now().UTC(),
	).Scan(&compnayID)

	if err != nil {
		return nil, fmt.Errorf("failed to insert a new row into companies table")
	}
	agg := aggregate.CompanyAggregate{
		ID:      compnayID,
		UUID:    req.UUID,
		Country: req.Country,
		State:   req.State,
		City:    req.City,
		Street:  req.Street,
		Number:  req.Number,
	}

	return &agg, nil
}

// NewCompanyRepository is a function constructor that returns a new instance of a company repository implementation
func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	repo := CompanyRepository{
		db: db,
	}

	return &repo
}
