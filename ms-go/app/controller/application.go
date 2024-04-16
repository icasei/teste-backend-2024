package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "[ms-go] | Success", "status": http.StatusOK})
}
