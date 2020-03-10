package company

import (
	"Sharykhin/go-cargo/domain/company/repository/sql"
	"Sharykhin/go-cargo/domain/company/service"
	"Sharykhin/go-cargo/infrastructure/database/postgres"
	"errors"
	"fmt"
	"net"
	"os"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"context"

	"google.golang.org/grpc"
)

type (
	server struct {
		handler service.CompanyService
	}
)

func handleError(err error) error {
	var vErr service.ValidationError
	if errors.As(err, &vErr) {
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
	fmt.Println("Country", req.GetCountry())
	company, err := s.handler.Create(ctx, service.CreateCompanyRequest{
		Name:    req.GetName(),
		Country: req.GetCountry(),
		State:   req.GetState(),
		City:    req.GetCity(),
		Street:  req.GetStreet(),
		Zip:     req.GetZip(),
		Number:  req.GetNumber(),
	})

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
	db := postgres.NewConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	RegisterCompanyServiceServer(s, &server{
		handler: service.NewSQLTransactionalDecorator(
			service.NewCompanyHandler(
				sql.NewCompanyRepository(db),
			),
			db,
		),
	})
	reflection.Register(s)

	return s.Serve(lis)
}
