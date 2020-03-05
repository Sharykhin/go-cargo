package service

import (
	"Sharykhin/go-cargo/domain/company/repository"
	"Sharykhin/go-cargo/domain/company/repository/request"
	"context"
	"fmt"
	"strings"

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
func (h *CompanyServiceHandler) Create(ctx context.Context, req CreateCompanyRequest) (*model.Company, error) {

	if err := h.validate(&req); err != nil {
		return nil, err
	}

	uuid := model.NewUUID()

	agg, err := h.companyRepository.CreateCompany(
		ctx,
		request.CreateCompany{
			string(uuid),
			req.Name,
			req.Country,
			req.State,
			req.City,
			req.Street,
			req.Number,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Could not create company: %v", err)
	}

	company := model.NewCompnay(uuid, agg.Country, agg.State, agg.City, agg.Street, agg.Number)

	return company, nil
}

func (h CompanyServiceHandler) validate(req *CreateCompanyRequest) error {
	if strings.Trim(req.Country, " ") == "" {
		return NewValidationError(FieldEmpty, "Country can not be empty", "Country")
	}

	if strings.Trim(req.City, " ") == "" {
		return NewValidationError(FieldEmpty, "City can not be empty", "City")
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
