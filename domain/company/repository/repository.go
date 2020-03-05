package repository

import (
	"Sharykhin/go-cargo/domain/company/repository/aggregate"
	"Sharykhin/go-cargo/domain/company/repository/request"
	"context"
)

type (
	// CompanyRepository describes storage interface
	CompanyRepository interface {
		CreateCompany(ctx context.Context, req request.CreateCompany) (*aggregate.CompanyAggregate, error)
	}
)
