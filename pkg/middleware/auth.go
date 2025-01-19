package middleware

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte("default-key")

func SetJWTKey(key string) {
	jwtkey = []byte(key)
}

// GenerateJWT generates a JWT token for the given username.
//
// The function takes a username as a parameter and returns a string representing the JWT token and an error if any.
// The JWT token is generated using the HS256 signing method and includes the following claims:
// - Issuer: The username.
// - Subject: "user_token".
// - Audience: "user_api".
// - IssuedAt: The current time.
// - NotBefore: The current time.
// - ExpiresAt: The current time plus 24 hours.
//
// The generated JWT token is signed using the jwtkey byte slice.
//
// If any error occurs during the token generation process, an error will be returned.
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    username,
		Subject:   "user_token",
		Audience:  jwt.ClaimStrings{"user_api"},
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		NotBefore: &jwt.NumericDate{Time: time.Now()},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
	})
	return token.SignedString(jwtkey)
}

// ValidateJWT validates a JWT token.
//
// This function takes a signed JWT token string and attempts to parse and validate it.
// It uses the global jwtkey for verification.
//
// Parameters:
//   - signedString: A string containing the signed JWT token to be validated.
//
// Returns:
//   - An error if the token is invalid or cannot be parsed.
//   - nil if the token is valid.
func ValidateJWT(signedString string) error {
	claim := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(signedString, &claim, func(t *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil || !token.Valid {
		return err
	}
	return nil
}

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		err := ValidateJWT(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
