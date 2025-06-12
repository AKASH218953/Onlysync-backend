package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if value, ok := os.LookupEnv("MONGOURI"); ok {
		return value
	}
	return "mongodb://localhost:27017"
}

func EnvPromURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if value, ok := os.LookupEnv("PROMETHEUS_URL"); ok {
		return value
	}
	return "http://portal.greenkwh.net:9090"
}

//Google - YOUR_CLIENT_ID
//Google - YOUR_CLIENT_SECRET

func EnvGoogleClientId() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("GOOGLECLIENTID")
}

func EnvGoogleClientSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("GOOGLECLIENTSECRET")
}

func LoadJWTSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("JWT_SECRET")
}

func Getisskey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("ISSKEY")
}
