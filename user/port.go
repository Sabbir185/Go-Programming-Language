package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
	MatchUserCredentials(email, password string) (*domain.User, error)
}
