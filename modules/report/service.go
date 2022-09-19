package report

import (
	"errors"
	"fmt"
	"test-majoo-new/helper"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetReportMerchant(context *gin.Context, input InputReportMerchant, pagination *helper.Pagination) (GetReportPerMerchants, *helper.Pagination, error)
	GetReportByMerchantID(context *gin.Context, input InputReportMerchantByid, pagination *helper.Pagination) (GetReportPerMerchants, *helper.Pagination, error)
	GetReportMerchantOutlet(context *gin.Context, input InputReportMerchantOutlet, pagination *helper.Pagination) (GetReportPerOutlets, *helper.Pagination, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// report merchant ambil merchant langsung dari login
func (s *service) GetReportMerchant(context *gin.Context, input InputReportMerchant, pagination *helper.Pagination) (GetReportPerMerchants, *helper.Pagination, error) {
	fmt.Println(input)
	GetReportPerMerchants, pagination, err := s.repository.FindReportPerMerchant(input, pagination)
	if err != nil {
		return GetReportPerMerchants, pagination, err
	}

	urlPath := context.Request.URL.Path

	// set first & last page pagination response
	pagination.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, 0)
	pagination.LastPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.PageCount)

	if pagination.Page > 0 {
		// set previous page pagination response
		pagination.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.Page-1)
	}

	if pagination.Page < pagination.PageCount {
		// set next page pagination response
		pagination.NextPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.Page+1)
	}

	if pagination.Page > pagination.PageCount {
		// reset previous page
		pagination.PreviousPage = ""
	}

	return GetReportPerMerchants, pagination, nil
}

//  report merchant bisa select merchant nya sendiri
func (s *service) GetReportByMerchantID(context *gin.Context, input InputReportMerchantByid, pagination *helper.Pagination) (GetReportPerMerchants, *helper.Pagination, error) {
	merchant, err := s.repository.FindByUserIDMerchant(input.UserID, input.MerchantID)
	if err != nil {
		return nil, nil, err
	}

	if merchant.UserID != input.UserID {
		fmt.Println(errors.New("Not an owner of the Merchant"))

		return nil, nil, errors.New("Not an owner of the Merchant")
	}

	GetReportPerMerchants, pagination, err := s.repository.FindReportPerMerchantID(input, pagination)
	if err != nil {
		return nil, nil, err
	}

	urlPath := context.Request.URL.Path
	// set first & last page pagination response
	pagination.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, 0)
	pagination.LastPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.PageCount)
	if pagination.Page > 0 {
		// set previous page pagination response
		pagination.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.Page-1)
	}
	if pagination.Page < pagination.PageCount {
		// set next page pagination response
		pagination.NextPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.Page+1)
	}
	if pagination.Page > pagination.PageCount {
		// reset previous page
		pagination.PreviousPage = ""
	}

	return GetReportPerMerchants, pagination, nil
}

// report merchant per outlet
func (s *service) GetReportMerchantOutlet(context *gin.Context, input InputReportMerchantOutlet, pagination *helper.Pagination) (GetReportPerOutlets, *helper.Pagination, error) {

	merchant, err := s.repository.FindMerchantOutlet(input.UserID, input.OutletID)
	if err != nil {
		return nil, nil, err
	}
	input.MerchantID = merchant.MerchantID
	input.MerchantName = merchant.MerchantName
	input.OutletName = merchant.OutletName

	GetReportPerOutlets, pagination, err := s.repository.FindReportPerOutlet(input, pagination)
	if err != nil {
		return GetReportPerOutlets, pagination, err
	}

	urlPath := context.Request.URL.Path
	// set first & last page pagination response
	pagination.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, 0)
	pagination.LastPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.PageCount)
	if pagination.Page > 0 {
		// set previous page pagination response
		pagination.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.Page-1)
	}
	if pagination.Page < pagination.PageCount {
		// set next page pagination response
		pagination.NextPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, pagination.Limit, pagination.Page+1)
	}
	if pagination.Page > pagination.PageCount {
		// reset previous page
		pagination.PreviousPage = ""
	}
	return GetReportPerOutlets, pagination, nil
}
