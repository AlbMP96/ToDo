package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AlbMP96/backend/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, created_at FROM users WHERE id = $1"
	err := DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, created_at FROM users WHERE email = $1"
	err := DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

var ErrEmailExists = errors.New("email already exists")

func CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	id := uuid.New()

	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`

	_, err := DB.ExecContext(ctx, query, id, user.Name, user.Email, user.Password)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == pq.ErrorCode("23505") {
				return nil, ErrEmailExists
			}
		}
		return nil, err
	}

	newUser := &models.User{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return newUser, nil
}
