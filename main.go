package main

import (
	"errors"
	"fmt"
)

var ProductDB map[string]Product

func init() {
	ProductDB = map[string]Product{
		"CH1": Product{Name: "Chai", Code: "CH1", Price: 3.11},
		"CF1": Product{Name: "Coffee", Code: "CF1", Price: 11.23},
		"AP1": Product{Name: "Apples", Code: "AP1", Price: 6.00},
		"MK1": Product{Name: "Milk", Code: "MK1", Price: 4.75},
	}
}

func (c *Cart) ApplyPriceItem(product Product, qty int, price float64, rule Rule) {
	fmt.Println("In ApplyPriceItem", qty)
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Product.Code == product.Code {
			c.Items[i].Price = price
			c.Items[i].Discount = Discount{
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

func (c *Cart) AddFreeItem(product Product, qty int, rule Rule) {
	fmt.Println("In AddFreeItem", qty)
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].Product.Code == product.Code {
			c.Items[i].Price = 0
			c.Items[i].Discount = Discount{
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
		freeItem := CartItem{
			product,
			0,
			Discount{
				Name:  rule.Name,
				Price: -product.Price,
			},
		}
		c.Items = append(c.Items, freeItem)
		c.CartMap[product.Code]++
		qty--
	}
}

// let's keep it simple for now

func (r Rule) ApplyRule(cart *Cart) error {
	fmt.Println("In ApplyRule")
	fmt.Println(cart)
	fmt.Println(r.Action.ProductCode)
	fmt.Println(r.Action.Operator)
	fmt.Println(r.Action.Qty)

	if qty, ok := cart.CartMap[r.Action.ProductCode]; ok {
		fmt.Println(qty)
		fmt.Println(ok)
		switch {
		case r.Action.Operator == "ge" && qty >= r.Action.Qty:
			fmt.Println("ApplyRule GreaterEq")
			r.ApplyResult(cart)
		case r.Action.Operator == "le" && qty <= r.Action.Qty:
			fmt.Println("ApplyRule LesserEq")
			r.ApplyResult(cart)
		case r.Action.Operator == "g" && qty > r.Action.Qty:
			fmt.Println("ApplyRule Lesser")
			r.ApplyResult(cart)
		case r.Action.Operator == "l" && qty > r.Action.Qty:
			fmt.Println("ApplyRule Lesser")
			r.ApplyResult(cart)
		case r.Action.Operator == "eq" && qty == r.Action.Qty:
			fmt.Println("ApplyRule Equal")
			r.ApplyResult(cart)
		default:
			fmt.Println("Invalid case", r.Action.Operator)
			return errors.New("Invalid case " + r.Action.Operator)
		}
	}
	return nil
}

func (r Rule) ApplyResult(cart *Cart) {
	fmt.Println("In ApplyResult")
	addQty := r.Result.Qty
	addProduct := r.Result.ProductCode
	if addQty < 0 {
		// Unlimited case
		sourceqty, _ := cart.CartMap[r.Action.ProductCode]
		addQty = sourceqty
	}
	if r.Rtype == "getFree" {
		if r.Action.ProductCode == r.Result.ProductCode {
			// If it is a BOGO offer
			addQty = addQty / 2
		}
		fmt.Println("addQty", addQty)
		cart.AddFreeItem(ProductDB[addProduct], addQty, r)
		// r.CalculateCart(cart) recalculate cartMap and price
		return
	}

	// ApplyPrice case
	appliePrice := r.Result.AppliedPrice
	var price float64
	if r.Result.AppliedPriceType == "percent" {
		price = ProductDB[r.Result.ProductCode].Price * appliePrice / 100
	} else {
		price = appliePrice
	}
	cart.ApplyPriceItem(ProductDB[addProduct], addQty, price, r)
}

func main() {
	cart := Cart{
		Items: []CartItem{
			CartItem{Product{Name: "Chai", Code: "CH1", Price: 3.11}, 3.11, Discount{}},
			CartItem{Product{Name: "Apples", Code: "AP1", Price: 6.00}, 6.00, Discount{}},
			CartItem{Product{Name: "Apples", Code: "AP1", Price: 6.00}, 6.00, Discount{}},
			CartItem{Product{Name: "Apples", Code: "AP1", Price: 6.00}, 6.00, Discount{}},
			CartItem{Product{Name: "Milk", Code: "MK1", Price: 4.75}, 4.75, Discount{}},
		},
		Total: 0,
		CartMap: CartMap{
			"CH1": 1, "AP1": 3, "MK1": 1,
		},
	}
	applyDiscount(cart)
}

func applyDiscount(cart Cart) {
	action1 := Action{
		ID:          1,
		RuleID:      1,
		ProductCode: "CH1",
		Operator:    "ge",
		Qty:         1,
	}

	result1 := Result{
		ID:          1,
		RuleID:      1,
		ProductCode: "MK1",
		Qty:         1,
	}

	rule1 := Rule{
		ID:     1,
		Rtype:  "getFree",
		Name:   "CHMK",
		Action: action1,
		Result: result1,
	}
	fmt.Println(rule1)

	action2 := Action{
		ID:          2,
		RuleID:      2,
		ProductCode: "AP1",
		Operator:    "ge",
		Qty:         3,
	}

	result2 := Result{
		ID:               2,
		RuleID:           2,
		ProductCode:      "AP1",
		Qty:              -1,
		ResultType:       "applyPrice",
		AppliedPrice:     4.5,
		AppliedPriceType: "absolute",
	}

	rule2 := Rule{
		ID:     2,
		Rtype:  "applyPrice",
		Name:   "APPL",
		Action: action2,
		Result: result2,
	}
	fmt.Println(rule1)
	fmt.Println(rule2)

	// for loop for rules
	err := rule1.ApplyRule(&cart)
	if err != nil {
		fmt.Println(err)
	}

	err = rule2.ApplyRule(&cart)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cart)
}
