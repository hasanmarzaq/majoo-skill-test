package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:45;null" json:"name"`
	UserName  string `gorm:"size:45;null" json:"user_name"`
	Password  string `gorm:"size:255;null;" json:"password"`
	Uuid      string `gorm:"size:50;not null" json:"uuid"`
	CreatedAt time.Time
	CreatedBy uint
	UpdatedAt time.Time
	UpdatedBy uint
	DeletedAt gorm.DeletedAt
	DeletedBy uint
	// Merchant  merchant.Merchant `gorm:"foreignKey:UserID"`
}
type Users []User
	