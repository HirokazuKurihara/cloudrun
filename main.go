package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Choose port
	port := "8080"
	if port == "" {
		port = "8080"
		e.Logger.Infof("Defaulting to port %s", port)
	}
	e.Logger.Infof("Listening on port %s", port)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}