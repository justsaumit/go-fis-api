package handlers

import (
	"database/sql"
	"github.com/justsaumit/go-fis-api/models"
	"github.com/justsaumit/go-fis-api/utils"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "hashstore.db")
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
}

// AddHash handles the file upload, store and response
func AddHash(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Failed to bind request:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	src, err := file.Open()
	if err != nil {
		log.Println("Failed to Open File", err)
		return err
	}
	defer src.Close()

	id, err := utils.GenerateID()
	if err != nil {
		log.Println("Failed to Generate ID", err)
	}

	hash, err := utils.GenerateHash(src)
	if err != nil {
		log.Println("Failed to Generate Hash", err)
	}

	// Store id and hash in the database
	_, err = db.Exec("INSERT INTO hashes (ID, HashValue) VALUES (?, ?)", id, hash)
	if err != nil {
		log.Println("Failed to insert into database:", err)
	}

	data := models.FileHashPair{
		ID:       id,
		FileHash: hash,
	}

	log.Println("Hash added successfully to the database")
	return c.JSON(http.StatusOK, data)
}
