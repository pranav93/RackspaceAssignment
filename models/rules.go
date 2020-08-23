package models

import (
	"errors"
	"fmt"
)

// Rule is a struct for storing rules data
type Rule struct {
	ID     int
	Rtype  string
	Name   string
	Action Action
	Result Result
}

// Action is a struct for storing actions for a rule
type Action struct {
	ID          int
	RuleID      int
	ProductCode string
	Operator    string
	Qty         int
}

// Result is a struct for storing results for a rule
type Result struct {
	ID               int
	RuleID           int
	ProductCode      string
	Qty              int
	AppliedPrice     float64
	AppliedPriceType string
	ResultType       string
}

// ApplyRule applies the rule
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
		// No need to repeat these switch statements
		case r.Action.Operator == "ge" && qty >= r.Action.Qty:
			fmt.Println("ApplyRule GreaterEq")
			err := r.ApplyResult(cart)
			return err
		case r.Action.Operator == "le" && qty <= r.Action.Qty:
			fmt.Println("ApplyRule LesserEq")
			err := r.ApplyResult(cart)
			return err
		case r.Action.Operator == "g" && qty > r.Action.Qty:
			fmt.Println("ApplyRule Lesser")
			err := r.ApplyResult(cart)
			return err
		case r.Action.Operator == "l" && qty > r.Action.Qty:
			fmt.Println("ApplyRule Lesser")
			err := r.ApplyResult(cart)
			return err
		case r.Action.Operator == "eq" && qty == r.Action.Qty:
			fmt.Println("ApplyRule Equal")
			err := r.ApplyResult(cart)
			return err
		default:
			fmt.Println("Invalid case", r.Action.Operator)
			return errors.New("Invalid case " + r.Action.Operator)
		}
	}
	return nil
}

// ApplyResult applies the result for a rule
func (r Rule) ApplyResult(cart *Cart) error {
	fmt.Println("In ApplyResult")
	addQty := r.Result.Qty
	addedProductID := r.Result.ProductCode
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

		var product Product
		if val, ok := ProductsDBMap[addedProductID]; ok {
			product = val
		} else {
			return errors.New("Invalid product id " + addedProductID)
		}

		cart.AddFreeItem(product, addQty, r)
		// r.CalculateCart(cart) recalculate cartMap and price
		return nil
	}

	// ApplyPrice case
	appliePrice := r.Result.AppliedPrice
	var price float64
	if r.Result.AppliedPriceType == "percent" {
		price = ProductsDBMap[r.Result.ProductCode].Price * appliePrice / 100
	} else {
		price = appliePrice
	}
	cart.ApplyPriceItem(ProductsDBMap[addedProductID], addQty, price, r)
	return nil
}
