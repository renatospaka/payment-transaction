package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionUsecase) authorizeTransaction(ctx context.Context, n *pb.AuthorizationResponse) (*pb.AuthorizationResponse, error) {
	log.Println("usecase.transactions.authorizeTransaction")

	return nil, nil
}
