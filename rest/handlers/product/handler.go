package product

import (
	middleware "ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svr         Service
}

func NewHandler(middlewares *middleware.Middlewares, svr Service) *Handler {
	return &Handler{
		middlewares: middlewares,
		svr:         svr,
	}
}
