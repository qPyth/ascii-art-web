package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mainPage, err := template.ParseFiles("./web/templates/main_page.html")
		if err != nil {
			log.Fatal(err)
		}
		err = mainPage.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
