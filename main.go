package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pranav93/RackspaceAssignment/controllers"
	"github.com/pranav93/RackspaceAssignment/scripts"
)

func init() {
	scripts.CreateRulesDB()
	scripts.CreateProductsDB()
	scripts.CreateCartsDB()
	// log.Println(models.ProductsDBMap)
	// log.Println(models.RulesDBMap)
	// log.Println(models.ActionsDBMap)
	// log.Println(models.ResultsDBMap)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	r.GET("/cart/:id/", controllers.GetCart)
	r.POST("/cart/save/", controllers.CreateCart)
	r.PATCH("/cart/save/:id/", controllers.UpdateCart)
	r.DELETE("/cart/:id/", controllers.DeleteCart)
	r.POST("/cart/checkout/:id/", controllers.CartCheckOut)

	r.Run()
}
