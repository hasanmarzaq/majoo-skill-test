package merchant

import (
	// "os/user"
	user "test-majoo-new/modules/User"

	// "test-majoo-new/modules/transaction"
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint64 `gorm:"not null;" json:"user_id"`
	MerchantName string `gorm:"size:40;not null" json:"merchant_name"`
	Uuid         string `gorm:"size:50;not null" json:"uuid"`
	CreatedAt    time.Time
	CreatedBy    uint64
	UpdatedAt    time.Time
	UpdatedBy    uint64
	DeletedAt    gorm.DeletedAt
	DeletedBy    uint64
	User         user.User `gorm:"foreignkey:user_id;references:id"`
	// Outlets      outlet.Outlets `gorm:"foreignkey:MerchantID"`
	// Transactions transaction.Transactions `gorm:"foreignkey:MerchantID"`
}
type Merchants []Merchant
