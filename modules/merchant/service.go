package merchant

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	GetAllMerchant() (Merchants, error)
	GetMerchantByUserID(userID uint64) (Merchant, error)
	GetMerchantByID(ID uint64) (Merchant, error)
	CreateMerchant(input CreateMerchantInput) (Merchant, error)
	UpdateMerchant(ID uint64, inputData CreateMerchantInput) (Merchant, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllMerchant() (Merchants, error) {

	merchants, err := s.repository.FindAll()
	if err != nil {
		return merchants, err
	}

	return merchants, nil
}

func (s *service) GetMerchantByUserID(userID uint64) (Merchant, error) {
	merchant, err := s.repository.FindByUserID(userID)
	if err != nil {
		return merchant, err
	}

	if merchant.ID == 0 {
		return merchant, errors.New("No merchant found on with that User ID")
	}

	return merchant, nil
}

func (s *service) GetMerchantByID(ID uint64) (Merchant, error) {
	merchant, err := s.repository.FindByID(ID)

	if err != nil {
		return merchant, err
	}

	return merchant, nil

}

func (s *service) CreateMerchant(input CreateMerchantInput) (Merchant, error) {
	merchant := Merchant{}

	merchant.UserID = input.User.ID
	merchant.MerchantName = input.MerchantName
	merchant.CreatedBy = input.User.ID
	merchant.Uuid = uuid.New().String()

	newMerchant, err := s.repository.Save(merchant)
	if err != nil {
		return newMerchant, err
	}

	return newMerchant, nil
}

func (s *service) UpdateMerchant(ID uint64, inputData CreateMerchantInput) (Merchant, error) {
	merchant, err := s.repository.FindByID(ID)
	if err != nil {
		return merchant, err
	}

	if merchant.UserID != inputData.User.ID {
		return merchant, errors.New("Not an owner of the merchant")
	}

	merchant.MerchantName = inputData.MerchantName

	updatedMerchant, err := s.repository.Update(merchant)
	if err != nil {
		return updatedMerchant, err
	}

	return updatedMerchant, nil
}
