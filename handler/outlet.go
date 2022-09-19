package handler

import (
	"net/http"
	"test-majoo-new/helper"

	// merchant "test-majoo-new/modules/merchant"
	user "test-majoo-new/modules/User"
	"test-majoo-new/modules/outlet"

	"github.com/gin-gonic/gin"
)

type outletHandler struct {
	service outlet.Service
}

func NewOutletHandler(service outlet.Service) *outletHandler {
	return &outletHandler{service}
}

func (h *outletHandler) GetOutlets(c *gin.Context) {

	outlets, err := h.service.GetAllOutlets()
	if err != nil {
		response := helper.APIResponse("Error to get outlets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of outlets", http.StatusOK, "success", outlets)
	c.JSON(http.StatusOK, response)
}

// func (h *merchantHandler) GetOutletByMerchantID(c *gin.Context) {
// 	currentUser := c.MustGet("currentUser").(user.User)
// 	userID := currentUser.ID
// 	merchants, err := h.service.GetMerchantByUserID(userID)
// 	if err != nil {
// 		response := helper.APIResponse("Error to get merchants", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	response := helper.APIResponse("List of merchants", http.StatusOK, "success", merchants)
// 	c.JSON(http.StatusOK, response)
// }

func (h *outletHandler) CreateOutlet(c *gin.Context) {
	var input outlet.CreateOutletInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create outlet", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newOutlet, err := h.service.CreateOutlet(input)
	if err != nil {
		response := helper.APIResponse("Failed to create outlet", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create outlet", http.StatusOK, "success", newOutlet)
	c.JSON(http.StatusOK, response)
}
