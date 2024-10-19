package main

import (
	"log"
	"net/http"

	"github.com/stainton/user-management/internal/db"
	"github.com/stainton/user-management/internal/handler"
	"github.com/stainton/user-management/internal/middleware"
)

func main() {
	db.Init()

	http.HandleFunc("/register", handler.RegisterHandler)
	http.HandleFunc("/login", handler.LoginHandler)

	protected := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is a protected route"))
	})
	http.Handle("/protected", middleware.JWTAuthMiddleware(protected))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
