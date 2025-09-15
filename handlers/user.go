package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"sample-go-echo/models"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided name
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUser(c echo.Context) error {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if req.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}

	user, err := models.CreateUser(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func GetUser(c echo.Context) error {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user")
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users in the system
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetAllUsers(c echo.Context) error {
	users, err := models.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get users")
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update a user's information by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UpdateUserRequest true "Updated user data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c echo.Context) error {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var req models.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if req.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Name is required")
	}

	user, err := models.UpdateUser(userID, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	err = models.DeleteUser(userID)
	if err != nil {
		if err == models.ErrUserNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.NoContent(http.StatusNoContent)
}