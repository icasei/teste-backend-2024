package controller

import (
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexProducts(c *gin.Context) {
	all, err := products.ListAll()

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": all})
}

func ShowProducts(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := products.Details(models.Product{ID: id})

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProducts(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := products.Create(params, true)

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func UpdateProducts(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	params.ID = id

	product, err := products.Update(params, true)

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
