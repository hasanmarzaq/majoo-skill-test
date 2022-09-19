package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	FromRow   int `json:"from_row"`
	ToRow     int `json:"to_row"`
	TotalRows int `json:"total_rows"`
	Limit     int `json:"limit"`
	// Sort         string `json:"sort"`
	Page         int    `json:"page"`
	PageCount    int    `json:"page_count"`
	Pages        []int  `json:"pages"`
	FirstPage    string `json:"first_page"`
	PreviousPage string `json:"previous_page"`
	NextPage     string `json:"next_page"`
	LastPage     string `json:"last_page"`
}

func GeneratePaginationRequest(context *gin.Context) *Pagination {
	// default limit, page & sort parameter
	limit := 10
	page := 1

	query := context.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
			// case "sort":
			// 	sort = queryValue
			// 	break
		}

	}

	return &Pagination{Limit: limit, Page: page}
}
