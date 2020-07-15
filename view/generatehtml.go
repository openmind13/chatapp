package view

import (
	"fmt"
	"html/template"
	"net/http"
)

// GenerateHTML ...
func GenerateHTML(w http.ResponseWriter, data interface{}, templateNames ...string) {
	var files []string

	for _, file := range templateNames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	tmpl := template.Must(template.ParseFiles(files...))
	err := tmpl.ExecuteTemplate(w, templateNames[0], data)
	if err != nil {
		fmt.Println("error in executing tempates", err)
	}
}
