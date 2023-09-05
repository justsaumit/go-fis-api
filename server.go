package main

import (
  "net/http"

  "github.com/labstack/echo/v4"

  "fmt"
  "github.com/justsaumit/go-fic-api/idgen"
  "github.com/justsaumit/go-fic-api/hasher"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
  e := echo.New()
  e.GET("/hello", Greetings)
	e.GET("/hello/:name", GreetingsWithParams)
	e.GET("/hello-queries", GreetingsWithQuery)
  e.GET("/genid", GenerateIDHandler)
  e.GET("/hasher", hasherHandler)
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })
  e.Logger.Fatal(e.Start(":3000"))
}

func Greetings(c echo.Context) error {
  return c.JSON(http.StatusOK, HelloWorld{
    Message: "Hello World",
  })
}

func GreetingsWithParams(c echo.Context) error {
  params := c.Param("name")
  return c.JSON(http.StatusOK, HelloWorld{
    Message: "Hello World, my name is " + params,
  })
}

func GreetingsWithQuery(c echo.Context) error {
	query := c.QueryParam("name")
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World, I'm using queries and my name is " + query,
	})
}

func GenerateIDHandler(c echo.Context) error {
    id := idgen.GenerateID()
    //Print the generated ID to the console.
    fmt.Println("Generated ID:", id)
    return c.JSON(http.StatusOK, map[string]string{"message": "Generated ID: " + id})
    //  return c.JSON(http.StatusOK, HelloWorld{
    //  Message: "Generated ID: " + id,
    //})
}

func hasherHandler(c echo.Context) error {
  filePaths := []string{
    "./message-orig.txt",
    "./message-copy.txt",
    "./message-modd.txt",
  }
  hashResults := make(map[string]string)

  for _, filePath := range filePaths {
      hash, err := hasher.CalculateBLAKE2Hash(filePath)
      if err != nil {
          return c.String(http.StatusInternalServerError, "Error calculating hash")
      }
      hashResults[filePath] = hash
  }

  response := "BLAKE2b hashes:\n"
  for filePath, hash := range hashResults {
      response += fmt.Sprintf("%s: %s\n", filePath, hash)
  }

  return c.String(http.StatusOK, response)
}
