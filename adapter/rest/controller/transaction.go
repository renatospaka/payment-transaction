package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/renatospaka/payment-transaction/core/dto"
	"github.com/renatospaka/payment-transaction/core/usecase"
)

type TransactionController struct {
	usecases *usecase.TransactionUsecase
}

func NewTransactionController(usecases *usecase.TransactionUsecase) *TransactionController {
	return &TransactionController{
		usecases: usecases,
	}
}


// Process a new Credit Cart Transaction
func (c *TransactionController) Process(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.process")

	var transaction dto.TransactionCreateDto
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		json.NewEncoder(w).Encode("error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := c.usecases.Create(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}


// Return a specific Transaction
func (c *TransactionController) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.get")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction, err := c.usecases.Find(id)
	if err != nil {
		json.NewEncoder(w).Encode("error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&transaction)
}


// Update basic information about a previously processed Transaction
func (c *TransactionController) Modify(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}


// Delete an existing Transaction
func (c *TransactionController) Remove(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}


// Return all existing Transactions (paginated) 
func (c *TransactionController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
