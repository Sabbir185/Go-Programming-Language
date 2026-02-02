package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
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

	port, port_err := strconv.ParseInt(httpPort, 10, 64)
	if port_err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	configuration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	loadConfig()
	return configuration
}
