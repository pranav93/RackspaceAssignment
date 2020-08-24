package controllers

import (
	"log"
	"net/http"

	"github.com/pranav93/RackspaceAssignment/models"

	"github.com/gin-gonic/gin"
)

// CreateCartInput is an input for cart creation
type CreateCartInput struct {
	CartItems map[string]int `json:"cartItems"`
}

// CreateCart Creates a cart
func CreateCart(c *gin.Context) {
	var input CreateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(input)

	cartID := models.CreateCart()
	log.Println(cartID)
	for k, val := range input.CartItems {
		models.AddToCart(cartID, k, val)
	}
	models.CalculateCart(cartID)
	cart, _ := models.GetCart(cartID)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"cart": cart}})
}

// GetCart gets a cart
func GetCart(c *gin.Context) {
	cart, err := models.GetCart(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"cart": cart}})
}

// UpdateCartInput UpdateCartInput
type UpdateCartInput struct {
	CartItems struct {
		Add    []string `json:"add"`
		Remove []string `json:"remove"`
	} `json:"cartItems"`
}

// UpdateCart Updates a cart
func UpdateCart(c *gin.Context) {
	cartID := c.Param("id")
	if cartID == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"error": "Invalid cart id"}})
		return
	}

	var input UpdateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(input)

	for i := 0; i < len(input.CartItems.Add); i++ {
		err := models.AddToCart(cartID, input.CartItems.Add[i], 1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	for i := 0; i < len(input.CartItems.Remove); i++ {
		err := models.RemoveFromCart(cartID, input.CartItems.Remove[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// I am not using db session and rollback here,
	// for simplicity key-val store in golang map is used
	// so if error occurs, partial operation will be done
	// and calculated value for cart will be wrong
	// Have to figure out how can we rollback it in golang map
	// Or maybe just db should be used
	err := models.ResetDiscount(cartID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.CalculateCart(cartID)
	cart, _ := models.GetCart(cartID)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"cart": cart}})
}

// DeleteCart DeleteCart
func DeleteCart(c *gin.Context) {
	err := models.DeleteCart(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"deleted": true}})
}

// CartCheckOut CartCheckOut
func CartCheckOut(c *gin.Context) {
	cart, err := models.GetCart(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": gin.H{"error": err.Error()}})
		return
	}
	cart.ApplyDiscount()
	cart.Calculate()
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"cart": cart}})
}
