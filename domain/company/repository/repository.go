package repository

import (
	"Sharykhin/go-cargo/domain/company/repository/aggregate"
	"context"
)

type (
	// CompanyRepository describes storage interface
	CompanyRepository interface {
		CreateCompany(ctx context.Context) (*aggregate.CompanyAggregate, error)
	}
)
