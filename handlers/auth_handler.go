package handlers

import (
	"database/sql"
	"encoding/json"
	"go-crud/config"
	"log"
	"net/http"
	"time"

	"go-crud/models"
	"go-crud/requests"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Wrong Request", http.StatusMethodNotAllowed)
		return
	}

	var req requests.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	log.Printf("%+v\n", req)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" || req.FullName == "" {
		http.Error(w, "Missing Fields ", http.StatusBadRequest)
		return
	}

	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users WHERE email = $1
		)
	`
	err = config.DB.QueryRowContext(
		r.Context(),
		query,
		req.Email,
	).Scan(&exists)
	if err != nil {
		http.Error(w, "Something wrong when checking for similar email", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "That Email already exists in the database ", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		http.Error(w, "Hash Error", http.StatusInternalServerError)
		return
	}

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, "Failed Begin Transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var userID string

	insertQuery := `
		INSERT INTO users(full_name,email,phone_number,password_hash)
		VALUES ($1,$2,$3,$4)
		RETURNING user_id
	`

	err = tx.QueryRow(
		insertQuery,
		req.FullName,
		req.Email,
		req.PhoneNumber,
		string(hashedPassword),
	).Scan(&userID)

	if err != nil {
		http.Error(w, "Failed to Create User", http.StatusInternalServerError)
		return
	}

	insertWalletQuery := `
		INSERT INTO wallet(user_id) VALUES ($1)
	`
	_, err = tx.Exec(
		insertWalletQuery,
		userID,
	)
	if err != nil {
		http.Error(w, "Failed Create Wallet", http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error Commiting Transaction", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Registered user",
		"user_id": userID,
	})

}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`

	jwt.RegisteredClaims
}

var jwtKey = []byte("SUPER_SECRET_KEY")

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Wrong Request", http.StatusMethodNotAllowed)
		return
	}

	var req requests.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Missing Fields", http.StatusInternalServerError)
		return
	}

	var u models.User
	query := `
		SELECT 
			user_id,
			full_name,
			email,
			phone_number,
			password_hash,
			is_verified,
			created_at,
			updated_at,
			role
		FROM users
		WHERE email = $1
	`
	err = config.DB.QueryRowContext(
		r.Context(),
		query,
		req.Email,
	).Scan(
		&u.UserID,
		&u.FullName,
		&u.Email,
		&u.PhoneNumber,
		&u.PasswordHash,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid Email", http.StatusInternalServerError)
			return
		}

		http.Error(w, "Database Error", http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalide Email Or passowrd", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: u.UserID.String(),
		Email:  u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Token Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login Success",
		"token":   tokenString,
		"user": map[string]interface{}{
			"user_id": u.UserID,
			"email":   u.Email,
		},
	})

}
