package scripts

import "github.com/pranav93/RackspaceAssignment/models"

// CreateProductsDB creates rules in in-memory key-val store
func CreateProductsDB() {
	models.ProductsDBMap = map[string]models.Product{
		"CH1": models.Product{Name: "Chai", Code: "CH1", Price: 3.11},
		"CF1": models.Product{Name: "Coffee", Code: "CF1", Price: 11.23},
		"AP1": models.Product{Name: "Apples", Code: "AP1", Price: 6.00},
		"MK1": models.Product{Name: "Milk", Code: "MK1", Price: 4.75},
	}
}
