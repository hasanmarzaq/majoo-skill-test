package handler

import (
	"net/http"
	"test-majoo-new/helper"
	user "test-majoo-new/modules/User"
	merchant "test-majoo-new/modules/merchant"

	"github.com/gin-gonic/gin"
)

type merchantHandler struct {
	service merchant.Service
}

func NewMerchantHandler(service merchant.Service) *merchantHandler {
	return &merchantHandler{service}
}

func (h *merchantHandler) GetMerchants(c *gin.Context) {

	merchants, err := h.service.GetAllMerchant()
	if err != nil {
		response := helper.APIResponse("Error to get merchants", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of merchants", http.StatusOK, "success", merchants)
	c.JSON(http.StatusOK, response)
}

func (h *merchantHandler) GetMerchantByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	merchants, err := h.service.GetMerchantByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Error to get merchants", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of merchants", http.StatusOK, "success", merchants)
	c.JSON(http.StatusOK, response)
}

func (h *merchantHandler) CreateMerchant(c *gin.Context) {
	var input merchant.CreateMerchantInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create merchant", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newMerchant, err := h.service.CreateMerchant(input)
	if err != nil {
		response := helper.APIResponse("Failed to create merchant", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create merchant", http.StatusOK, "success", newMerchant)
	c.JSON(http.StatusOK, response)
}
