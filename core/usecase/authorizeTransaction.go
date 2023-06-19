package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionUsecase) authorizeTransaction(ctx context.Context, n *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	log.Println("usecase.transactions.authorize")

	return nil, nil
}
