package playlist

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	Service *Service
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetPlaylist(w, r)
	}
}

func (h *Handler) GetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistName := strings.TrimPrefix(r.URL.Path, "/playlist/")
	fetchedPlaylist, code := h.Service.FetchPlaylistData(playlistName)
	w.WriteHeader(code)
	formattedRes, _ := json.Marshal(fetchedPlaylist)
	w.Write(formattedRes)
}
