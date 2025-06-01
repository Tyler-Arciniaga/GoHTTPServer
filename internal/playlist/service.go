package playlist

import (
	"net/http"
)

type Service struct {
	PlaylistStore map[string]Playlist
}

func (s *Service) FetchPlaylistData(name string) (Playlist, int) {
	p, ok := s.PlaylistStore[name]
	if ok {
		return p, http.StatusOK
	} else {
		return Playlist{"Invalid", "Invalid", "Invalid", []Track{}}, http.StatusNotFound
	}
}

func (s *Service) FetchAllPlaylists() (map[string]Playlist, int) {
	if len(s.PlaylistStore) > 1 {
		return s.PlaylistStore, http.StatusOK
	} else {
		return nil, http.StatusBadRequest
	}
}

func (s *Service) StoreNewPlaylist(p Playlist) int {
	_, ok := s.PlaylistStore[p.Name]
	if ok {
		return http.StatusConflict
	} else {
		s.PlaylistStore[p.Name] = p
		return http.StatusCreated
	}

}
