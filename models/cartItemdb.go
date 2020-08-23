package models

import (
	"github.com/google/uuid"
)

// CartItemsDBMap key-val storage for cart items
var CartItemsDBMap map[string]*CartItem

// CreateCartItem creates a cart item
func CreateCartItem(productCode string) *CartItem {
	cartItemID := uuid.New().String()
	product := ProductsDBMap[productCode]

	cartItem := &CartItem{
		ID:      cartItemID,
		Product: product,
		Price:   product.Price,
	}
	CartItemsDBMap[cartItemID] = cartItem
	return cartItem
}

// DeleteCartItem DeleteCartItem
func DeleteCartItem(cartItemID string) {
	if _, ok := CartItemsDBMap[cartItemID]; ok {
		delete(CartItemsDBMap, cartItemID)
	}
}
