package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type BaseUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type User struct {
	BaseUser
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(db *sql.DB, user BaseUser) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (username, password, email, role) VALUES (?, ?, ?, ?)", user.Username, hashedPassword, user.Email, user.Role)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (User, error) {
	var user User
	row := db.QueryRow("SELECT id, username, password, email, role, created_at FROM users WHERE username = ?", username)
	var timeStr string
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role, &timeStr)
	if err == sql.ErrNoRows {
		return User{}, nil
	}
	// 将字符串转化为time.Time表示的时间
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", timeStr)
	return user, err
}
