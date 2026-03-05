package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, password string) (*User, error)
	MatchUserCredentials(email, password string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(u User) (*User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
		VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
		RETURNING id
	`
	var userId int
	rows, err := r.db.NamedQuery(query, u)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return nil, err
		}
	}
	u.ID = userId
	return &u, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	var user User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) MatchUserCredentials(email, password string) (*User, error) {
	var user User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
