package handlers

import (
	"encoding/json"
	"go-crud/config"
	"go-crud/models"
	"go-crud/requests"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	var query = `
		SELECT * FROM stocks;
	`

	rows, err := config.DB.QueryContext(
		r.Context(),
		query,
	)

	if err != nil {
		http.Error(w, "Error Getting Data From Database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allStocks []models.Stock

	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(
			&stock.StockID,
			&stock.Symbol,
			&stock.Name,
			&stock.Price,
			&stock.CreatedAt,
		)

		if err != nil {
			http.Error(w, "Error Giving Data", http.StatusInternalServerError)
			return
		}

		allStocks = append(allStocks, stock)
	}

	if len(allStocks) < 1 {
		http.Error(w, "Nothing in the database master Evan", http.StatusNotFound)
		return
	}

	// 	if variable := something(); condition {
	// 	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"total":  len(allStocks),
		"stocks": allStocks,
	}); err != nil {
		http.Error(w, "Error giving response", http.StatusInternalServerError)
		return
	}
}

func InsertAStock(w http.ResponseWriter, r *http.Request) {

	var s requests.InsertStockRequest
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid Body Request Format", http.StatusBadRequest)
		return
	}

	query := `
		SELECT COUNT(*) from stocks WHERE name = $1 OR symbol = $2
	`

	var exist int
	err = config.DB.QueryRowContext(
		r.Context(),
		query,
		s.Name,
		s.Symbol,
	).Scan(&exist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exist >= 1 {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	query = `
		INSERT INTO stocks(symbol,name,price) VALUES ($1,$2,$3) RETURNING name
	`
	var stockName string
	err = config.DB.QueryRowContext(
		r.Context(),
		query,
		s.Symbol,
		s.Name,
		s.Price,
	).Scan(&stockName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"Message":        "Succesfully Inserted",
		"New Stock Name": stockName,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateAStock(w http.ResponseWriter, r *http.Request) {

	var id = chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Id Must Be valid", http.StatusBadRequest)
		return
	}

	var temp requests.UpdateStockRequest
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var exists int
	err = config.DB.QueryRowContext(r.Context(), "SELECT COUNT(*) FROM stocks WHERE stock_id = $1", id).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exists == 0 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var conflict int
	err = config.DB.QueryRowContext(
		r.Context(),
		"SELECT COUNT(*) FROM stocks WHERE symbol = $1 OR name = $2",
		temp.Symbol,
		temp.Name,
	).Scan(&conflict)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if conflict > 0 {
		http.Error(w, "conflict with other data(s) master evan", http.StatusConflict)
		return
	}

	query := `
		UPDATE
			stocks
		SET
			symbol = $1,
			name = $2,
			price = $3
		WHERE
			stock_id = $4
	`
	_, err = config.DB.ExecContext(
		r.Context(),
		query,
		temp.Symbol,
		temp.Name,
		temp.Price,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "Succesfully Updated " + temp.Name,
		"result":  temp,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// var idParam = chi.URLParam(r, "id")
// id, err := strconv.Atoi(idParam)
func DeleteAStock(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")
	//id, err := strconv.Atoi(id)

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	query := `SELECT COUNT(*) FROM stocks WHERE stock_id = $1`

	var stockExists int
	err := config.DB.QueryRowContext(r.Context(), query, id).Scan(&stockExists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if stockExists == 0 {
		http.Error(w, "stock not found", http.StatusNotFound)
		return
	}

	query = `DELETE FROM stocks WHERE stock_id = $1`
	_, err = config.DB.ExecContext(r.Context(), query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(map[string]interface{}{
		"Message": "Sucess Deletion for id : " + id,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
