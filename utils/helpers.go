package helpers

import (
	"ecommerce_api/initializers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyExistence(model interface{}, id string, c *gin.Context) {
	notFoundMessage := fmt.Sprintf("Could not delete. No products found with ID of %s", id)
	if result := initializers.DB.First(&model, id); result.Error.Error() == "record not found" {
		c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
}
