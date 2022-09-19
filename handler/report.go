package handler

import (
	"net/http"
	"test-majoo-new/helper"
	user "test-majoo-new/modules/User"
	"test-majoo-new/modules/report"

	"github.com/gin-gonic/gin"
)

type reportHandler struct {
	service report.Service
}

func NewReportHandler(service report.Service) *reportHandler {
	return &reportHandler{service}
}

func (h *reportHandler) ReportMerchant(c *gin.Context) {
	var input report.InputReportMerchant
	pagination := helper.GeneratePaginationRequest(c)

	// data := h.staffService.Pagination(c, pagination)
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.UserID = currentUser.ID
	merchants, paginations, err := h.service.GetReportMerchant(c, input, pagination)
	if err != nil {
		response := helper.APIResponse("Error to get merchants", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of merchants", http.StatusOK, "success", gin.H{
		"data":       merchants,
		"pagination": paginations,
	})
	c.JSON(http.StatusOK, response)
}

func (h *reportHandler) ReportMerchantByid(c *gin.Context) {
	var input report.InputReportMerchantByid
	pagination := helper.GeneratePaginationRequest(c)

	// data := h.staffService.Pagination(c, pagination)
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.UserID = currentUser.ID
	merchants, paginations, err := h.service.GetReportByMerchantID(c, input, pagination)
	if err != nil {
		response := helper.APIResponse("Error to get merchants", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of merchants", http.StatusOK, "success", gin.H{
		"data":       merchants,
		"pagination": paginations,
	})
	c.JSON(http.StatusOK, response)
}

func (h *reportHandler) ReportOutletByid(c *gin.Context) {
	var input report.InputReportMerchantOutlet
	pagination := helper.GeneratePaginationRequest(c)

	// data := h.staffService.Pagination(c, pagination)
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("report failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.UserID = currentUser.ID
	merchants, paginations, err := h.service.GetReportMerchantOutlet(c, input, pagination)
	if err != nil {
		response := helper.APIResponse("Error to get merchants", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of merchants", http.StatusOK, "success", gin.H{
		"data":       merchants,
		"pagination": paginations,
	})
	c.JSON(http.StatusOK, response)
}
