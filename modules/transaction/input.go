package transaction

import user "test-majoo-new/modules/User"

type CreateTransactionInput struct {
	MerchantID uint64  `json:"merchant_id"`
	OutletID   uint64  `json:"outlet_id" form:"outlet_id" binding:"required"`
	BillTotal  float64 `json:"bill_total" form:"bill_total" binding:"required"`
	User       user.User
}
