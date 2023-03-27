package views

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	Landing = template.Must(template.ParseFiles("resources/global/html/landing.html"))
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
	if err := Landing.Execute(w, nil); err != nil {
		fmt.Printf("VIEWS: %v\n", err)
		return
	}
}
