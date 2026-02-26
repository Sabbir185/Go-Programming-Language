package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// sqlc -- fast speed, hard developer friendly
// sqlx -- medium speed, medium developer friendly
// gorm -- slow speed, easy developer friendly, powerful task manager

func GetConnectionString() string {
	return "host=localhost port=5432 user=postgres password=postgres dbname=ecommercego sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
