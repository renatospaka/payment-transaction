package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/renatospaka/payment-transaction/core/dto"
)

// Process a new Credit Cart Transaction
func (c *TransactionController) ReprocessPending(w http.ResponseWriter, r *http.Request) {
	log.Println("http.transactions.reprocessPending")

	transactioId := chi.URLParam(r, "transactioId")
	if transactioId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tr dto.TransactionCreateDto
	if err := json.NewDecoder(r.Body).Decode(&tr); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}
	tr.ID = transactioId

	transaction, err := c.usecases.ReprocessTransactionPendingAuthorization(&tr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&transaction)
}
