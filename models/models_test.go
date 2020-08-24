package models_test

import (
	"testing"

	"github.com/pranav93/RackspaceAssignment/models"
	"github.com/pranav93/RackspaceAssignment/scripts"
)

func init() {
	scripts.CreateRulesDB()
	scripts.CreateProductsDB()
	scripts.CreateCartsDB()
}

func TestCreateCart(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
}

func TestGetCart(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
	cart, err := models.GetCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if cart.ID != cartID {
		t.Fatalf("Returned wrong cart. %s != %s\n", cartID, cart.ID)
	}
}

func TestDeleteCart(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
	err := models.DeleteCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAddToCart(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
	err := models.AddToCart(cartID, "CF1", 2)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = models.AddToCart(cartID, "AP1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	cart, err := models.GetCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(cart.Items) != 3 {
		t.Fatalf("Cart should contain 3 cart items\n")
	}

	if val, ok := cart.CartMap["CF1"]; !ok {
		t.Fatalf("CF1 should exist in cartMap\n")
	} else {
		if val != 2 {
			t.Fatalf("CF1 value should be 2, but it is %d\n", val)
		}
	}

	if val, ok := cart.CartMap["AP1"]; !ok {
		t.Fatalf("AP1 should exist in cartMap\n")
	} else {
		if val != 1 {
			t.Fatalf("AP1 value should be 1, but it is %d\n", val)
		}
	}

	err = models.AddToCart(cartID, "AP1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(cart.Items) != 4 {
		t.Fatalf("Cart should contain 4 cart items\n")
	}
	if val, ok := cart.CartMap["AP1"]; !ok {
		t.Fatalf("AP1 should exist in cartMap\n")
	} else {
		if val != 2 {
			t.Fatalf("AP1 value should be 2, but it is %d\n", val)
		}
	}
}

func TestRemoveFromCart(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
	err := models.AddToCart(cartID, "CF1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = models.AddToCart(cartID, "AP1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	cart, err := models.GetCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	productCode := cart.Items[0].Code
	cartItemID := cart.Items[0].ID
	err = models.RemoveFromCart(cartID, cartItemID)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(cart.Items) != 1 {
		t.Fatalf("Cart should contain 1 cart items\n")
	}
	if val, ok := cart.CartMap[productCode]; !ok {
		t.Fatalf("%s should exist in cartMap with qunatity 0\n", productCode)
	} else {
		if val != 0 {
			t.Fatalf("AP1 value should be 0, but it is %d\n", val)
		}
	}
}

func TestCalculateCart(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}

	err := models.AddToCart(cartID, "CF1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = models.AddToCart(cartID, "AP1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = models.CalculateCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}
	cart, err := models.GetCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if cart.Total != 17.23 {
		t.Fatalf("Cart total is incorrect. Expected %f, but got %f.\n", 17.23, cart.Total)
	}
}

func TestDiscountShouldNotApply(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
	err := models.AddToCart(cartID, "AP1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	cart, err := models.GetCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	cart.ApplyDiscount()
	cart.Calculate()

	if cart.Items[0].Discount != nil {
		t.Fatalf("Cart item %s should not have any discount\n", cart.Items[0].Code)
	}
}

func TestDiscountShouldApplyGetFree(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
	err := models.AddToCart(cartID, "CH1", 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	cart, err := models.GetCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	cart.ApplyDiscount()
	cart.Calculate()

	if len(cart.Items) != 2 {
		t.Fatalf("Cart should contain 2 cart items\n")
	}
	if val, ok := cart.CartMap["MK1"]; ok {
		if val != 1 {
			t.Fatalf("%s quantity should be 1\n", "MK1")
		}
		if cart.Items[1].Code != "MK1" {
			t.Fatalf("Appended cart item should be %s, but is %s\n", "MK1", cart.Items[1].Code)
		}
		if cart.Items[1].Price != 0 {
			t.Fatalf("Appended cart item %s's price should be %d\n", "MK1", 0)
		}
		if cart.Items[1].Discount == nil {
			t.Fatalf("Appended cart item %s's discount should not be nil\n", "MK1")
		}
		if cart.Items[1].Discount.Name != "CHMK" {
			t.Fatalf("Appended cart item %s's discount name should be %s\n", "MK1", "CHMK")
		}
	} else {
		t.Fatalf("%s should exist in cartMap\n", "MK1")
	}

	if cart.Items[0].Discount != nil {
		t.Fatalf("Cart item %s should not have any discount\n", cart.Items[0].Code)
	}
	chaiProduct, ok := models.ProductsDBMap["CH1"]
	if !ok {
		t.Fatalf("CH1 should exist in ProductsDBMap\n")
	}

	if cart.Total != chaiProduct.Price {
		t.Fatalf("Cart total should be chai's price %f, but is %f\n", chaiProduct.Price, cart.Total)
	}
}

func TestDiscountShouldApplyPriceReduction(t *testing.T) {
	cartID := models.CreateCart()
	if cartID == "" {
		t.Fatalf("Cart ID is not generated. %s is invalid.\n", cartID)
	}
	err := models.AddToCart(cartID, "AP1", 3)
	if err != nil {
		t.Fatalf(err.Error())
	}
	cart, err := models.GetCart(cartID)
	if err != nil {
		t.Fatalf(err.Error())
	}

	cart.ApplyDiscount()
	cart.Calculate()

	if len(cart.Items) != 3 {
		t.Fatalf("Cart should contain 3 cart items\n")
	}
	if val, ok := cart.CartMap["AP1"]; ok {
		if val != 3 {
			t.Fatalf("%s quantity should be 3\n", "AP1")
		}
		for _, cartItem := range cart.Items {
			if cartItem.Code != "AP1" {
				t.Fatalf("Cart item should be %s\n", "AP1")
			}
			if cartItem.Price != 4.5 {
				t.Fatalf("Cart item %s's price should be %f, but is %f\n", "AP1", 4.5, cartItem.Price)
			}
			if cartItem.Discount == nil {
				t.Fatalf("Appended cart item %s's discount should not be nil\n", "AP1")
			}
			if cartItem.Discount.Name != "APPL" {
				t.Fatalf("Appended cart item %s's discount name should be %s\n", "AP1", "APPL")
			}
			if cartItem.Discount.Price != -1.5 {
				t.Fatalf("Appended cart item %s's discount price should be %f, but is %f\n", "AP1",
					-1.5, cartItem.Discount.Price)
			}
		}
	} else {
		t.Fatalf("%s should exist in cartMap\n", "AP1")
	}

	if cart.Total != 13.5 {
		t.Fatalf("Cart total should be %f, but is %f\n", 13.5, cart.Total)
	}
}
