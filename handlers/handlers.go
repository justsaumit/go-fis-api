package handlers

import (
	"database/sql"
	"github.com/justsaumit/go-fis-api/models"
	"github.com/justsaumit/go-fis-api/utils"
	"github.com/labstack/echo/v4"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite", "hashstore.db")
	if err != nil {
		log.Fatalf("sqlite3 not installed: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create the 'hashes' table if it does not exist
	_, err = db.Exec(`
            CREATE TABLE IF NOT EXISTS hashes (
            ID TEXT PRIMARY KEY,
            HashValue TEXT NOT NULL
            );
        `)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "API is running!"})
}

// handleError handles errors by logging and returning a JSON response
func handleError(c echo.Context, errMsg string, status int) error {
	log.Println(errMsg)
	return c.JSON(status, map[string]string{"message": errMsg})
}

// AddHash handles the file upload, store and response
func AddHash(c echo.Context) (err error) {
	file, err := c.FormFile("FileInput")
	if err != nil {
		log.Println("Failed to bind request:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		return handleError(c, "Failed to bind request", http.StatusBadRequest)
	}

	// Open the file stream
	src, err := file.Open()
	if err != nil {
		log.Println("Failed to Open File", err)
		return handleError(c, "Failed to Open File", http.StatusInternalServerError)
	}
	// Close the file stream on function exit
	defer src.Close()

	id, err := utils.GenerateID()
	if err != nil {
		log.Println("Failed to Generate ID", err)
		return handleError(c, "Failed to Generate ID", http.StatusInternalServerError)
	}
	log.Printf("ID generated successfully: %s", id)

	hash, err := utils.GenerateHash(src)
	if err != nil {
		log.Println("Failed to Generate Hash", err)
		return handleError(c, "Failed to Generate Hash", http.StatusInternalServerError)
	}
	log.Printf("Hash generated successfully: %s", hash)

	// Store id and hash in the database
	_, err = db.Exec("INSERT INTO hashes (ID, HashValue) VALUES (?, ?)", id, hash)
	if err != nil {
		log.Println("Failed to insert into database:", err)
		return handleError(c, "Failed to insert into database", http.StatusInternalServerError)
	}

	data := models.FileHashPair{
		ID:       id,
		FileHash: hash,
	}

	log.Println("Hash added to the database successfully: %v", data)
	return c.JSON(http.StatusOK, data)
}

// VerifyHash handles the file verification and response

func VerifyHash(c echo.Context) error {
	// Extract the ID from the request
	id := c.FormValue("idInput")

	// Handle file upload
	file, err := c.FormFile("FileInput")
	if err != nil {
		log.Println("Failed to bind request:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// Open the file stream
	src, err := file.Open()
	if err != nil {
		log.Println("Failed to Open File", err)
		return err
	}
	// Close the file stream on function exit
	defer src.Close()

	// Generate hash of the uploaded file
	hash, err := utils.GenerateHash(src)
	if err != nil {
		log.Println("Failed to Generate Hash", err)
	}

	// Retrieve the stored hash from the database
	var storedHash string
	err = db.QueryRow("SELECT HashValue FROM hashes WHERE ID = ?", id).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("ID not Found")
			return c.JSON(http.StatusNotFound, map[string]string{"message": "2"})
		}
		log.Println("Failed to query database:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve stored hash"})
	}

	// Compare the generated hash with the stored hash
	if hash != storedHash {
		log.Println("Hash verification failed")
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "1"})
	}
	log.Println("Hash verified successfully")
	return c.JSON(http.StatusOK, map[string]string{"message": "0"})
}
