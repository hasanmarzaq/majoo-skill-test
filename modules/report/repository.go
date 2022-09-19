package report

// package merchant

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"test-majoo-new/helper"
	"test-majoo-new/modules/merchant"

	// "test-majoo-new/modules/merchant"

	"gorm.io/gorm"
)

type Repository interface {
	FindReportPerMerchant(input InputReportMerchant, pagination *helper.Pagination) ([]GetReportPerMerchant, *helper.Pagination, error)
	FindByUserIDMerchant(UserID uint64, MerchantID uint64) (merchant.Merchant, error)
	FindReportPerMerchantID(input InputReportMerchantByid, pagination *helper.Pagination) ([]GetReportPerMerchant, *helper.Pagination, error)
	FindMerchantOutlet(UserID uint64, OutletID uint64) (MerchantCheckOutlet, error)
	FindReportPerOutlet(input InputReportMerchantOutlet, pagination *helper.Pagination) (GetReportPerOutlets, *helper.Pagination, error)
}
type repository struct {
	db       *gorm.DB
	dbManual *sql.DB
}

func NewRepository(db *gorm.DB, dbManual *sql.DB) *repository {
	return &repository{db, dbManual}
}

func (r *repository) FindReportPerMerchant(input InputReportMerchant, pagination *helper.Pagination) ([]GetReportPerMerchant, *helper.Pagination, error) {
	var merchant MerchantDetailUser
	totalPages, fromRow, toRow := 0, 0, 0

	err := r.db.Debug().Where("user_id = ?", input.UserID).Table("merchants").Find(&merchant).Error
	// fmt.Println(merchant)
	if err != nil {
		return nil, nil, err
	}

	var GetReportPerMerchants []GetReportPerMerchant
	sqlPaging := fmt.Sprintf("select tbl_tanggal.gen_date,ifnull(tbl_transaction.merchant_name,'%s') as merchant_name,ifnull(tbl_transaction.omset,0) as omset from (select * from (select adddate('1970-01-01',t4*10000 + t3*1000 + t2*100 + t1*10 + t0) gen_date from (select 0 t0 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t0,(select 0 t1 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t1,(select 0 t2 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t2,(select 0 t3 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t3,(select 0 t4 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t4) v where gen_date between '%s' and '%s') tbl_tanggal left join (SELECT trx.merchant_id,m.merchant_name,date(trx.created_at) as tanggal,sum(trx.bill_total) as omset FROM `transactions` trx LEFT JOIN `merchants` m on m.id = trx.merchant_id LEFT JOIN `users` u on u.id =m.user_id where u.id =%d group by merchant_id,date(created_at)) tbl_transaction on tbl_tanggal.gen_date = tbl_transaction.tanggal", merchant.MerchantName, input.StartDate, input.EndDate, input.UserID)

	offset := (pagination.Page - 1) * pagination.Limit

	sqlPaging = fmt.Sprintf("%s ORDER BY tbl_tanggal.gen_date ASC LIMIT %d OFFSET %d", sqlPaging, pagination.Limit, offset)
	fmt.Println(sqlPaging)
	// fmt.Println(sqlPaging)
	merchants, err := r.dbManual.Query(sqlPaging)

	if err != nil {
		return nil, nil, err
	}

	for merchants.Next() {
		// var role Role
		var getReportPerMerchant GetReportPerMerchant

		if err := merchants.Scan(&getReportPerMerchant.Date, &getReportPerMerchant.MerchantName, &getReportPerMerchant.Omzet); err != nil {
			fmt.Println(err)

			// panic(err.Error())
			return nil, nil, err
		}

		GetReportPerMerchants = append(GetReportPerMerchants, getReportPerMerchant)

	}
	var totalRows int64

	sqlCount := "select count(tbl_tanggal.gen_date) as jumlah from (select * from (select adddate('1970-01-01',t4*10000 + t3*1000 + t2*100 + t1*10 + t0) gen_date from (select 0 t0 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t0, (select 0 t1 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t1, (select 0 t2 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t2, (select 0 t3 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t3,  (select 0 t4 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t4) v where gen_date between '2021-11-01' and '2021-11-30') tbl_tanggal"
	fmt.Println(sqlCount)

	if err := r.dbManual.QueryRow(sqlCount).Scan(&totalRows); err != nil {
		return nil, nil, err
	}

	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	// var limit = pagination.Limit
	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}
	if pagination.Page < 1 || pagination.Page > totalPages {
		return nil, nil, errors.New("page not found")
	}

	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow
	pagination.TotalRows = int(totalRows)
	pagination.PageCount = int(totalPages)

	pages := make([]int, totalPages)
	for i := 0; i < totalPages; i++ {
		pages[i] = i + 1
	}
	pagination.Pages = pages

	// fmt.Println(GetReportPerMerchants)

	return GetReportPerMerchants, pagination, nil
}

