package main

import (
    "fmt"
    "backend/controllers"

    "github.com/labstack/echo/v4"
)

func main() {
    listen := echo.New()
    listen.GET("/info/*", controllers.Info)
    listen.GET("/latest", controllers.Latest)
    fmt.Println("Server Starting ... ")
    listen.Logger.Fatal(listen.Start(":8080"))
}
