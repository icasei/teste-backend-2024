package router

import (
	"fmt"
	"log"
	"ms-go/app/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := SetupRouter()

	fmt.Println("* Listening on tcp://localhost:3030")

	if err := r.Run(":3030"); err != nil {
		log.Fatal("Unable to start the server")
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("", controller.IndexHome)

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/products", controller.IndexProducts)
		apiV1.GET("/products/:id", controller.ShowProducts)
		apiV1.POST("/products", controller.CreateProducts)
		apiV1.PATCH("/products/:id", controller.UpdateProducts)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Page not found"})
	})

	return router
}
