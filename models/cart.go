package models

import "fmt"

// CartItem is a struct for product
type CartItem struct {
	ID string
	Product
	Price float64
}

// CartMap is a map for ProductCode and Quantity
type CartMap map[string]int

// Cart is a struct for CartItems and related data
type Cart struct {
	ID    string
	Items []*CartItem
	Total float64
	CartMap
}

// ApplyPriceItem applies the provided price for cart items
func (c *Cart) ApplyPriceItem(product Product, qty int, price float64) {
	fmt.Println("In ApplyPriceItem", qty)
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Product.Code == product.Code {
			c.Items[i].Price = price
			qty--
			if qty == 0 {
				break
			}
		}
	}
}

// AddFreeItem adds or updates the items for getFree rules
func (c *Cart) AddFreeItem(product Product, qty int) {
	fmt.Println("In AddFreeItem", qty)
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Product.Code == product.Code {
			c.Items[i].Price = 0
			qty--
			if qty == 0 {
				break
			}
		}
	}
	for qty != 0 {
		freeItem := CreateCartItem(product.Code)
		c.Items = append(c.Items, freeItem)
		c.CartMap[product.Code]++
		qty--
	}
}

// AddItem AddItem
func (box *Cart) AddItem(item *CartItem) []*CartItem {
	box.Items = append(box.Items, item)
	return box.Items
}

// func (c *Cart) AddItem(cartItem CartItem) []CartItem {
// 	c.Items = append(c.Items, cartItem)
// 	fmt.Println("Added item to cart")
// 	fmt.Println("c.Items", c.Items)
// 	return c.Items
// }