func (r *repository) FindReportPerMerchantID(input InputReportMerchantByid, pagination *helper.Pagination) ([]GetReportPerMerchant, *helper.Pagination, error) {
	var merchant MerchantDetailUser
	totalPages, fromRow, toRow := 0, 0, 0

	err := r.db.Where("user_id = ?", input.UserID).Table("merchants").Find(&merchant).Error
	fmt.Println(merchant)
	if err != nil {
		return nil, nil, err
	}

	var GetReportPerMerchants []GetReportPerMerchant
	sqlPaging := fmt.Sprintf("select tbl_tanggal.gen_date,ifnull(tbl_transaction.merchant_name,'%s') as merchant_name,ifnull(tbl_transaction.omset,0) as omset from (select * from (select adddate('1970-01-01',t4*10000 + t3*1000 + t2*100 + t1*10 + t0) gen_date from (select 0 t0 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t0,(select 0 t1 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t1,(select 0 t2 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t2,(select 0 t3 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t3,(select 0 t4 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t4) v where gen_date between '%s' and '%s') tbl_tanggal left join (SELECT trx.merchant_id,m.merchant_name,date(trx.created_at) as tanggal,sum(trx.bill_total) as omset FROM `transactions` trx LEFT JOIN `merchants` m on m.id = trx.merchant_id LEFT JOIN `users` u on u.id =m.user_id where u.id =%d and trx.merchant_id =%d  group by merchant_id,date(created_at)) tbl_transaction on tbl_tanggal.gen_date = tbl_transaction.tanggal", merchant.MerchantName, input.StartDate, input.EndDate, input.UserID, input.MerchantID)
	offset := (pagination.Page - 1) * pagination.Limit

	sqlPaging = fmt.Sprintf("%s ORDER BY tbl_tanggal.gen_date ASC LIMIT %d OFFSET %d", sqlPaging, pagination.Limit, offset)

	// fmt.Println(sqlPaging)
	merchants, err := r.dbManual.Query(sqlPaging)

	if err != nil {
		return nil, nil, err
	}

	for merchants.Next() {
		// var role Role
		var getReportPerMerchant GetReportPerMerchant

		if err := merchants.Scan(&getReportPerMerchant.Date, &getReportPerMerchant.MerchantName, &getReportPerMerchant.Omzet); err != nil {
			fmt.Println(err)

			// panic(err.Error())
			return nil, nil, err
		}

		GetReportPerMerchants = append(GetReportPerMerchants, getReportPerMerchant)

	}
	var totalRows int64

	sqlCount := "select count(tbl_tanggal.gen_date) as jumlah from (select * from (select adddate('1970-01-01',t4*10000 + t3*1000 + t2*100 + t1*10 + t0) gen_date from (select 0 t0 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t0, (select 0 t1 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t1, (select 0 t2 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t2, (select 0 t3 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t3,  (select 0 t4 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t4) v where gen_date between '2021-11-01' and '2021-11-30') tbl_tanggal"
	fmt.Println(sqlCount)

	if err := r.dbManual.QueryRow(sqlCount).Scan(&totalRows); err != nil {
		return nil, nil, err
	}

	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	// var limit = pagination.Limit
	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}
	if pagination.Page < 1 || pagination.Page > totalPages {
		return nil, nil, errors.New("page not found")
	}

	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow
	pagination.TotalRows = int(totalRows)
	pagination.PageCount = int(totalPages)

	pages := make([]int, totalPages)
	for i := 0; i < totalPages; i++ {
		pages[i] = i + 1
	}
	pagination.Pages = pages

	// fmt.Println(GetReportPerMerchants)

	return GetReportPerMerchants, pagination, nil
}

