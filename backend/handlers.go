package main

import (
	"encoding/json"
	"net/http"
)

func GetAssets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, category, status FROM assets")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.ID, &asset.Name, &asset.Category, &asset.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		assets = append(assets, asset)
	}

	json.NewEncoder(w).Encode(assets)
}

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	var asset Asset
	json.NewDecoder(r.Body).Decode(&asset)

	query := `INSERT INTO assets (name, category, status) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, asset.Name, asset.Category, asset.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
