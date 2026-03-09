package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// sqlc -- fast speed, hard developer friendly
// sqlx -- medium speed, medium developer friendly
// gorm -- slow speed, easy developer friendly, powerful task manager

func GetConnectionString(cnf *config.DBConfig) string {
	connection_str := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
	)
	if !cnf.EnableSSLMode {
		connection_str += " sslmode=disable"
	}
	return connection_str
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
