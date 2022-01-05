package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		TransactionType: "non_type",
		Amount:          2000,
	}
	// Act

	appError := request.Validate()

	//Assert
	if appError.Message != "Transaction type should be withdrawal or deposit" {
		t.Error("Invalid message while testing transaction type")
	}

	if appError.Code != http.StatusBadRequest {
		t.Error("Invalid code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// Arrange

	request := NewTransactionRequest{
		TransactionType: "withdrawal",
		Amount:          -200,
	}

	// Act

	appError := request.Validate()

	// Assert

	if appError.Message != "Amount cannot be less than zero" {
		t.Error("Invalid message while validating amount")
	}

	if appError.Code != http.StatusBadRequest {
		t.Error("Invalid code while validating amount")
	}
}
