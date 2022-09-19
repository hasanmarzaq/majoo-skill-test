package outlet

import user "test-majoo-new/modules/User"

type CreateOutletInput struct {
	MerchantID uint64 `json:"merchant_id" `
	OutletName string `json:"outlet_name" form:"outlet_name" binding:"required"`
	User       user.User
}
