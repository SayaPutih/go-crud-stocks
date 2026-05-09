package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("SUPER_SECRET_KEY")

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`

	jwt.RegisteredClaims
}

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		//Tipe Auth
		if authHeader == "" {
			http.Error(w, "Error No Authentication Header", http.StatusUnauthorized)
			return
		}

		//Check ada token
		splitted := strings.Split(authHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			http.Error(w, "Invalid Auth Format", http.StatusUnauthorized)
			return
		}

		//Parsing
		tokenString := splitted[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		//Ini Buat apa?
		ctx := context.WithValue(
			r.Context(),
			UserIDKey,
			claims.UserID,
		)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
