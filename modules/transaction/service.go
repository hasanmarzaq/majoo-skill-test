package transaction

import (
	"github.com/google/uuid"
)

type Service interface {
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository Repository
	// outletrepository outlet.Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}

	merchant, err := s.repository.FindMerchantOutlet(input.User.ID, input.OutletID)
	if err != nil {
		return transaction, err
	}
	transaction.MerchantID = merchant.MerchantID
	transaction.OutletID = input.OutletID
	transaction.BillTotal = input.BillTotal
	transaction.CreatedBy = input.User.ID
	transaction.Uuid = uuid.New().String()
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
