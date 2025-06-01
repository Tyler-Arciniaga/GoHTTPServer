package main

import (
	"log"
	"net/http"

	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/playlist"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	store := map[string]playlist.Playlist{
		"Playlist1":   {Name: "Playlist1", Author: "Tyler", Created_at: "2016", Tracks: []playlist.Track{}},
		"Chill-Vibes": {Name: "Chill-Vibes", Author: "Derek", Created_at: "2020", Tracks: []playlist.Track{}},
	}
	playlistService := &playlist.Service{PlaylistStore: store}
	playlistHandler := &playlist.Handler{Service: playlistService}

	r.Route("/playlist", func(r chi.Router) {
		r.Get("/{name}", playlistHandler.GetSinglePlaylist)
		r.Get("/", playlistHandler.GetAllPlaylists)
		r.Post("/", playlistHandler.PostPlaylist)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
