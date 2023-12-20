package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/justsaumit/go-fis-api/handlers"
	"github.com/labstack/echo/v4"
)

const (
	defaultPort           = "3000"
	defaultDomain         = "localhost"
	defaultAPIEndpointURL = "http://localhost:3000"
)

func main() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	domain, apiEndpointUrl := os.Getenv("DOMAIN")
	api_endpoint_url := os.Getenv("API_ENDPOINT_URL")
	if domain == "" || api_endpoint_url == "" {
		log.Println("Warning: DOMAIN and API_ENDPOINT_URL environment variable not set. Using defaults")
		domain = defaultDomain
		api_endpoint_url = defaultAPIEndpointURL
	}

	certPath := "/etc/letsencrypt/live/" + domain + "/fullchain.pem"
	keyPath := "/etc/letsencrypt/live/" + domain + "/privkey.pem"

	e := echo.New()
	e.POST("/upload", handlers.AddHash)
	e.POST("/verify", handlers.VerifyHash)

	environment := os.Getenv("ENVIRONMENT")
	switch environment {
	case "development":
		e.Logger.Fatal(e.Start(":" + port))
	case "production":
		e.Logger.Fatal(e.StartTLS(":"+port, certPath, keyPath))
	default:
		log.Printf("Unknown environment '%s', starting on default port %s\n", environment, port)
		e.Logger.Fatal(e.Start(":" + port))
	}
}
