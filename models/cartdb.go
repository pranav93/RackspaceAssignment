package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// CartsDBMap key-val storage for carts
var CartsDBMap map[string]*Cart

// CreateCart creates an empty cart
func CreateCart() string {
	cartID := uuid.New().String()
	cart := Cart{
		ID:      cartID,
		Items:   []*CartItem{},
		CartMap: CartMap{},
		Total:   0,
	}
	CartsDBMap[cartID] = &cart
	return cartID
}

// GetCart gets a cart with ID
func GetCart(cartID string) (*Cart, error) {
	if val, ok := CartsDBMap[cartID]; ok {
		return val, nil
	}
	return nil, errors.New("Cart not found")
}

// DeleteCart deletes a cart with ID
func DeleteCart(cartID string) error {
	if val, ok := CartsDBMap[cartID]; ok {
		for i := 0; i < len(val.Items); i++ {
			DeleteCartItem(val.Items[i].ID)
		}
		delete(CartsDBMap, cartID)
		return nil
	}
	return errors.New("Cart not found")
}

// AddToCart adds a product to cart
func AddToCart(cartID string, productCode string, qty int) error {
	var cart *Cart
	log.Println(cartID, productCode, qty)
	if val, ok := CartsDBMap[cartID]; !ok {
		return errors.New("CartID " + cartID + " is invalid")
	} else {
		cart = val
	}
	log.Println(cart)
	// product := ProductsDBMap[productCode]

	for i := 0; i < qty; i++ {
		cartItem := CreateCartItem(productCode)
		cart.AddItem(cartItem)
	}
	log.Println("cart.Items", cart.Items)

	if val, ok := cart.CartMap[productCode]; ok {
		cart.CartMap[productCode] = val + qty
	} else {
		cart.CartMap[productCode] = qty
	}
	log.Println(cart)
	return nil
}

// RemoveFromCart removes a cartItem from cart
func RemoveFromCart(cartID string, cartItemID string) error {
	cart := CartsDBMap[cartID]

	// map can be used instead of items array
	foundIndex := -1
	for i := 0; i < len(cart.Items); i++ {
		if cart.Items[i].ID == cartItemID {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return errors.New("Cart Item with ID " + cartItemID + " does not exist in cart with ID " + cartID)
	}
	productCode := cart.Items[foundIndex].Code

	cart.Items = append(cart.Items[:foundIndex], cart.Items[foundIndex+1:]...)
	if val, ok := cart.CartMap[productCode]; ok {
		cart.CartMap[productCode] = val - 1
	}

	return nil
}

// CalculateCart calculates the cart total
func CalculateCart(cartID string) error {
	cart, err := GetCart(cartID)
	if err != nil {
		return err
	}
	cart.Calculate()
	return nil
}

// ResetDiscount resets discount
func ResetDiscount(cartID string) error {
	cart, err := GetCart(cartID)
	if err != nil {
		return err
	}

	for i := 0; i < len(cart.Items); i++ {
		if val, ok := ProductsDBMap[cart.Items[i].Product.Code]; ok {
			cart.Items[i].Price = val.Price
			cart.Items[i].Discount = nil
		} else {
			fmt.Println(cart.Items[i].Product.Code + " is invalid product code")
			continue
		}
	}
	return nil
}

// 	cart.ApplyDiscount()
// var total float64
// for i := 0; i < len(cart.Items); i++ {
// 	total += cart.Items[i].Price
// }
// cart.Total = total
// 	return nil
// }
