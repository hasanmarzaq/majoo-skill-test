package outlet

import (
	"database/sql"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (Outlets, error)
	FindByMerchantID(MerchantID uint64) (Outlets, error)
	FindByID(ID uint64) (Outlet, error)
	Save(outlet Outlet) (Outlet, error)
	// FindMerchantOutlet(UserID uint64, OutletID uint64) (Outlet, error)
}
type repository struct {
	db       *gorm.DB
	dbManual *sql.DB
}

func NewRepository(db *gorm.DB, dbManual *sql.DB) *repository {
	return &repository{db, dbManual}
}

func (r *repository) FindAll() (Outlets, error) {
	var outlets Outlets

	err := r.db.Preload("Merchant").Preload("Merchant.User").Find(&outlets).Error
	if err != nil {
		return outlets, err
	}

	return outlets, nil
}

func (r *repository) FindByMerchantID(MerchantID uint64) (Outlets, error) {
	var outlet Outlets

	err := r.db.Where("merchant_id = ?", MerchantID).Preload("Merchant").Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) FindByID(ID uint64) (Outlet, error) {
	var outlet Outlet

	err := r.db.Preload("Merchant").Where("id = ?", ID).Find(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
func (r *repository) Save(outlet Outlet) (Outlet, error) {
	err := r.db.Create(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

// func (r *repository) FindMerchantOutlet(UserID uint64, OutletID uint64) (Outlet, error) {
// 	var outlet Outlet
// 	sqlCount := fmt.Sprintf("SELECT m.id as merchant_id,o.outlet_name FROM `merchants` m LEFT JOIN `outlets` o on o.merchant_id =m.id LEFT JOIN users u on u.id = m.user_id where u.id = %d and o.id = %d ", UserID, OutletID)
// 	// fmt.Println(sqlCount)

// 	if err := r.dbManual.QueryRow(sqlCount).Scan(&outlet.MerchantID, &outlet.OutletName); err != nil {
// 		return outlet, errors.New("Not an owner of the Outlet")
// 	}

// 	return outlet, nil
// }
