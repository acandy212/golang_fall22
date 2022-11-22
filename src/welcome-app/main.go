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
	Value1 string `json:"FirstName"`
	Value2 string `json:"LastName"`
	JsonNested JsonNested `json:"User-Information"`
}

type JsonNested struct{
	AddyValue1 string `json:"Address"`
	ContactValue2 string `json:"Contact-Information"`}

func main() {
	
	welcome := Welcome{"New User", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	
	nested := JsonNested{
		AddyValue1: "4364 Elk Creek Road Gainesville, GA 30501",
		ContactValue2: "JohnW1966@hotmail.com, 678-908-1519",
	}
	jsonResp := JsonResponse{
		Value1: "John",
		Value2: "Williams",
		JsonNested: nested,
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
	http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode(jsonResp)
	})
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
