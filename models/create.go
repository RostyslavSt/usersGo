package models

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(person)

	//check for empty values
	if person.LastName == "" || person.FirstName == "" || person.Age == 0 {
		fmt.Println("unvalid json")
		http.Error(w, http.StatusText(400), 400)
		return
	}
	result, err := db.Exec("INSERT INTO people (FirstName, LastName, Age) VALUES(?, ?, ?)",
		person.FirstName, person.LastName, person.Age)
	if err != nil {
		fmt.Println("error with EXEC")
		http.Error(w, http.StatusText(500), 500)
		return
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "Person %v created successfully (%d row affected)\n", person.UserID, rowsAffected)
}
