package usecase

import (
	"context"
	"log"
	// "math/rand"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionUsecase) authorizeTransaction(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	log.Println("usecase.transactions.authorize")

	// // Issue #21 open to solve this lack of communication to the server.
	// // Meanwhile, this is an workaround to respond to the caller
	// min, max, status := 0, 100, "pending"

	// random := rand.Intn(max-min) + min
	// if random <= 30 {
	// 	status = "denied"
	// } else {
	// 	status = "approved"
	// }

	var response *pb.AuthorizationResponse

	// Execute the gRPC call
	auth, err := t.services.AuthorizeTransaction(ctx, in)
	if err != nil {
		response = &pb.AuthorizationResponse{
			AuthorizationId: "",
			ClientId:        in.ClientId,
			TransactionId:   in.TransactionId,
			Status:          auth.Status,
			Value:           in.Value,
			ErrorMessage:    auth.ErrorMessage,
		}
		return response, err
	}

	response = &pb.AuthorizationResponse{
		AuthorizationId: auth.AuthorizationId,
		ClientId:        in.ClientId,
		TransactionId:   in.TransactionId,
		Status:          auth.Status,
		Value:           in.Value,
	}
	return response, nil
}
