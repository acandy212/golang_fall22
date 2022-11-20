package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

// Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

type JsonResponse struct {
	Value1: string `json:"key1"`
	Value2: string `json:"key2"`
}

func main() {
	
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	jsonResp := JsonResponse{
		Value1: "some Data",
		Value2: "other Data",
	}

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request){
		fet.Fprint(w, "JSON goes here")
	})
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
