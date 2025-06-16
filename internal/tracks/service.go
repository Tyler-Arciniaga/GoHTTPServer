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

func (s *Service) IncrementTrackVote(id int, u string) int {
	t, ok := s.TrackStore[id]
	if ok {
		if _, exists := t.Voters[u]; exists {
			return http.StatusBadRequest
		}
		temp := t
		temp.Votes++
		temp.Voters[u] = struct{}{}
		s.TrackStore[id] = temp

		return http.StatusCreated
	} else {
		return http.StatusNotFound
	}
}
