package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"sample-go-echo/database"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

func CreateUser(req CreateUserRequest) (*User, error) {
	query := psql.Insert("users").
		Columns("name").
		Values(req.Name)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	result, err := database.DB.Exec(sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetUserByID(int(lastInsertID))
}

func GetUserByID(userID int) (*User, error) {
	query := psql.Select("user_id", "name", "created_at").
		From("users").
		Where(squirrel.Eq{"user_id": userID})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var user User
	err = database.DB.QueryRow(sqlQuery, args...).Scan(&user.UserID, &user.Name, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	query := psql.Select("user_id", "name", "created_at").
		From("users").
		OrderBy("created_at DESC")

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := database.DB.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Name, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(userID int, req UpdateUserRequest) (*User, error) {
	query := psql.Update("users").
		Set("name", req.Name).
		Where(squirrel.Eq{"user_id": userID})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	result, err := database.DB.Exec(sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	return GetUserByID(userID)
}

func DeleteUser(userID int) error {
	query := psql.Delete("users").
		Where(squirrel.Eq{"user_id": userID})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	result, err := database.DB.Exec(sqlQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}