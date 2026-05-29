package main

import (
	"html/template"
	"net/http"
)

func backend(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("victor.html")
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", backend)
	http.ListenAndServe(":5000", nil)
}
