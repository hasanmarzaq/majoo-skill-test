package merchant

import user "test-majoo-new/modules/User"

type CreateMerchantInput struct {
	MerchantName string `json:"merchant_name" form:"merchant_name" binding:"required"`
	User         user.User
}
