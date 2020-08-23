package models

import "fmt"

// Discount Discount
type Discount struct {
	Name  string
	Tag   string
	Price float64
}

// CartItem is a struct for product
type CartItem struct {
	ID string
	Product
	Price    float64
	Discount *Discount
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

// Calculate Calculate
func (c *Cart) Calculate() {
	var total float64
	for i := 0; i < len(c.Items); i++ {
		total += c.Items[i].Price
	}
	c.Total = total
}

// ApplyDiscount ApplyDiscount
func (c *Cart) ApplyDiscount() {
	rules := GetAllRulesDB()

	for ID, rule := range rules {
		fmt.Println("Applying rule with id", ID)
		rule.ApplyRule(c)
	}
}

// ApplyPriceItem applies the provided price for cart items
func (c *Cart) ApplyPriceItem(product Product, qty int, price float64, rule Rule) {
	fmt.Println("In ApplyPriceItem", qty)
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Product.Code == product.Code {
			c.Items[i].Price = price
			c.Items[i].Discount = &Discount{
				Name:  rule.Name,
				Price: price - product.Price,
			}
			qty--
			if qty == 0 {
				break
			}
		}
	}
}

// AddFreeItem adds or updates the items for getFree rules
func (c *Cart) AddFreeItem(product Product, qty int, rule Rule) {
	fmt.Println("In AddFreeItem", qty)
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Product.Code == product.Code {
			c.Items[i].Price = 0
			c.Items[i].Discount = &Discount{
				Name:  rule.Name,
				Price: -product.Price,
			}
			qty--
			if qty == 0 {
				break
			}
		}
	}
	for qty != 0 {
		freeItem := CreateCartItem(product.Code)
		freeItem.Price = 0
		freeItem.Discount = &Discount{
			Name:  rule.Name,
			Price: -product.Price,
		}
		c.AddItem(freeItem)
		c.CartMap[product.Code]++
		qty--
	}

	fmt.Println("FreeItem", *c.Items[0])
	fmt.Println("FreeItem", *c.Items[1])
	fmt.Println("FreeItem", *c.Items[2])
}

// AddItem AddItem
func (c *Cart) AddItem(item *CartItem) []*CartItem {
	c.Items = append(c.Items, item)
	return c.Items
}

// func (c *Cart) AddItem(cartItem CartItem) []CartItem {
// 	c.Items = append(c.Items, cartItem)
// 	fmt.Println("Added item to cart")
// 	fmt.Println("c.Items", c.Items)
// 	return c.Items
// }
