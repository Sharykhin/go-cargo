package company

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"context"

)

type (
	server struct {

	}
)

func (s *server) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	return nil, nil
}

func ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to open tcp connection on %s: %v", addr, err)
	}

	s := grpc.NewServer()
	RegisterCompanyServer(s, &server{})

	return s.Serve(lis)
}
