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
    e.Logger.Fatal(e.Start(":3000"))
}

func Greetings(c echo.Context) error {
    return c.JSON(http.StatusOK, HelloWorld{
        Message: "Hello World",
    })
}
