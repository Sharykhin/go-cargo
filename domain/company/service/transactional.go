package service

import (
	"Sharykhin/go-cargo/domain/company/model"
	"context"
	"database/sql"
)

type (
	//ContextKey is an alias of string that wiil be used with context.WithValue to prevent potential collisions
	ContextKey string
	// SQLTransactionalDecorator is a decorator aroung CompanyService interface
	// that wraps functions in sql transaction
	SQLTransactionalDecorator struct {
		service CompanyService
		db      *sql.DB
	}
)

var (
	//TODO: since context keys can be used at leaset on repository level, you should move it out of this package.

	// SQLTransactionTX it's a context key regarding sql transaction struct
	SQLTransactionTX = ContextKey("sql-transaction-tx")
)

// Create wraps excution code into sql transaction
func (d SQLTransactionalDecorator) Create(ctx context.Context, request CreateCompanyRequest) (c *model.Company, err error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	ctx = context.WithValue(ctx, SQLTransactionTX, tx)
	c, err = d.service.Create(ctx, request)

	return c, err
}

// NewSQLTransactionalDecorator is a function constructor
// that returns a new instance of SQLTransactionalDecorator
func NewSQLTransactionalDecorator(
	service CompanyService,
	db *sql.DB,
) *SQLTransactionalDecorator {
	return &SQLTransactionalDecorator{
		service: service,
		db:      db,
	}
}
