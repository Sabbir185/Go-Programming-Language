package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to the database")

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := user.NewHandler(cnf, userRepo)
	productHandler := product.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
