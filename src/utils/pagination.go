package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationParams holds the pagination parameters
type PaginationParams struct {
	Page int
	Size int
}

// GetPaginationParams extracts and validates pagination parameters from the request context
func GetPaginationParams(c *gin.Context, defaultPage, defaultSize int) PaginationParams {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = defaultPage
	}

	sizeStr := c.Query("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = defaultSize
	}

	return PaginationParams{
		Page: page,
		Size: size,
	}
}
