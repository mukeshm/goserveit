package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
	}
}
