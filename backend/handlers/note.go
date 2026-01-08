package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/db"
	"backend/models"

	"github.com/gorilla/mux"
)

// GET /notes
func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.DB.Query("SELECT id, title FROM notes ORDER BY id DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []models.Note

	for rows.Next() {
		var note models.Note
		rows.Scan(&note.ID, &note.Title)
		notes = append(notes, note)
	}

	json.NewEncoder(w).Encode(notes)
}

// POST /notes
func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = db.DB.QueryRow(
		"INSERT INTO notes(title) VALUES($1) RETURNING id",
		note.Title,
	).Scan(&note.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(note)
}

// DELETE /notes/{id}
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("DELETE FROM notes WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Note deleted",
	})
}
