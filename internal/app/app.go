package app

import (
	"ascii-art-web/internal/handlers"
	"net/http"
)

func Run() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.HandleFunc("/", handlers.MainPageHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	http.ListenAndServe("localhost:8080", nil)
}
