package main

type Product struct {
	Name  string
	Code  string
	Price float64
}

type Cart struct {
	Items []CartItem
	Total float64
	CartMap
}

type Discount struct {
	Name  string
	Tag   string
	Price float64
}

type CartItem struct {
	Product
	Price float64
	Discount
}

type CartMap map[string]int

type Rule struct {
	ID     int
	Rtype  string
	Name   string
	Action Action
	Result Result
}

type Action struct {
	ID          int
	RuleID      int
	ProductCode string
	Operator    string
	Qty         int
}

type Result struct {
	ID               int
	RuleID           int
	ProductCode      string
	Qty              int
	AppliedPrice     float64
	AppliedPriceType string
	ResultType       string
}
