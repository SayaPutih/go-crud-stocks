package handlers

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"go-crud/config"
// 	"net/http"
// 	"time"

// 	"go-crud/models"

// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"
// )

// var jwtKey = []byte("SUPER_SECRET_KEY")

// type Claims struct {
// 	UserID int    `json:"user_id"`
// 	Email  string `json:"email"`
// 	jwt.RegisteredClaims
// }

// func Register(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Wrong Request", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var u models.User

// 	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 		http.Error(w, "Invalid Body", http.StatusBadRequest)
// 		return
// 	}

// 	var count int
// 	err := config.DB.QueryRowContext(
// 		r.Context(),
// 		"SELECT COUNT(*) FROM users WHERE email = $1",
// 		u.Email,
// 	).Scan(&count)
// 	if err != nil {
// 		http.Error(w, "Something wrong when checking for similar email", http.StatusInternalServerError)
// 		return
// 	}
// 	if count > 0 {
// 		http.Error(w, "That Email already exists in the database ", http.StatusConflict)
// 		return
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword(
// 		[]byte(u.PasswordHash),
// 		bcrypt.DefaultCost,
// 	)
// 	if err != nil {
// 		http.Error(w, "Hash Error", http.StatusInternalServerError)
// 		return
// 	}

// 	var userID int

// 	err = config.DB.QueryRowContext(
// 		r.Context(),
// 		`
// 		INSERT INTO users(
// 			full_name,
// 			email,
// 			phone_number,
// 			password_hash,
// 			is_verified
// 		)VALUES ($1,$2,$3,$4,$5)
// 		RETURNING user_id
// 		`,
// 		u.FullName,
// 		u.Email,
// 		u.PhoneNumber,
// 		string(hashedPassword),
// 		false,
// 	).Scan(&userID)

// 	if err != nil {
// 		http.Error(w, "Db Insert Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message": "Registered user",
// 		"user_id": userID,
// 	})

// }

// func Login(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Wrong Request", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var u models.User

// 	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 		http.Error(w, "Invalid Body", http.StatusBadRequest)
// 		return
// 	}

// 	//var exists int
// 	var validPassword string
// 	err := config.DB.QueryRowContext(
// 		r.Context(),
// 		"SELECT user_id,password_hash FROM users WHERE email = $1",
// 		u.Email,
// 	).Scan(&u.ID, &validPassword)
// 	if err == sql.ErrNoRows {
// 		http.Error(w, "Invalid Email Or Password", http.StatusNotFound)
// 		return
// 	}
// 	if err != nil {
// 		http.Error(w, "Didnt Get Password From Db", http.StatusInternalServerError)
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(validPassword), []byte(u.PasswordHash))
// 	if err != nil {
// 		http.Error(w, "Invalide Email Or passowrd", http.StatusUnauthorized)
// 		return
// 	}

// 	expirationTime := time.Now().Add(24 * time.Hour)
// 	claims := &Claims{
// 		UserID: u.ID,
// 		Email:  u.Email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)

// 	if err != nil {
// 		http.Error(w, "Token Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message": "Login Success",
// 		"token":   tokenString,
// 		"user": map[string]interface{}{
// 			"user_id": u.ID,
// 			"email":   u.Email,
// 		},
// 	})

// }
