package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    e.GET("/health", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
    })

    e.GET("/service/vessel", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello from vessel endpoint !!")
    })

    e.GET("/service/container", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello from container endpoint !!")
    })

    e.Logger.Fatal(e.Start(":8080"))

}
