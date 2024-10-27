package main

import (
	"html/template"
	"net/http"
)

func main() {
	// Parse the layout and index templates
	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Execute the template with data (optional)
		err := tmpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8000", nil)
}
