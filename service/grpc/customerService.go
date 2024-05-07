package grpc

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/pstano1/customer-api/internal/pkg"
	pb "github.com/pstano1/customer-api/proto"
	"github.com/pstano1/customer-api/service"
	"google.golang.org/grpc"
)

type CustomerService struct {
	api service.IAPI
	pb.UnimplementedCustomerServiceServer
}

func NewService(grpcServer *grpc.Server, api service.IAPI) {
	customersGRPC := &CustomerService{api: api}
	pb.RegisterCustomerServiceServer(grpcServer, customersGRPC)
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *pb.CustomerCreateRequest) (*pb.CustomerCreateResponse, error) {
	var response pb.CustomerCreateResponse
	var customerCreate pkg.CustomerCreate
	customerCreate.Name = req.GetName()
	customerCreate.Tag = req.GetTag()
	customer, err := s.api.CreateCustomer(customerCreate)
	if err != nil {
		return &response, err
	}
	err = copier.Copy(&response, customer)
	if err != nil {
		return &response, err
	}
	return &response, nil
}

func (s *CustomerService) ExchangeTagForId(ctx context.Context, req *pb.ExchangeTagForIdRequest) (*pb.ExchangeTagForIdResponse, error) {
	var response pb.ExchangeTagForIdResponse
	tag := req.GetTag()
	if tag == "" {
		return &response, errors.New("tag is required")
	}
	customerId, err := s.api.ExchangeTagForId(tag)
	if err != nil {
		return &response, err
	}
	if customerId != nil {
		response.Id = *customerId
	}
	return &response, nil
}

func (s *CustomerService) ValidateId(ctx context.Context, req *pb.ValidateIdRequest) (*pb.ValidateIdResponse, error) {
	var response pb.ValidateIdResponse
	id := req.GetId()
	if id == "" {
		return &response, errors.New("id is required")
	}
	ok, err := s.api.ValidateId(id)
	if err != nil {
		return &response, err
	}
	response.Ok = ok
	return &response, nil
}
