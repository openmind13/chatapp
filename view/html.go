package view

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	templateFolder = "templates/"
)

var (
	indexTemplate  *template.Template
	loginTemplate  *template.Template
	signupTemplate *template.Template
	chatTemplate   *template.Template
	errorTemplate  *template.Template
)

func init() {
	indexTemplate = template.Must(template.ParseFiles(
		templateFolder+"index.html",
		templateFolder+"header.html"))

	loginTemplate = template.Must(template.ParseFiles(
		templateFolder+"login.html",
		templateFolder+"header.html"))

	signupTemplate = template.Must(template.ParseFiles(
		templateFolder+"signup.html",
		templateFolder+"header.html"))

	chatTemplate = template.Must(template.ParseFiles(
		templateFolder+"chat.html",
		templateFolder+"header.html"))

	errorTemplate = template.Must(template.ParseFiles(
		templateFolder+"error.html",
		templateFolder+"header.html",
	))
}

// ShowIndexPage ...
func ShowIndexPage(w http.ResponseWriter, data interface{}) {
	err := indexTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in ShowIndexPage()")
	}
}

// ShowLoginPage ...
func ShowLoginPage(w http.ResponseWriter, data interface{}) {
	err := loginTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in ShowLoginPage()")
	}
}

// ShowSignupPage ...
func ShowSignupPage(w http.ResponseWriter, data interface{}) {
	err := signupTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in ShowSignupPage()")
	}
}

// ShowChatPage ...
func ShowChatPage(w http.ResponseWriter, data interface{}) {
	err := chatTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in ShowChatPage()")
	}
}

// ShowErrorPage ...
func ShowErrorPage(w http.ResponseWriter, data interface{}) {
	err := errorTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in ShowErrorPage()")
	}
}
