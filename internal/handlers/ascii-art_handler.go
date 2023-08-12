package handlers

import (
	asciiart "ascii-art-web/internal/services/ascii_art"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {

	Request := struct {
		SimpleText string `json:"text"`
		Font       string `json:"font"`
	}{}

	type response struct {
		AsciiText string `json:"asciiText"`
	}
	if http.MethodPost == r.Method {
		data, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(data, &Request)
		if err != nil {
			log.Fatal(err)
		}
	}

	asciiMaker := asciiart.NewAsciiArt()
	asciiText, err := asciiMaker.AsciiMake(Request.SimpleText)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	Response := response{
		AsciiText: asciiText,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(Response)
	if err != nil {
		log.Fatal(err)
	}
}
