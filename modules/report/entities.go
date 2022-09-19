package report

type GetReportPerMerchant struct {
	Date         string  `json:"date"`
	MerchantName string  `json:"merchant_name" `
	Omzet        float64 `json:"omzet"`
}
type GetReportPerMerchants []GetReportPerMerchant

type GetReportPerOutlet struct {
	Date         string  `json:"date"`
	MerchantName string  `json:"merchant_name" `
	OutletName   string  `json:"outlet_name" `
	Omzet        float64 `json:"omzet"`
}
type GetReportPerOutlets []GetReportPerOutlet

type MerchantDetailUser struct {
	MerchantName string `json:"merchant_name" form:"merchant_name" binding:"required"`
	// User         user.User
}

type MerchantCheckOutlet struct {
	MerchantID   uint64 `gorm:"not null;" json:"merchant_id"`
	MerchantName string `json:"merchant_name" form:"merchant_name"`
	OutletName   string `json:"outlet_name"`
}
