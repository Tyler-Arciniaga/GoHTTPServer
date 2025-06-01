package playlist

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handler struct {
	Service *Service
}

/*
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetPlaylist(w, r)
	case http.MethodPost:
		h.PostPlaylist(w, r)
	}
}
*/

func (h *Handler) GetSinglePlaylist(w http.ResponseWriter, r *http.Request) {
	//trim prefix to extract just the playlist name
	playlistName := strings.TrimPrefix(r.URL.Path, "/playlist/")

	//fetch playlist data from services and add code to res header
	fetchedPlaylist, code := h.Service.FetchPlaylistData(playlistName)
	w.WriteHeader(code)

	//marshal fetched playlist data as JSON
	formattedRes, e := json.Marshal(fetchedPlaylist)
	checkErr(e, "Error marshaling single playlist data as JSON", w)

	//write to response writer
	w.Write(formattedRes)
}

func (h *Handler) GetAllPlaylists(w http.ResponseWriter, r *http.Request) {
	PlaylistMap, code := h.Service.FetchAllPlaylists()

	w.WriteHeader(code)
	formattedRes, e := json.Marshal(PlaylistMap)
	checkErr(e, "Error marshaling all playlist data as JSON", w)
	w.Write(formattedRes)
}

func (h *Handler) PostPlaylist(w http.ResponseWriter, r *http.Request) {
	//decode incoming io.Reader and convert into playlist struct type
	var newPlaylist Playlist
	json.NewDecoder(r.Body).Decode(&newPlaylist)

	//return status code in response
	statusCode := h.Service.StoreNewPlaylist(newPlaylist)
	w.WriteHeader(statusCode)
}

func checkErr(e error, m string, w http.ResponseWriter) {
	if e != nil {
		http.Error(w, m, http.StatusInternalServerError)
		return
	}
}
