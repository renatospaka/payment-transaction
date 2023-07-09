package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/renatospaka/payment-transaction/core/dto"
)

// Process a new Credit Cart Transaction
func (c *TransactionController) Process(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.process")

	var tr dto.TransactionCreateDto
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	transaction, err := c.usecases.CreateTransactionAndProcessAuthorization(&tr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&transaction)
}
