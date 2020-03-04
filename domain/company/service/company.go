package service

import (
	"Sharykhin/go-cargo/domain/company/repository"
	"context"
	"fmt"

	"Sharykhin/go-cargo/domain/company/model"
)

type (
	// CompanyServiceHandler implements CompanyService interface and handles
	// general business logic
	CompanyServiceHandler struct {
		companyRepository repository.CompanyRepository
	}
)

// Create creates a new company
func (h CompanyServiceHandler) Create(ctx context.Context, r CreateCompanyRequest) (*model.Company, error) {

	if err := h.validate(&r); err != nil {
		return nil, err
	}

	agg, err := h.companyRepository.CreateCompany(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not create company: %v", err)
	}

	company := model.NewCompnay("SOMEUUID", agg.Country, agg.State, agg.City, agg.Street, agg.Number)
	return company, nil
}

func (h CompanyServiceHandler) validate(r *CreateCompanyRequest) error {
	if r.Country == "" {
		return NewValidationError(FieldEmpty, "Country can not be empty", "Country")
	}

	return nil
}

// NewCompanyHandler is a function constructor that returns
// company handler
func NewCompanyHandler(
	companyRepository repository.CompanyRepository,
) *CompanyServiceHandler {
	return &CompanyServiceHandler{
		companyRepository: companyRepository,
	}
}
