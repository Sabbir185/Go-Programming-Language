package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// dependencies
type Server struct {
	cnf            config.Config
	userHandler    *user.Handler
	productHandler *product.Handler
}

// injection of dependencies for creating server instance
func NewServer(cnf config.Config, userHandler *user.Handler, productHandler *product.Handler) *Server {
	return &Server{
		cnf:            cnf,
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {
	// common middleware
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	// HTTP Multiplexer --> router
	mux := http.NewServeMux()
	wrapMux := manager.WrapMux(mux)

	// register routes
	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Starting server on", addr)
	err := http.ListenAndServe(addr, wrapMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
