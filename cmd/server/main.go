package main

import (
	"log"
	"net/http"

	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/playlist"
	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/tracks"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	PlaylistStore := map[string]playlist.Playlist{
		"Playlist1":   {Name: "Playlist1", Author: "Tyler", Created_at: "2016", Tracks: []playlist.Track{}},
		"Chill-Vibes": {Name: "Chill-Vibes", Author: "Derek", Created_at: "2020", Tracks: []playlist.Track{}},
	}
	playlistService := &playlist.Service{PlaylistStore: PlaylistStore}
	playlistHandler := &playlist.Handler{Service: playlistService}

	TrackStore := map[int]tracks.Track{
		1: {Title: "How Time Stretches", Artist: "Whirr", Album: "Feels Like You", Votes: 1},
		2: {Title: "Mellow", Artist: "Whirr", Album: "Feels Like You", Votes: 1},
	}
	tracksService := &tracks.Service{TrackStore: TrackStore}
	trackshandler := &tracks.Handler{Service: tracksService}

	r.Route("/playlist", func(r chi.Router) {
		r.Get("/{name}", playlistHandler.GetSinglePlaylist)
		r.Get("/", playlistHandler.GetAllPlaylists)
		r.Post("/", playlistHandler.PostPlaylist)
		r.Post("/{name}/tracks", playlistHandler.PostPlaylistTrack)
	})

	r.Route("/tracks", func(r chi.Router) {
		r.Get("/{id}", trackshandler.GetTrack)
		r.Post("/{id}", trackshandler.AddTrackVote)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
