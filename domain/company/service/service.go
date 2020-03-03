package service

import (
	"Sharykhin/go-cargo/domain/company/model"
	"context"
)

type (
	CreateCompanyRequest struct {
		Country string
		State string
		City string
		Street string
		Number string
	}

	CompanyService interface {
		Create(ctx context.Context, request CreateCompanyRequest) (*model.Company, error)
	}
)

