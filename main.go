package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "sample-go-echo/docs"
	"sample-go-echo/database"
	"sample-go-echo/handlers"
)

// HealthResponse represents the health check response
// @Schema HealthResponse
type HealthResponse struct {
	Status string `json:"status" example:"OK"`
}

// ProtectedResponse represents the protected endpoint response
// @Schema ProtectedResponse
type ProtectedResponse struct {
	Message string `json:"message" example:"Access granted"`
	UserID  string `json:"user_id" example:"authenticated_user"`
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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// BearerTokenAuth returns a middleware function that validates Bearer token
func BearerTokenAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get expected token from environment variable
			expectedToken := os.Getenv("BEARER_TOKEN")
			if expectedToken == "" {
				expectedToken = "your-secret-bearer-token" // fallback for development
			}

			// Get Authorization header
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			// Check if it's a Bearer token
			if !strings.HasPrefix(auth, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			// Extract token
			token := strings.TrimPrefix(auth, "Bearer ")

			// Validate token
			if token != expectedToken {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Bearer token")
			}

			return next(c)
		}
	}
}

func main() {
	// Initialize database
	database.InitDB()
	defer database.CloseDB()

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/hello", hello)

	// User CRUD routes - no authentication required
	e.POST("/users", handlers.CreateUser)
	e.GET("/users", handlers.GetAllUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	// Protected routes - require Bearer token
	protected := e.Group("")
	protected.Use(BearerTokenAuth())
	protected.GET("/protected", protectedEndpoint)

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

// protectedEndpoint godoc
// @Summary Protected endpoint
// @Description Returns protected data that requires Bearer token authentication
// @Tags protected
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} ProtectedResponse
// @Failure 401 {object} map[string]string
// @Router /protected [get]
func protectedEndpoint(c echo.Context) error {
	response := ProtectedResponse{
		Message: "Access granted to protected resource",
		UserID:  "authenticated_user",
	}
	return c.JSON(http.StatusOK, response)
}