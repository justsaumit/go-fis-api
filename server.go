package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"fmt"

	"github.com/justsaumit/go-fic-api/hasher"
	"github.com/justsaumit/go-fic-api/idgen"
)

// type HelloWorld struct {
// 	Message string `json:"message"`
// }

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/hello", Greetings)
	app.Get("/hello/:name", GreetingsWithParams)
	app.Get("/hello-queries", GreetingsWithQueries)
	app.Get("/genid", GenerateIdHandler)
	app.Get("/hasher", hasherHandler)

	log.Fatal(app.Listen(":3000"))
}

func Greetings(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World",
	})
}

func GreetingsWithParams(c *fiber.Ctx) error {
	params := c.Params("name")
	return c.JSON(fiber.Map{
		"message": "Hello World, my name is " + params,
	})
}

func GreetingsWithQueries(c *fiber.Ctx) error {
	queries := c.Query("name")
	return c.JSON(fiber.Map{
		"message": "Hello World, I am using Queries and my name is " + queries,
	})
}

func GenerateIdHandler(c *fiber.Ctx) error {
	id := idgen.GenerateID()
	fmt.Println("Generated ID: " + id)
	return c.JSON(fiber.Map{
		"id": "Generated ID:" + id,
	})
}

func hasherHandler(c *fiber.Ctx) error {
	filePaths := []string{
		"./message-orig.txt",
		"./message-copy.txt",
		"./message-modd.txt",
	}
	hashResults := make(map[string]string)

	for _, filePath := range filePaths {
		hash, err := hasher.CalculateBLAKE2Hash(filePath)
		if err != nil {
			return c.SendString("Error calculating hash")
		}
		hashResults[filePath] = hash
	}

	response := "BLAKE@ Hash Results: \n"
	for filePath, hash := range hashResults {
		response += filePath + ": " + hash + "\n"
	}

	return c.SendString(response)

}
