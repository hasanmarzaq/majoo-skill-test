package outlet

import (
	// "test-majoo-new/modules/transaction"
	"test-majoo-new/modules/merchant"
	"time"

	"gorm.io/gorm"
)

type Outlet struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	MerchantID uint64 `gorm:"not null;" json:"merchant_id"`
	OutletName string `gorm:"size:40;not null" json:"outlet_name"`
	Uuid       string `gorm:"size:50;not null" json:"uuid"`
	CreatedAt  time.Time
	CreatedBy  uint64
	UpdatedAt  time.Time
	UpdatedBy  uint64
	DeletedAt  gorm.DeletedAt
	DeletedBy  uint64
	Merchant   merchant.Merchant `gorm:"foreignkey:merchant_id;references:id"`
}
type Outlets []Outlet
