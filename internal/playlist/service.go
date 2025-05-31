package playlist

import (
	"net/http"
)

type Service struct {
	PlaylistStore map[string]Playlist
}

func (s *Service) FetchPlaylistData(name string) (Playlist, int) {
	if name == "Playlist1" {
		return Playlist{"Playlist1", "Tyler", "2016", []Track{}}, http.StatusOK
	}
	if name == "Chill-Vibes" {
		return Playlist{"Chill-Vibes", "Derek", "2020", []Track{}}, http.StatusOK
	}
	return Playlist{"Invalid", "Invalid", "Invalid", []Track{}}, http.StatusNotFound
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
