package service

import (
	"context"

	"Sharykhin/go-cargo/domain/company/model"
)

type (
	CompanyServiceHandler struct {
	}
)

func (h CompanyServiceHandler) Create(ctx context.Context, r CreateCompanyRequest) (*model.Company, error) {
	if r.Country == "" {
		return nil, NewValidationError(Empty, "Country can not be empty", "Country")
	}

	return nil, nil
}

func NewCompanyHandler() *CompanyServiceHandler {
	return &CompanyServiceHandler{}
}
