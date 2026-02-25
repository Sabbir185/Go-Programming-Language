package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := user.NewHandler(cnf, userRepo)
	productHandler := product.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