func (r *repository) FindReportPerOutlet(input InputReportMerchantOutlet, pagination *helper.Pagination) (GetReportPerOutlets, *helper.Pagination, error) {
	totalPages, fromRow, toRow := 0, 0, 0
	var GetReportPerOutlets GetReportPerOutlets
	sqlPaging := fmt.Sprintf("select tbl_tanggal.gen_date,ifnull(tbl_transaction.merchant_name,'%s') as merchant_name,ifnull(tbl_transaction.outlet_name,'%s') as outlet_name,ifnull(tbl_transaction.omset,0) as omset from (select * from (select adddate('1970-01-01',t4*10000 + t3*1000 + t2*100 + t1*10 + t0) gen_date from (select 0 t0 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t0,(select 0 t1 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t1,(select 0 t2 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t2,(select 0 t3 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t3,(select 0 t4 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t4) v where gen_date between '%s' and '%s') tbl_tanggal left join (SELECT trx.merchant_id,m.merchant_name,o.outlet_name,date(trx.created_at) as tanggal,sum(trx.bill_total) as omset FROM `transactions` trx LEFT JOIN `outlets` o on o.id =trx.outlet_id LEFT JOIN `merchants` m on m.id = trx.merchant_id LEFT JOIN `users` u on u.id =m.user_id where u.id =%d and trx.merchant_id =%d and outlet_id =%d  group by trx.merchant_id,trx.outlet_id,date(trx.created_at)) tbl_transaction on tbl_tanggal.gen_date = tbl_transaction.tanggal", input.MerchantName, input.OutletName, input.StartDate, input.EndDate, input.UserID, input.MerchantID, input.OutletID)
	offset := (pagination.Page - 1) * pagination.Limit

	sqlPaging = fmt.Sprintf("%s ORDER BY tbl_tanggal.gen_date ASC LIMIT %d OFFSET %d", sqlPaging, pagination.Limit, offset)
	fmt.Println(sqlPaging)
	merchants, err := r.dbManual.Query(sqlPaging)
	if err != nil {
		return nil, nil, err
	}

	for merchants.Next() {
		var GetReportPerOutlet GetReportPerOutlet
		if err := merchants.Scan(&GetReportPerOutlet.Date, &GetReportPerOutlet.MerchantName, &GetReportPerOutlet.OutletName, &GetReportPerOutlet.Omzet); err != nil {
			// fmt.Println(err)

			// panic(err.Error())
			return nil, nil, err
		}

		GetReportPerOutlets = append(GetReportPerOutlets, GetReportPerOutlet)

	}
	var totalRows int64

	sqlCount := "select count(tbl_tanggal.gen_date) as jumlah from (select * from (select adddate('1970-01-01',t4*10000 + t3*1000 + t2*100 + t1*10 + t0) gen_date from (select 0 t0 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t0, (select 0 t1 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t1, (select 0 t2 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t2, (select 0 t3 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t3,  (select 0 t4 union select 1 union select 2 union select 3 union select 4 union select 5 union select 6 union select 7 union select 8 union select 9) t4) v where gen_date between '2021-11-01' and '2021-11-30') tbl_tanggal"
	fmt.Println(sqlCount)

	if err := r.dbManual.QueryRow(sqlCount).Scan(&totalRows); err != nil {
		return nil, nil, err
	}

	totalPages = int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))

	// var limit = pagination.Limit
	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}
	if pagination.Page < 1 || pagination.Page > totalPages {
		return nil, nil, errors.New("page not found")
	}

	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow
	pagination.TotalRows = int(totalRows)
	pagination.PageCount = int(totalPages)

	pages := make([]int, totalPages)
	for i := 0; i < totalPages; i++ {
		pages[i] = i + 1
	}
	pagination.Pages = pages

	// fmt.Println(GetReportPerMerchants)

	return GetReportPerOutlets, pagination, nil
}

func (r *repository) FindByUserIDMerchant(UserID uint64, MerchantID uint64) (merchant.Merchant, error) {
	var merchant merchant.Merchant

	err := r.db.Preload("User").Where("id = ? And user_id = ?", MerchantID, UserID).Find(&merchant).Error

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repository) FindMerchantOutlet(UserID uint64, OutletID uint64) (MerchantCheckOutlet, error) {
	var merchant MerchantCheckOutlet
	sqlCount := fmt.Sprintf("SELECT m.id as merchant_id,m.merchant_name,o.outlet_name FROM `merchants` m LEFT JOIN `outlets` o on o.merchant_id =m.id LEFT JOIN users u on u.id = m.user_id where u.id = %d and o.id = %d ", UserID, OutletID)
	// fmt.Println(sqlCount)

	if err := r.dbManual.QueryRow(sqlCount).Scan(&merchant.MerchantID, &merchant.MerchantName, &merchant.OutletName); err != nil {
		return merchant, errors.New("Not an owner of the Outlet")
	}

	return merchant, nil
}
