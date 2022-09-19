package transaction

// package merchant

import (
	"database/sql"
	"errors"
	"fmt"
	"test-majoo-new/modules/outlet"

	// "test-majoo-new/modules/merchant"

	"gorm.io/gorm"
)

type Repository interface {
	Save(transaction Transaction) (Transaction, error)
	FindMerchantOutlet(UserID uint64, OutletID uint64) (outlet.Outlet, error)
}
type repository struct {
	db       *gorm.DB
	dbManual *sql.DB
}

func NewRepository(db *gorm.DB, dbManual *sql.DB) *repository {
	return &repository{db, dbManual}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) FindMerchantOutlet(UserID uint64, OutletID uint64) (outlet.Outlet, error) {
	var outlet outlet.Outlet
	sqlCount := fmt.Sprintf("SELECT m.id as merchant_id,o.outlet_name FROM `merchants` m LEFT JOIN `outlets` o on o.merchant_id =m.id LEFT JOIN users u on u.id = m.user_id where u.id = %d and o.id = %d ", UserID, OutletID)
	// fmt.Println(sqlCount)

	if err := r.dbManual.QueryRow(sqlCount).Scan(&outlet.MerchantID, &outlet.OutletName); err != nil {
		return outlet, errors.New("Not an owner of the Outlet")
	}

	return outlet, nil
}
