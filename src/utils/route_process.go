package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PaginationParams holds the pagination parameters
type PaginationParams struct {
	Page int
	Size int
}

const DEFAULT_PAGE = 1
const DEFAULT_SIZE = 10

// GetPaginationParams extracts and validates pagination parameters from the request context
func GetPaginationParams(pageStr, sizeStr string) PaginationParams {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = DEFAULT_PAGE
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = DEFAULT_SIZE
	}

	return PaginationParams{
		Page: page,
		Size: size,
	}
}

// GetObjectIDFromParam extracts and validates a MongoDB ObjectID from a URL parameter
func GetObjectIDFromParam(c *gin.Context, param string) (primitive.ObjectID, bool) {
	id := c.Param(param)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return primitive.NilObjectID, false
	}
	return objID, true
}

func GetIdStrFromParam(c *gin.Context, param string) string {
	id := c.Param(param)
	return id
}
