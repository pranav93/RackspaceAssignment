package scripts

import "github.com/pranav93/RackspaceAssignment/models"

// CreateCartsDB creates cart entities in in-memory key-val store
func CreateCartsDB() {
	models.CartsDBMap = map[string]*models.Cart{}
	models.CartItemsDBMap = map[string]*models.CartItem{}
}
