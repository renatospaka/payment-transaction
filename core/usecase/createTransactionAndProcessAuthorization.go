package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/entity"
	"github.com/renatospaka/payment-transaction/core/service"
)

func (t *TransactionUsecase) createTransactionAndProcessAuthorization(ctx context.Context, tr *dto.TransactionCreateDto) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.createTransactionAndProcessAuthorization")

	if t.services == nil {
		return nil, service.ErrGRPCNotDefined
	}

	transaction, err := entity.NewTransaction(tr.ClientID, tr.Value)
	if err != nil {
		return nil, err
	}
	transactionId := transaction.GetID()
	clientId := transaction.GetClientID()
	value := transaction.GetValue()
	tr.ID = transactionId

	err = t.repo.Create(transaction)
	if err != nil {
		return nil, err
	}

	// Call the Authotization service and update status accordingly
	auth := &pb.AuthorizationProcessNewRequest{
		TransactionId: transactionId,
		ClientId:      clientId,
		Value:         value,
	}
	authorize, err := t.authorizeNewTransaction(ctx, auth)
	if err != nil {
		return nil, err
	}

	// request the gRPC server to authorize the transaction
	tran := &dto.TransactionAuthorizeDto{
		ID:              transactionId,
		AuthorizationId: authorize.AuthorizationId,
		ClientID:        clientId,
		Value:           value,
	}
	if authorize.Status == entity.TR_APPROVED {
		err = t.approveTransaction(ctx, tran)
	} else if authorize.Status == entity.TR_DENIED {
		err = t.denyTransaction(ctx, tran)
	}
	if err != nil {
		return nil, err
	}

	// Find the new Transaction the reply to the caller
	created, err := t.findTransactionById(ctx, transactionId)
	if err != nil {
		return nil, err
	}
	return created, nil
}
