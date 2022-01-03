package service

import (
	"time"

	"github.com/nvs2394/just-bank/domain"
	"github.com/nvs2394/just-bank/dto"
	"github.com/nvs2394/just-bank/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (accountService DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	data := domain.Account{
		Id:          "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Status:      "1",
		Amount:      req.Amount,
	}

	newAccount, err := accountService.repo.Save(data)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func (accountService DefaultAccountService) MakeTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := accountService.repo.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdrawal(req.Amount) {
			return nil, errs.NewBadRequestError("Insufficient balance in the account")
		}
	}

	data := domain.Transaction{
		AccountId:       req.AccountId,
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, appErr := accountService.repo.SaveTransaction(data)
	if appErr != nil {
		return nil, appErr
	}

	response := newTransaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
