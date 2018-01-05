package models

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/go-chi/chi"
	"database/sql"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM people")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	people := make([]*Person, 0)
	for rows.Next() {
		pl := new(Person)
		err := rows.Scan(&pl.UserID, &pl.FirstName, &pl.LastName, &pl.Age)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, pl)

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "personID")

	if id == "" {
		http.Error(w, http.StatusText(400), 400)
	}
	row := db.QueryRow("SELECT * FROM people WHERE userID = ?", id)
	person := new(Person)
	err := row.Scan(&person.UserID, &person.FirstName, &person.LastName, &person.Age)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(person)
}