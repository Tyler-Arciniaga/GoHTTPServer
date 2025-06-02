package tracks

import (
	"net/http"
)

type Service struct {
	TrackStore map[int]Track
}

func (s *Service) FetchTrack(id int) (Track, int) {
	t, ok := s.TrackStore[id]
	if ok {
		return t, http.StatusOK
	} else {
		return Track{}, http.StatusNotFound
	}

}
