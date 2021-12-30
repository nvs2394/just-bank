package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (accountHandlers *AccountHandler) CreateAccount(response http.ResponseWriter, request *http.Request) {
	var accountRequest dto.NewAccountRequest
	customerId := mux.Vars(request)["customer_id"]

	if err := json.NewDecoder(request.Body).Decode(&accountRequest); err != nil {
		writeResponse(response, http.StatusBadRequest, err.Error())
	} else {
		accountRequest.CustomerId = customerId
		account, appErr := accountHandlers.service.NewAccount(accountRequest)
		if appErr != nil {
			writeResponse(response, appErr.Code, appErr.Message)
		} else {
			writeResponse(response, http.StatusCreated, account)
		}
	}

}
