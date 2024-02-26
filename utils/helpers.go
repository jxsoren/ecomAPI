package helpers

import (
	"ecommerce_api/initializers"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func VerifyExistence(model interface{}, id string, c *gin.Context) {
	notFoundMessage := fmt.Sprintf("Could not delete. No products found with ID of %s", id)
	if result := initializers.DB.First(model, id); result.Error != nil {
		// Search for record not found err in error tree
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
}
