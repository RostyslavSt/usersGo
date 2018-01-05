package models

import (
	"net/http"
	"github.com/go-chi/chi"
	"fmt"
)

func DeletePerson (w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "personID")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
	}
	result, err := db.Exec("DELETE FROM people WHERE userID = ?", id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "Person with ID = %v deleted successfully (%d row affected)\n", id, rowsAffected)
}
