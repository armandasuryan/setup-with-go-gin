package utils

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	FlexibleType       interface{}
	PaginationResponse struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Error      string      `json:"error,omitempty"`
		Data       interface{} `json:"data,omitempty"`
		Meta       Meta        `json:"meta,omitempty"`
	}

	Meta struct {
		Pagination Pagination `json:"pagination"`
	}

	Pagination struct {
		Total       int   `json:"total"`
		Count       int   `json:"count"`
		PerPage     int   `json:"per_page"`
		CurrentPage int   `json:"current_page"`
		TotalPages  int   `json:"total_pages"`
		Links       Links `json:"links"`
	}

	Links struct {
		Next string `json:"next"`
	}
)

func GetPaginated(ctx *gin.Context, page, limit int, data interface{}) (interface{}, Pagination) {
	// Check whether the data sent is sliced (partial data) or not
	sliceValue := reflect.ValueOf(data)
	if sliceValue.Kind() != reflect.Slice {
		return nil, Pagination{}
	}

	// get length data
	dataLength := sliceValue.Len()

	// get startindex and endindex
	startIndex := (page - 1) * limit
	endIndex := page * limit
	if endIndex > dataLength {
		endIndex = dataLength
	}

	// Check if the current page has data, if not set to the previous page
	if startIndex >= dataLength && page > 1 {
		page--
		startIndex = (page - 1) * limit
		endIndex = page * limit
		if endIndex > dataLength {
			endIndex = dataLength
		}
	}

	// get data after paginated
	paginateddata := sliceValue.Slice(startIndex, endIndex).Interface()

	pagination := Pagination{
		Total:       dataLength,
		Count:       endIndex - startIndex,
		PerPage:     limit,
		CurrentPage: page,
		TotalPages:  (dataLength + limit - 1) / limit,
		Links: Links{
			Next: getNextPageURL(ctx, page, limit, dataLength),
		},
	}

	return paginateddata, pagination
}

func getNextPageURL(ctx *gin.Context, currentPage, limit, totalItems int) string {
	if (currentPage * limit) >= totalItems {
		return ""
	}
	nextPage := currentPage + 1

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := scheme + "://" + ctx.Request.Host

	return baseURL + ctx.Request.URL.Path + "?page=" + strconv.Itoa(nextPage) + "&limit=" + strconv.Itoa(limit)
}
