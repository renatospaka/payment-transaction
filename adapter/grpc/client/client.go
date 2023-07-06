package client

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/service"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	ctx      context.Context
	conn     *grpc.ClientConn
	services service.AuthorizationServiceInterface
	client   pb.AuthorizationServiceClient
}

func NewGrpcClient(ctx context.Context, conn *grpc.ClientConn, services service.AuthorizationServiceInterface) *GrpcClient {
	log.Println("estabelecendo conexão com o servidor gRPC")
	cli := &GrpcClient{
		ctx:      ctx,
		conn:     conn,
		services: services,
	}

	cli.client = pb.NewAuthorizationServiceClient(cli.conn)
	return cli
}
