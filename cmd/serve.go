package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
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

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	userSvr := user.NewService(userRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := userHandler.NewHandler(cnf, userSvr)
	productHandler := productHandler.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
