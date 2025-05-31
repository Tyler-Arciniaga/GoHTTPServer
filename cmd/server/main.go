package main

import (
	"log"
	"net/http"

	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/playlist"
)

func main() {
	playlistService := &playlist.Service{}
	playlistHandler := &playlist.Handler{Service: playlistService}
	log.Fatal(http.ListenAndServe(":8080", playlistHandler))
}
