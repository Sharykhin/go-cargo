package service

import (
	"Sharykhin/go-cargo/domain/company/model"
	"context"
	"database/sql"
)

type (
	ContextKey                string
	SQLTransactionalDecorator struct {
		service CompanyService
		db      *sql.DB
	}
)

var (
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
	ctx = context.WithValue(ctx, "sql-transaction-tx", tx)
	c, err = d.service.Create(ctx, request)

	return c, err
}

func NewSQLTransactionalDecorator(
	service CompanyService,
	db *sql.DB,
) *SQLTransactionalDecorator {
	return &SQLTransactionalDecorator{
		service: service,
		db:      db,
	}
}
