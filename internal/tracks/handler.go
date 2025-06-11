package tracks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	Service *Service
}

func (h *Handler) GetTrack(w http.ResponseWriter, r *http.Request) {
	trackID, err := extractTrackID(r.URL.Path)
	checkErr(err, "Invalid query param, requires an integer id", w)
	//call service's FetchTrack() to get back track and status code
	track, code := h.Service.FetchTrack(trackID)

	//marshal fetched track data into JSON
	formattedRes, err := json.Marshal(track)
	checkErr(err, "Failure to marshal data into JSON", w)

	//write information back into response writer
	w.WriteHeader(code)
	w.Write(formattedRes)

}

func (h *Handler) AddTrackVote(w http.ResponseWriter, r *http.Request) {
	trackID, err := extractTrackID(r.URL.Path)
	checkErr(err, "Invalid query param, requires an integer id", w)

	code := h.Service.IncrementTrackVote(trackID)

	w.WriteHeader(code)
}

func extractTrackID(url_path string) (int, error) {
	//extract track id from URL param and convert to int
	trackIDString := strings.TrimPrefix(url_path, "/tracks/")
	trackID, err := strconv.Atoi(trackIDString)
	return trackID, err
}

func checkErr(e error, m string, w http.ResponseWriter) {
	if e != nil {
		http.Error(w, m, http.StatusInternalServerError)
		return
	}
}
