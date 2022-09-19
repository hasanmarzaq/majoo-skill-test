package report

type InputReportMerchant struct {
	UserID    uint64 `json:"user_id"`
	StartDate string `json:"start_date" form:"start_date" binding:"required"`
	EndDate   string `json:"end_date" form:"end_date" binding:"required"`
}

type InputReportMerchantByid struct {
	UserID     uint64 `json:"user_id"`
	MerchantID uint64 `json:"merchant_id" form:"merchant_id" binding:"required"`
	StartDate  string `json:"start_date" form:"start_date" binding:"required"`
	EndDate    string `json:"end_date" form:"end_date" binding:"required"`
}

type InputReportMerchantOutlet struct {
	UserID       uint64 `json:"user_id"`
	MerchantID   uint64 `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	OutletName   string `json:"outlet_name"`
	OutletID     uint64 `json:"outlet_id" form:"outlet_id" binding:"required"`
	StartDate    string `json:"start_date" form:"start_date" binding:"required"`
	EndDate      string `json:"end_date" form:"end_date" binding:"required"`
}
