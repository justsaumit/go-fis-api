package main

import (
  "net/http"

  "github.com/labstack/echo/v4"

  "fmt"
  "github.com/justsaumit/go-fic-api/idgen"
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
