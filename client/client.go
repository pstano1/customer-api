package client

import (
	"context"

	pb "github.com/pstano1/customer-api/client/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ICustomerServiceClient interface {
	CreateCustomer(request *pb.CustomerCreateRequest) (*pb.CustomerCreateResponse, error)
	ExchangeTagForId(tag string) (*pb.ExchangeTagForIdResponse, error)
	ValidateId(id string) (*pb.ValidateIdResponse, error)
}

type CustomerServiceClient struct {
	client pb.CustomerServiceClient
	logger *zap.Logger
}

func NewCustomerService(conn *grpc.ClientConn, logger *zap.Logger) ICustomerServiceClient {
	c := pb.NewCustomerServiceClient(conn)

	return &CustomerServiceClient{
		client: c,
		logger: logger,
	}
}

func (c *CustomerServiceClient) CreateCustomer(request *pb.CustomerCreateRequest) (*pb.CustomerCreateResponse, error) {
	return c.client.CreateCustomer(context.Background(), request)
}

func (c *CustomerServiceClient) ExchangeTagForId(tag string) (*pb.ExchangeTagForIdResponse, error) {
	return c.client.ExchangeTagForId(context.Background(), &pb.ExchangeTagForIdRequest{Tag: tag})
}

func (c *CustomerServiceClient) ValidateId(id string) (*pb.ValidateIdResponse, error) {
	return c.client.ValidateId(context.Background(), &pb.ValidateIdRequest{Id: id})
}
