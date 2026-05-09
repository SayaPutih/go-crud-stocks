package middleware

import (
	"net/http"

	"go-crud/config"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(UserIDKey).(string)

		if !ok || userID == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var role string
		query := `SELECT role FROM users WHERE user_id = $1`
		err := config.DB.QueryRowContext(
			r.Context(),
			query,
			userID,
		).Scan(&role)

		if err != nil {
			http.Error(w, "Error Getting From Database", http.StatusUnauthorized)
			return
		}

		if role != "admin" {
			http.Error(w, "Admin Only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})

}
