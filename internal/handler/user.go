package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stainton/user-management/internal/models"
	"github.com/stainton/user-management/pkg/middleware"
)

var db *sql.DB

func RegisterUserManager(connectedDB *sql.DB, router *gin.Engine) {
	db = connectedDB

	router.POST("/register", func(ctx *gin.Context) {
		RegisterHandler(ctx.Writer, ctx.Request)
	})
	router.POST("/login", func(ctx *gin.Context) {
		LoginHandler(ctx.Writer, ctx.Request)
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.BaseUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	gotUser, err := models.GetUserByUsername(db, user.Username)
	if err == nil || gotUser.Username == user.Username {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	if err := models.CreateUser(db, user); err != nil {
		http.Error(w, "Couldn't create user", http.StatusInternalServerError)
		return
	}

	token, _ := middleware.GenerateJWT(user.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.BaseUser
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		b, _ := json.Marshal(creds)
		http.Error(w, "Invalid input "+err.Error()+string(b), http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByUsername(db, creds.Username)
	if err != nil || !models.CheckPasswordHash(creds.Password, user.Password) {
		http.Error(w, "Invalid credentials "+err.Error(), http.StatusUnauthorized)
		return
	}

	token, _ := middleware.GenerateJWT(user.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
