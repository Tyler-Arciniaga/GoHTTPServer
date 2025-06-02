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
	trackIDString := strings.TrimPrefix(r.URL.Path, "/tracks/")
	trackID, err := strconv.Atoi(trackIDString)
	checkErr(err, "Invalid query param, requires an integer id", w)

	track, code := h.Service.FetchTrack(trackID)

	formattedRes, err := json.Marshal(track)
	checkErr(err, "Failure to marshal data into JSON", w)

	w.WriteHeader(code)
	w.Write(formattedRes)

}

func checkErr(e error, m string, w http.ResponseWriter) {
	if e != nil {
		http.Error(w, m, http.StatusInternalServerError)
		return
	}
}
