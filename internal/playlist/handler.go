package playlist

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Tyler-Arciniaga/MixTapeAPI/internal/auth"
)

type Handler struct {
	Service     *Service
	UserService UserService
}

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
	//retrieve map of all playlists in db/store
	PlaylistMap, code := h.Service.FetchAllPlaylists()

	//marhsal playlist map into JSON
	formattedRes, e := json.Marshal(PlaylistMap)
	checkErr(e, "Error marshaling all playlist data as JSON", w)

	//write code and JSON response
	w.WriteHeader(code)
	w.Write(formattedRes)
}

func (h *Handler) PostPlaylist(w http.ResponseWriter, r *http.Request) {
	//decode incoming io.Reader and convert into playlist struct type
	var newPlaylist Playlist
	json.NewDecoder(r.Body).Decode(&newPlaylist)

	username, ok := auth.GetUsernameFromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	newPlaylist.Author = &username

	//return status code in response
	statusCode := h.Service.StoreNewPlaylist(newPlaylist)
	w.WriteHeader(statusCode)
}

func (h *Handler) PostPlaylistTrack(w http.ResponseWriter, r *http.Request) {
	//extract just the playlist name
	playlistName := strings.TrimPrefix(r.URL.Path, "/playlist/")
	playlistName = strings.TrimSuffix(playlistName, "/tracks")

	//decode the new track from json into Track object
	var newTrack Track
	json.NewDecoder(r.Body).Decode(&newTrack)

	//add new track to desired playlist
	statusCode := h.Service.AddNewPlaylistTrack(playlistName, newTrack)

	//return status code
	w.WriteHeader(statusCode)
}

func checkErr(e error, m string, w http.ResponseWriter) {
	if e != nil {
		http.Error(w, m, http.StatusInternalServerError)
		return
	}
}
