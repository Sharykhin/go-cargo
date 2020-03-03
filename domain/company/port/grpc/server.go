package company

import (
	"Sharykhin/go-cargo/domain/company/service"
	"errors"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"context"
	"google.golang.org/grpc"
)

type (
	server struct {
		handler service.CompanyService
	}
)

func handleError(err error) error {
	if errors.Is(err, service.ValidationError{}) {
		var vErr service.ValidationError
		errors.As(err, &vErr)
		st := status.New(codes.FailedPrecondition, "Validation Failed")

		v := &errdetails.BadRequest_FieldViolation{
			Field:       vErr.Field,
			Description: vErr.Message,
		}

		br := &errdetails.BadRequest{}
		br.FieldViolations = append(br.FieldViolations, v)

		st, _ = st.WithDetails(br)

		return st.Err()
	}

	return err
}

// CreateCompany create a new company in a system
func (s *server) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	company, err := s.handler.Create(ctx, service.CreateCompanyRequest{})

	if err != nil {
		err = handleError(err)

		return nil, err
	}

	return &CreateCompanyResponse{
		Uuid: string(company.UUID),
	}, nil
}

// ListenAndServe starts GRPC server and listens incoming requests
func ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to open tcp connection on %s: %v", addr, err)
	}

	s := grpc.NewServer()

	RegisterCompanyServer(s, &server{
		handler: service.NewCompanyHandler(),
	})

	return s.Serve(lis)
}
