package main

import (
	"fmt"
	"html/template"
	"net/http"

	"workofJoshua/data/about"
	"workofJoshua/data/contact"
	"workofJoshua/data/home"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// render and catch errors
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fs := http.Dir("static")
	//handle css files
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(fs)))

	//Setup Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := home.GetPageData()
		renderTemplate(w, "index", data)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := about.GetPageData()
		renderTemplate(w, "about", data)
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		data := contact.GetPageData()
		renderTemplate(w, "contact", data)
	})

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
