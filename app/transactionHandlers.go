package app

import (
	"encoding/json"
	"net/http"

	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (transactionHandlers *TransactionHandler) CreateTransaction(response http.ResponseWriter, request *http.Request) {
	var transactionRequest dto.NewTransactionRequest

	if err := json.NewDecoder(request.Body).Decode(&transactionRequest); err != nil {
		writeResponse(response, http.StatusBadRequest, err.Error())
	} else {
		account, appErr := transactionHandlers.service.NewTransaction(transactionRequest)
		if appErr != nil {
			writeResponse(response, appErr.Code, appErr.Message)
		} else {
			writeResponse(response, http.StatusCreated, account)
		}
	}

}
