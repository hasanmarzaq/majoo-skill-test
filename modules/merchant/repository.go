package merchant

import (
	"database/sql"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (Merchants, error)
	FindByUserID(userID uint64) (Merchant, error)
	FindByID(ID uint64) (Merchant, error)
	Save(merchant Merchant) (Merchant, error)
	Update(merchant Merchant) (Merchant, error)
	// FindByReportMerchant(userID uint64) (Merchant, error)
}
type repository struct {
	db       *gorm.DB
	dbManual *sql.DB
}

func NewRepository(db *gorm.DB, dbManual *sql.DB) *repository {
	return &repository{db, dbManual}
}

func (r *repository) FindAll() (Merchants, error) {
	var merchants Merchants

	err := r.db.Preload("User").Find(&merchants).Error
	if err != nil {
		return merchants, err
	}

	return merchants, nil
}

func (r *repository) FindByUserID(userID uint64) (Merchant, error) {
	var merchant Merchant

	err := r.db.Where("user_id = ?", userID).Preload("User").Find(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repository) FindByID(ID uint64) (Merchant, error) {
	var merchant Merchant

	err := r.db.Preload("User").Where("id = ?", ID).Find(&merchant).Error

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repository) Save(merchant Merchant) (Merchant, error) {
	err := r.db.Create(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repository) Update(merchant Merchant) (Merchant, error) {
	err := r.db.Save(&merchant).Error

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

// func (r *repository) FindByUserIDMerchant(UserID uint64, MerchantID uint64) (Merchant, error) {
// 	var merchant Merchant

// 	err := r.db.Preload("User").Where("id = ? And user_id = ?", MerchantID, UserID).Find(&merchant).Error

// 	if err != nil {
// 		return merchant, err
// 	}

// 	return merchant, nil
// }
