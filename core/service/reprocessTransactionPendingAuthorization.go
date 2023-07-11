package service

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
)

func (t *TransactionService) reprocessTransactionPendingAuthorization(ctx context.Context, in *pb.AuthorizationReprocessPendingRequest) (*pb.AuthorizationReprocessPendingResponse, error) {
	log.Println("service.transactions.reprocessTransactionPendingAuthorization")

	response, err := t.grpServices.Client.ReprocessPendingAuthorization(ctx, in)
	reprocessResp := &pb.AuthorizationReprocessPendingResponse{
		AuthorizationId: response.AuthorizationId,
		ClientId:        response.ClientId,
		TransactionId:   response.TransactionId,
		Status:          response.Status,
		Value:           response.Value,
		ErrorMessage:    "",
	}

	if err != nil {
		reprocessResp.ErrorMessage = response.ErrorMessage
		return reprocessResp, err
	}
	return reprocessResp, nil
}
