package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server
		MongoDB
	}

	Server struct {
		Host string
		Port string
	}

	MongoDB struct {
		Host string
		Port string
	}
)

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	serverHOST := os.Getenv("SERVER_HOST")
	serverPORT := os.Getenv("SERVER_PORT")
	mongoHOST := os.Getenv("MONGO_HOST")
	mongoPORT := os.Getenv("MONGO_PORT")

	if serverPORT == "" {
		serverPORT = "8080"
	}
	if mongoPORT == "" {
		mongoPORT = "27017"
	}

	if serverHOST == "" {
		serverHOST = "localhost"
	}

	if mongoHOST == "" {
		mongoHOST = "localhost"
	}

	return &Config{
		Server: Server{
			Host: serverHOST,
			Port: serverPORT,
		},
		MongoDB: MongoDB{
			Host: mongoHOST,
			Port: mongoPORT,
		},
	}
}
