package controllers

import (
	"fmt"
	"net/http"

	"github.com/pranav93/RackspaceAssignment/models"

	"github.com/gin-gonic/gin"
)

// CreateCartInput is an input for cart creation
type CreateCartInput struct {
	CartItems []struct {
		ProductCode string `json:"productCode"`
	} `json:"cartItems"`
}

// CreateCart Creates a cart
func CreateCart(c *gin.Context) {
	var input CreateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)

	cartID := models.CreateCart()
	fmt.Println(cartID)
	for i := 0; i < len(input.CartItems); i++ {
		models.AddToCart(cartID, input.CartItems[i].ProductCode, 1)
	}

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
		Add []struct {
			ProductCode string `json:"productCode"`
		} `json:"add"`
		Remove []struct {
			CartItemID string `json:"cartItemId"`
		} `json:"remove"`
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
	fmt.Println(input)

	for i := 0; i < len(input.CartItems.Add); i++ {
		err := models.AddToCart(cartID, input.CartItems.Add[i].ProductCode, 1)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	for i := 0; i < len(input.CartItems.Remove); i++ {
		err := models.RemoveFromCart(cartID, input.CartItems.Remove[i].CartItemID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

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
