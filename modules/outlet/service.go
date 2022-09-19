package outlet

import (
	"errors"
	"test-majoo-new/modules/merchant"

	"github.com/google/uuid"
)

type Service interface {
	GetAllOutlets() (Outlets, error)
	GetOutletByMerchantID(MerchantID uint64) (Outlets, error)
	GetOutletByID(ID uint64) (Outlet, error)
	CreateOutlet(input CreateOutletInput) (Outlet, error)
}

type service struct {
	repository         Repository
	merchantrepository merchant.Repository
}

func NewService(repository Repository, merchantrepository merchant.Repository) *service {
	return &service{repository, merchantrepository}
}

func (s *service) GetAllOutlets() (Outlets, error) {

	outlets, err := s.repository.FindAll()
	if err != nil {
		return outlets, err
	}

	return outlets, nil
}

func (s *service) GetOutletByMerchantID(MerchantID uint64) (Outlets, error) {
	outlet, err := s.repository.FindByMerchantID(MerchantID)
	if err != nil {
		return outlet, err
	}

	// if outlet.ID == 0 {
	// 	return outlet, errors.New("No merchant found on with that User ID")
	// }

	return outlet, nil
}

func (s *service) GetOutletByID(ID uint64) (Outlet, error) {
	outlet, err := s.repository.FindByID(ID)

	if err != nil {
		return outlet, err
	}

	return outlet, nil

}

func (s *service) CreateOutlet(input CreateOutletInput) (Outlet, error) {
	outlet := Outlet{}

	merchant, err := s.merchantrepository.FindByUserID(input.User.ID)
	if err != nil {
		return outlet, err
	}
	if merchant.UserID != input.User.ID {
		return outlet, errors.New("Not an owner of the merchant")
	}
	outlet.MerchantID = merchant.ID
	outlet.OutletName = input.OutletName
	outlet.CreatedBy = input.User.ID
	outlet.Uuid = uuid.New().String()
	newMerchant, err := s.repository.Save(outlet)
	if err != nil {
		return newMerchant, err
	}

	return newMerchant, nil
}

// func (s *service) UpdateMerchant(ID uint64, inputData CreateMerchantInput) (Merchant, error) {
// 	merchant, err := s.repository.FindByID(ID)
// 	if err != nil {
// 		return merchant, err
// 	}

// 	if merchant.UserID != inputData.User.ID {
// 		return merchant, errors.New("Not an owner of the merchant")
// 	}

// 	merchant.MerchantName = inputData.MerchantName

// 	updatedMerchant, err := s.repository.Update(merchant)
// 	if err != nil {
// 		return updatedMerchant, err
// 	}

// 	return updatedMerchant, nil
// }
