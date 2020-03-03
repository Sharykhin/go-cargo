package service

import (
	"Sharykhin/go-cargo/domain/company/repository"
	"context"

	"Sharykhin/go-cargo/domain/company/model"
)

type (
	// CompanyServiceHandler implements CompanyService interface and handles
	// general business logic
	CompanyServiceHandler struct {
		companyRepository repository.CompanyRepository
	}
)

func (h CompanyServiceHandler) Create(ctx context.Context, r CreateCompanyRequest) (*model.Company, error) {

	if err := h.validate(&r); err != nil {
		return nil, err
	}

	return &model.Company{}, nil
}

func (h CompanyServiceHandler) validate(r *CreateCompanyRequest) error {
	if r.Country == "" {
		return NewValidationError(Empty, "Country can not be empty", "Country")
	}

	return nil
}

func NewCompanyHandler(
	companyRepository repository.CompanyRepository,
) *CompanyServiceHandler {
	return &CompanyServiceHandler{
		companyRepository: companyRepository,
	}
}
