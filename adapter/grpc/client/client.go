package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

type GrpcClient struct {
	ctx      context.Context
	conn     *grpc.ClientConn
	Client   pb.AuthorizationServiceClient
}

func NewGrpcClient(ctx context.Context, conn *grpc.ClientConn) *GrpcClient {
	log.Println("estabelecendo conexão com o servidor gRPC")
	cli := &GrpcClient{
		ctx:      ctx,
		conn:     conn,
	}

	cli.Client = pb.NewAuthorizationServiceClient(cli.conn)
	return cli
}
