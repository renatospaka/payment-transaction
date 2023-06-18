package usecase

import (
	"context"
	"log"

	"github.com/renatospaka/payment-transaction/adapter/grpc/pb"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/repository"
)

type TransactionUsecase struct {
	repo repository.TransactionInterface
}

func NewTransactionUsecase(repo repository.TransactionInterface) *TransactionUsecase {
	return &TransactionUsecase{
		repo: repo,
	}
}

// All business validations related to creating new transactions occur at the usecase level
func (t *TransactionUsecase) Create(tr *dto.TransactionCreateDto) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.create")

	ctx := context.Background()

	// Create the Transaction
	err := t.createTransaction(ctx, tr)
	if err != nil {
		return nil, err
	}

	// // Call the Authotization service and update status accordingly
	// auth, err := authorization.Process()
	// if err != nil {
	// 	return nil, err
	// }
	// authorization_id := authorization.authorization_id

	// if auth.Status == entity.TR_APPROVED {
	// 	err = t.approveTransaction(ctx, tr, authorization_id)
	// } else if auth.Status == entity.TR_DENIED {
	// 	err = t.denyTransaction(ctx, tr, authorization_id)
	// }

	// Find the new Transaction the reply to the caller
	transaction, err := t.findTransaction(ctx, tr.ID)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// All business validations related to finding transactions occur at the usecase level
func (t *TransactionUsecase) Find(id string) (*dto.TransactionDto, error) {
	log.Println("usecase.transactions.find")
	return t.findTransaction(context.Background(), id)
}

// There is no business validations related to retrieving all transactions at the usecase level
func (t *TransactionUsecase) FindAll(page int, limit int) (*dto.TransactionFindAllResponseDto, error) {
	return t.findAllTransactions(context.Background(), page, limit)
}

// All business validations related to deleting transactions occur at the usecase level
func (t *TransactionUsecase) Delete(id string) error {
	log.Println("usecase.transactions.delete")
	return t.deleteTransaction(context.Background(), id)
}

// All business validations related to updating transactions occur at the usecase level
func (t *TransactionUsecase) Update(tr *dto.TransactionUpdateDto) error {
	return t.updateTransaction(context.Background(), tr)
}


// Request to the gRPC Server to process the authorization
func (t *TransactionUsecase) Authorize(auth *pb.AuthorizationResponse) (*pb.AuthorizationResponse, error) {
	return t.authorizeTransaction(context.Background(), auth)
}
