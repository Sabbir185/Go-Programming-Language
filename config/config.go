package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
	DB          *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the .env file")
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")
	jwtSecret := os.Getenv("JWT_SECRET")

	port, port_err := strconv.ParseInt(httpPort, 10, 64)
	if port_err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	// --- Database Env ---
	db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		fmt.Println("DB Host is required")
		os.Exit(1)
	}

	db_port := os.Getenv("DB_PORT")
	if db_port == "" {
		fmt.Println("db_port is required")
		os.Exit(1)
	}
	dbPortInt, err := strconv.Atoi(db_port)
	if err != nil {
		fmt.Println("db_port must be a number")
		os.Exit(1)
	}

	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		fmt.Println("db_name is required")
		os.Exit(1)
	}
	db_user := os.Getenv("DB_USER")
	if db_user == "" {
		fmt.Println("db_user is required")
		os.Exit(1)
	}
	db_password := os.Getenv("DB_PASSWORD")
	if db_password == "" {
		fmt.Println("db_password is required")
		os.Exit(1)
	}

	db_enable_ssl := os.Getenv("DB_ENABLE_SSL_MODE")
	if db_enable_ssl == "" {
		fmt.Println("db_enable_ssl is required")
		os.Exit(1)
	}

	enable_ssl, err := strconv.ParseBool(db_enable_ssl)

	if err != nil {
		fmt.Println("Invalid enable ssl mode value", err)
		os.Exit(1)
	}
	// --- Database Env End ---

	db_config := &DBConfig{
		Host:          db_host,
		Port:          dbPortInt,
		Name:          db_name,
		User:          db_user,
		Password:      db_password,
		EnableSSLMode: enable_ssl,
	}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JwtSecret:   jwtSecret,
		DB:          db_config,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
