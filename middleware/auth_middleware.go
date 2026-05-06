package middleware

// import (
// 	"context"
// 	"net/http"
// 	"strings"

// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtKey = []byte("SUPER_SECRET_KEY")

// type contextKey string

// const UserContextKey contextKey = "user"

// func JWTMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request)){
// 		authHeader := r.Header.Get("Authorization")

// 		if authHeader == ""{
// 			http.Error(w,"Error No Authentication Header",http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := string.Replace(authHeader,"Bearer ","",1)
// 		token,err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error){
// 			return jwtKey,nil
// 		})

// 		if err != nil || !token.Valid{
// 			http.Error(w,"Invalid Token",http.StatusUnauthorized)
// 		}

// 		ctx := context.WithValue(
// 			r.Context(),
// 			USerContextKey,
// 			token
// 		)

// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}
// }
