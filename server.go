package main

import (
	"github.com/justsaumit/go-fis-api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/upload", handlers.AddHash)
	e.POST("/verify", handlers.VerifyHash)
        //e.Logger.Fatal(e.Start(":3000"))
	e.Logger.Fatal(e.StartTLS(":3000", "/etc/letsencrypt/live/draconyan.xyz/fullchain.pem", "/etc/letsencrypt/live/draconyan.xyz/privkey.pem"))
}
