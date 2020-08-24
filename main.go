package main

import (
	"github.com/pranav93/RackspaceAssignment/scripts"
	"github.com/pranav93/RackspaceAssignment/setup"
)

func init() {
	scripts.CreateRulesDB()
	scripts.CreateProductsDB()
	scripts.CreateCartsDB()
}

func main() {
	setup.Server().Run()
}
