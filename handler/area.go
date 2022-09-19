package handler

import (
	"net/http"
	"test-majoo-new/helper"
	"test-majoo-new/modules/area"

	"github.com/gin-gonic/gin"
)

type areaHandler struct {
	service area.Service
}

func NewAreaHandler(service area.Service) *areaHandler {
	return &areaHandler{service}
}

func (h *areaHandler) CreateArea(c *gin.Context) {
	var input area.CreateAreaInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to create area", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentUser := c.MustGet("currentUser").(user.User)

	// input.User = currentUser

	newArea, err := h.service.InsertArea(input.Param1, input.Param2, input.TypeArea)
	if err != nil {
		response := helper.APIResponse("Failed to create area", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create area", http.StatusOK, "success", newArea)
	c.JSON(http.StatusOK, response)
}
