package main

import (
  "net/http"

  "github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
  e := echo.New()
  e.GET("/hello", Greetings)
	e.GET("/hello/:name", GreetingsWithParams)
	e.GET("/hello-queries", GreetingsWithQuery)
  e.Logger.Fatal(e.Start(":3000"))
}

//http://localhost:3000/hello
//{"message":"Hello World"}
func Greetings(c echo.Context) error {
  return c.JSON(http.StatusOK, HelloWorld{
    Message: "Hello World",
  })
}

//http://localhost:3000/hello/Saumit
//{"message":"Hello World, my name is Saumit"}
func GreetingsWithParams(c echo.Context) error {
  params := c.Param("name")
  return c.JSON(http.StatusOK, HelloWorld{
    Message: "Hello World, my name is " + params,
  })
}

//http://localhost:3000/hello-queries?name=Saumit
//{"message":"Hello World, I'm using queries and my name is Saumit"}
func GreetingsWithQuery(c echo.Context) error {
	query := c.QueryParam("name")
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World, I'm using queries and my name is " + query,
	})
}
