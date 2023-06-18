package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type GrpcClient struct {
	ctx    context.Context
	Client grpc.ClientConn
}

func NewGrpcClient(ctx context.Context) *GrpcClient {
	log.Println("estabelecendo conexão com o servidor gRPC")

	cli := &GrpcClient{
		ctx: ctx,
	}
	return cli
}

// func (g *GrpcClient) Connect(port string) {
// 	options := make([]grpc.DialOption, 0)
// 	g.Client, err := grpc.Dial(":" + port, options...)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	srv, err := pb.AuthorizationServiceClient(g.Client)
// }
