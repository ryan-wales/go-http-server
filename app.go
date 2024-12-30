package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/blog/{year}/{month}/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		year := vars["year"]
		month := vars["month"]
		title := vars["title"]
		fmt.Fprintf(w, "You have requested the blog article %s written in %s %s", title, month, year)
	})

	http.ListenAndServe(":8080", r)
}
