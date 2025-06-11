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

func (s *Service) IncrementTrackVote(id int) int {
	t, ok := s.TrackStore[id]
	if ok {
		temp := t
		temp.Votes++
		s.TrackStore[id] = temp

		return http.StatusCreated
	} else {
		return http.StatusNotFound
	}
}
