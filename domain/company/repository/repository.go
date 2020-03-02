package repository

import (
	"Sharykhin/go-cargo/domain/company/repository/aggregate"
	"context"
)

type (
	CompanyRepository interface {
		CreateCompany(ctx context.Context) (*aggregate.CompanyAggregate, error)
	}
)