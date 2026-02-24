package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()
	userHandler := user.NewHandler()
	productHandler := product.NewHandler()

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
