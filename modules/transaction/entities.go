package transaction

import (
	"test-majoo-new/modules/merchant"
	"test-majoo-new/modules/outlet"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID         uint64  `gorm:"primary_key;auto_increment" json:"id"`
	MerchantID uint64  `gorm:"not null;" json:"merchant_id"`
	OutletID   uint64  `gorm:"not null;" json:"outlet_id"`
	BillTotal  float64 `gorm:"not null" json:"bill_total"`
	Uuid       string  `gorm:"size:50;not null" json:"uuid"`
	CreatedAt  time.Time
	CreatedBy  uint64
	UpdatedAt  time.Time
	UpdatedBy  uint64
	DeletedAt  gorm.DeletedAt
	DeletedBy  uint64
	Merchant   merchant.Merchant `gorm:"foreignkey:merchant_id;references:id"`
	Outlet     outlet.Outlet     `gorm:"foreignkey:outlet_id;references:id"`
}
type Transactions []Transaction
