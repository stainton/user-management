package handler

import (
	"encoding/json"
	"net/http"

	"github.com/stainton/user-management/internal/db"
	"github.com/stainton/user-management/internal/middleware"
	"github.com/stainton/user-management/internal/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := models.CreateUser(db.DB, user); err != nil {
		http.Error(w, "Couldn't create user", http.StatusInternalServerError)
		return
	}

	token, _ := middleware.GenerateJWT(user.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByUsername(db.DB, creds.Username)
	if err != nil || !models.CheckPasswordHash(creds.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := middleware.GenerateJWT(user.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
