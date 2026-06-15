package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type pagef struct{
	HomePage string
	User   string
	IdPremium bool
	Items []string
}

func handler(w http.ResponseWriter, r *http.Request){
	data := pagef {
		HomePage: "ASCII ART",
		User: "<script> alert('hello ascii')</script> gopher",
		IdPremium: true,
		Items: []string{"hello", "victor", "if i let this down "},

	}

	temp, err := template.ParseFiles("pagef")
	if err != nil {
		fmt.Println("error findind file", err)
		http.Error(w, "error finding file ", http.StatusInternalServerError)
	
		return
	}

	temp.Execute(w, data)
	

}

func main(){
	http.HandleFunc("/", handler)


	http.ListenAndServe(":8080", nil)

}