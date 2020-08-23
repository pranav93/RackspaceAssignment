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
	// fmt.Println(models.ProductsDBMap)
	// fmt.Println(models.RulesDBMap)
	// fmt.Println(models.ActionsDBMap)
	// fmt.Println(models.ResultsDBMap)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	r.GET("/cart/:id", controllers.GetCart)
	r.POST("/cart", controllers.CreateCart)
	r.PATCH("/cart/:id", controllers.UpdateCart)
	r.DELETE("/cart/:id", controllers.DeleteCart)

	r.Run()
}
