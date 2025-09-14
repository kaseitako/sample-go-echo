package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "sample-go-echo/docs"
)

// HealthResponse represents the health check response
// @Schema HealthResponse
type HealthResponse struct {
	Status string `json:"status" example:"OK"`
}

// @title Sample Go Echo API
// @version 1.0
// @description This is a sample server using Echo framework
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/hello", hello)

	// Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// hello godoc
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /hello [get]
func hello(c echo.Context) error {
	response := HealthResponse{
		Status: "OK",
	}
	return c.JSON(http.StatusOK, response)
}