package user

import "ecommerce/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, password string) (*domain.User, error)
	MatchUserCredentials(email string, password string) (*domain.User, error)
}
