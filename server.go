package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/justsaumit/go-fis-api/handlers"
	"github.com/labstack/echo/v4"
)

const (
	defaultPort = "3000"
)

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Println("No port specified; running on the default port 3000")
	}

	e := echo.New()
	e.POST("/upload", handlers.AddHash)
	e.POST("/verify", handlers.VerifyHash)

	environment := os.Getenv("ENVIRONMENT")
	switch environment {
	case "production":
		startProductionServer(e, port)
	case "development":
		startDevelopmentServer(e, port)
	default:
		log.Printf("Unknown environment '%s', running without TLS encryption\n", environment)
		e.Logger.Fatal(e.Start(":" + port))
	}
}

func startProductionServer(e *echo.Echo, port string) {
	certPath := os.Getenv("CERTPATH")
	keyPath := os.Getenv("KEYPATH")

	if certPath == "" || keyPath == "" {
		log.Fatal("Certificate or key path not provided for production environment")
	}

	log.Printf("Starting production server on port %s\n", port)
	e.Logger.Fatal(e.StartTLS(":"+port, certPath, keyPath))
}

func startDevelopmentServer(e *echo.Echo, port string) {
	log.Printf("Starting development server on port %s\n", port)
	e.Logger.Fatal(e.Start(":" + port))
}
