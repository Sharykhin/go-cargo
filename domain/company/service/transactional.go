package service

import (
	"Sharykhin/go-cargo/domain/company"
	"Sharykhin/go-cargo/domain/company/model"
	"context"
	"database/sql"
)

type (
	// SQLTransactionalDecorator is a decorator aroung CompanyService interface
	// that wraps functions in sql transaction
	SQLTransactionalDecorator struct {
		service CompanyService
		db      *sql.DB
	}
)

// Create wraps excution code into sql transaction
func (d SQLTransactionalDecorator) Create(ctx context.Context, req CreateCompanyRequest) (c *model.Company, err error) {
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
	ctx = context.WithValue(ctx, company.SQLTransactionTX, tx)
	c, err = d.service.Create(ctx, req)

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
