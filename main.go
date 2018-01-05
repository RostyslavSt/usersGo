package main

import (
	"usersCRUD/models"
	"github.com/go-chi/chi"
	"net/http"
	"fmt"
)
const port = ":3000"

func main() {
	models.Hello()
	models.InitDB("root:root@tcp(:3306)/people")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(".root"))
	})
	r.Get("/people", models.GetPeople)
	r.Get("/people/{personID}", models.GetPerson)
	r.Post("/people", models.CreatePerson)
	r.Put("/people/{personID}", models.UpdatePerson)
	r.Delete("/people/{personID}", models.DeletePerson)

	fmt.Printf("Serv on port %v....\n", port)
	http.ListenAndServe(port, r)
}

