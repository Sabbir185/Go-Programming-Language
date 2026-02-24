package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()
	middleware := middleware.NewMiddlewares(cnf)

	userHandler := user.NewHandler()
	productHandler := product.NewHandler(middleware)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
