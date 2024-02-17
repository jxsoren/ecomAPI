package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello!!! 🐹 🐹 🐹 \n")
}
