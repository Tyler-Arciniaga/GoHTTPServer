package main

import (
	"log"
	"net/http"

	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/auth"
	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/playlist"
	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/tracks"
	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/user"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	t := "Tyler"
	d := "Derek"
	PlaylistStore := map[string]playlist.Playlist{
		"Playlist1":   {Name: "Playlist1", Author: &t, Created_at: "2016", Tracks: []playlist.Track{}},
		"Chill-Vibes": {Name: "Chill-Vibes", Author: &d, Created_at: "2020", Tracks: []playlist.Track{}},
	}
	playlistService := &playlist.Service{PlaylistStore: PlaylistStore}
	playlistHandler := &playlist.Handler{Service: playlistService}

	TrackStore := map[int]tracks.Track{
		1: {Title: "How Time Stretches", Artist: "Whirr", Album: "Feels Like You", Votes: 1, Voters: map[string]struct{}{}},
		2: {Title: "Mellow", Artist: "Whirr", Album: "Feels Like You", Votes: 1, Voters: map[string]struct{}{}},
	}
	tracksService := &tracks.Service{TrackStore: TrackStore}
	trackshandler := &tracks.Handler{Service: tracksService}

	UserStore := map[string]user.UserDB{
		"TitleFight97": {UUID: "1234", Username: "TitleFight97", HashedPassword: "isitreallyhashed?"},
	}
	userService := &user.Service{UserStore: UserStore}
	userHandler := &user.Handler{Service: userService}

	r.Route("/playlist", func(r chi.Router) {
		r.Get("/{name}", playlistHandler.GetSinglePlaylist)
		r.Get("/", playlistHandler.GetAllPlaylists)
		r.With(auth.AuthMiddleWare).Post("/", playlistHandler.PostPlaylist)
		r.With(auth.AuthMiddleWare).Post("/{name}/tracks", playlistHandler.PostPlaylistTrack)
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/signup", userHandler.CreateUser)
		r.Post("/login", userHandler.LoginUser)
	})

	r.Route("/tracks", func(r chi.Router) {
		r.Get("/{id}", trackshandler.GetTrack)
		r.With(auth.AuthMiddleWare).Post("/{id}", trackshandler.AddTrackVote)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
