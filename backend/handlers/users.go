package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AlbMP96/backend/auth"
	"github.com/AlbMP96/backend/db"
	"github.com/AlbMP96/backend/models"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	ctx := r.Context()
	newUser, err := db.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, db.ErrEmailExists) {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id := query.Get("id")
	email := query.Get("email")

	if id != "" {
		ctx := r.Context()
		user, err := db.GetUserById(ctx, id)
		if err != nil {
			http.Error(w, "Error retrieving user", http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	if email != "" {
		ctx := r.Context()
		user, err := db.GetUserByEmail(ctx, email)
		if err != nil {
			http.Error(w, "Error retrieving user", http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	if id == "" && email == "" {
		http.Error(w, "Either user ID or email is required", http.StatusBadRequest)
		return
	}
	http.Error(w, "Invalid request", http.StatusBadRequest)
}
