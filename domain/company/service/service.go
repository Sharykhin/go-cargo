package service

import (
	"Sharykhin/go-cargo/domain/company/model"
	"context"
)

type (
	// CreateCompanyRequest represents an income request to create a new company
	CreateCompanyRequest struct {
		Name    string
		Country string
		State   string
		City    string
		Street  string
		Zip     string
		Number  string
	}
	// CompanyService describes a general interface around company model
	CompanyService interface {
		Create(ctx context.Context, request CreateCompanyRequest) (*model.Company, error)
	}
)
