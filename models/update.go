package models

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/go-chi/chi"
	"database/sql"
	"io/ioutil"
)

func UpdatePerson(w http.ResponseWriter, r *http.Request){
	//get ID from request
	id := chi.URLParam(r, "personID")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
	}
	fmt.Println("id - ", id)

	//read item from DB
	var personFromDB Person
	row := db.QueryRow("SELECT * FROM people WHERE userID = ?", id)
	err := row.Scan(&personFromDB.UserID, &personFromDB.FirstName, &personFromDB.LastName, &personFromDB.Age)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Println(personFromDB)


	//read JSON from request body
	body, err := ioutil.ReadAll(r.Body)
	var person Person
	if err != nil {
		log.Println("handlers SaveIDID error:", err)
		http.Error(w, "can’t read body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &person)
	if err != nil {
		log.Println("handlers SaveID error:", err)
		http.Error(w, "can’t Unmarshal json body", http.StatusBadRequest)
		return
	}

	//compare params DB person and JSON person
	if person.FirstName != "" {
		personFromDB.FirstName = person.FirstName
	}
	if person.LastName != "" {
		personFromDB.LastName = person.LastName
	}
	if person.Age != 0 {
		personFromDB.Age = person.Age
	}

	//save to database
	result, err := db.Exec("UPDATE people SET FirstName = ?, LastName = ?, age = ? WHERE userID = ?",
		personFromDB.FirstName, personFromDB.LastName, personFromDB.Age, id)
	if err != nil {
		fmt.Println("error with update")
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "Person with ID = %v updated successfully (%d row affected)\n", id, rowsAffected)
}
