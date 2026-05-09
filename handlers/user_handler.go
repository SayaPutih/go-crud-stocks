package handlers

import (
	"encoding/json"
	"net/http"

	"go-crud/config"
	"go-crud/middleware"
	"go-crud/models"
)

func GetMe(w http.ResponseWriter, r *http.Request) {

	// ambil user_id dari middleware
	userID := r.Context().Value(
		middleware.UserIDKey,
	).(string)

	var user models.User

	query := `
		SELECT
			user_id,
			full_name,
			email,
			phone_number,
			role,
			is_verified,
			created_at
		FROM users
		WHERE user_id = $1
	`

	err := config.DB.QueryRowContext(
		r.Context(),
		query,
		userID,
	).Scan(
		&user.UserID,
		&user.FullName,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.IsVerified,
		&user.CreatedAt,
	)

	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// func GetAllUsers(w http.ResponseWriter, r *http.Request) {

// 	query := `
// 	SELECT user_id, full_name, email, phone_number, password_hash, is_verified, created_at FROM users
// 	`

// 	rows, err := config.DB.QueryContext(r.Context(), query)
// 	if err != nil {
// 		http.Error(w, "DB ERROR : "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	defer rows.Close()

// 	var users []models.User

// 	for rows.Next() {
// 		var selected_users models.User

// 		err := rows.Scan(
// 			&selected_users.ID,
// 			&selected_users.FullName,
// 			&selected_users.Email,
// 			&selected_users.PhoneNumber,
// 			&selected_users.PasswordHash,
// 			&selected_users.IsVerified,
// 			&selected_users.CreatedAt,
// 		)
// 		if err != nil {
// 			http.Error(w, "DB ERROR : "+err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		users = append(users, selected_users)
// 	}

// 	if err := rows.Err(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if len(users) == 0 {
// 		http.Error(w, "No user(s) in the database master Evan but api is live", 404)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	if err := json.NewEncoder(w).Encode(users); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// }

// func RegisterAUser(w http.ResponseWriter, r *http.Request) {
// 	var u models.User

// 	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
// 		return
// 	}

// 	//Check Kalo email udah
// 	query := `
// 		SELECT count(*) from users WHERE email = $1
// 	`

// 	var count int
// 	err := config.DB.QueryRowContext(r.Context(), query, u.Email).Scan(&count)
// 	if err != nil {
// 		http.Error(w, "DB Error : "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	if count > 0 {
// 		http.Error(w, "email Exists master Evan", http.StatusConflict)
// 		return
// 	}

// 	var new_user string
// 	inserQuery := `
// 		INSERT INTO users (
// 		full_name,
// 		email,
// 		phone_number,
// 		password_hash,
// 		is_verified
// 		) VALUES ($1,$2,$3,$4,$5)
// 		RETURNING full_name
// 		`

// 	err = config.DB.QueryRowContext(
// 		r.Context(),
// 		inserQuery,
// 		u.FullName,
// 		u.Email,
// 		u.PhoneNumber,
// 		u.PasswordHash,
// 		u.IsVerified,
// 	).Scan(&new_user)

// 	if err != nil {
// 		http.Error(w, "Insert Error : "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	//Setting Content Type
// 	w.Header().Set("Content-Type", "application/json")

// 	//Header Di Pake gimana?
// 	w.WriteHeader(http.StatusCreated)

// 	//w.Write([]byte("New User Created Named : " + new_user));

// 	//Basic Cara return di .net kek
// 	/*
// 		return(new Message{
// 		})
// 	*/
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "User Created",
// 		"name":    new_user,
// 	})
// }

// func UpdateAUser(w http.ResponseWriter, r *http.Request) {
// 	idParam := chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idParam)

// 	if err != nil {
// 		http.Error(w, "Invalid User ID", http.StatusBadRequest)
// 		return
// 	}

// 	var u requests.UpdateUserRequest

// 	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 		http.Error(w, "Invalid Request Body: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var user_exist int
// 	err = config.DB.QueryRowContext(r.Context(),
// 		"SELECT COUNT(*) FROM users WHERE user_id = $1",
// 		id,
// 	).Scan(&user_exist)
// 	if err != nil {
// 		http.Error(w, "DB Error : "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	if user_exist != 1 {
// 		http.Error(w, "User Does not exist with id ", http.StatusNotFound)
// 		return
// 	}

// 	var taken_id int
// 	err = config.DB.QueryRowContext(r.Context(),
// 		"SELECT COUNT(*) FROM users WHERE email = $1 OR phone_number = $2",
// 		u.Email, u.PhoneNumber,
// 	).Scan(&taken_id)
// 	if err != nil {
// 		http.Error(w, "DB error : ", http.StatusInternalServerError)
// 		return
// 	}
// 	if taken_id == 1 {
// 		http.Error(w, "Email or Phone Taken", http.StatusConflict)
// 		return
// 	}

// 	update_query := `
// 		UPDATE
// 			users
// 		SET
// 			full_name = $2,
// 			email = $3,
// 			phone_number = $4,
// 			is_verified = $5
// 		WHERE
// 			user_id = $1
// 	`

// 	_, err = config.DB.ExecContext(r.Context(),
// 		update_query,
// 		id,
// 		u.FullName,
// 		u.Email,
// 		u.PhoneNumber,
// 		u.IsVerified,
// 	)
// 	if err != nil {
// 		http.Error(w, "DB Error : "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"Message":  "User Updated",
// 		"New Name": u.FullName,
// 	})
// }

// func DeleteAUser(w http.ResponseWriter, r *http.Request) {
// 	var idParam = chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idParam)

// 	// if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
// 	// 	http.Error(w, "Invalide Request body", http.StatusBadRequest)
// 	// 	return
// 	// }

// 	var userExist int
// 	err = config.DB.QueryRowContext(r.Context(),
// 		"SELECT COUNT(*) FROM users WHERE user_id = $1",
// 		id,
// 	).Scan(&userExist)

// 	if err != nil {
// 		http.Error(w, "DB Query Error ", http.StatusInternalServerError)
// 		return
// 	}

// 	if userExist != 1 {
// 		http.Error(w, "Nothing with that id ", http.StatusNotFound)
// 		return
// 	}

// 	delete_query := `
// 		DELETE FROM users
// 		WHERE user_id = $1
// 	`

// 	_, err = config.DB.ExecContext(r.Context(),
// 		delete_query,
// 		id,
// 	)

// 	if err != nil {
// 		http.Error(w, "DB Error : "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"Message": "User Deleted",
// 	})

// }
