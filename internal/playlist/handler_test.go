package playlist

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPlaylistHandler_GET(t *testing.T) {
	TestService := &Service{}
	TestHandler := &Handler{Service: TestService}
	PlaylistDB := map[string]Playlist{
		"Playlist1":   {"Playlist1", "Tyler", "2016", []Track{}},
		"Chill-Vibes": {"Chill-Vibes", "Derek", "2020", []Track{}},
	}
	t.Run("test good get request and response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/playlist/Playlist1", nil)
		response := httptest.NewRecorder()

		TestHandler.ServeHTTP(response, request)

		var PlaylistRes Playlist
		json.Unmarshal(response.Body.Bytes(), &PlaylistRes)
		//gotBody := response.Body
		wantBody := PlaylistDB["Playlist1"]

		gotCode := response.Code
		wantCode := 200

		CheckResBody(t, PlaylistRes, wantBody)
		CheckStatusCodes(t, gotCode, wantCode)
	})

	t.Run("test good get request on different req", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/playlist/Chill-Vibes", nil)
		response := httptest.NewRecorder()

		TestHandler.ServeHTTP(response, request)

		var PlaylistRes Playlist
		json.Unmarshal(response.Body.Bytes(), &PlaylistRes)
		//gotBody := response.Body.String()
		wantBody := PlaylistDB["Chill-Vibes"]

		gotCode := response.Code
		wantCode := 200

		CheckResBody(t, PlaylistRes, wantBody)
		CheckStatusCodes(t, gotCode, wantCode)
	})

	t.Run("test bad get request and 404 response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/playlist/Random", nil)
		response := httptest.NewRecorder()

		TestHandler.ServeHTTP(response, request)

		got := response.Code
		want := 404
		CheckStatusCodes(t, got, want)
	})
}

func TestPlaylistHandler_Post(t *testing.T) {

	TestService := &Service{make(map[string]Playlist)}
	TestHandler := &Handler{Service: TestService}
	t.Run("test post method returns accepted status code", func(t *testing.T) {
		newPlaylist := `{"name": "Playlist2", "author": "Mom", "created_at": "2022", "tracks" : []}`
		body := strings.NewReader(newPlaylist)
		request, _ := http.NewRequest(http.MethodPost, "/playlist", body)
		response := httptest.NewRecorder()

		TestHandler.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted

		CheckStatusCodes(t, got, want)
	})

	t.Run("test post method correct stores a new playlist object", func(t *testing.T) {
		newPlaylist := `{"name": "Playlist2", "author": "Mom", "created_at": "2022", "tracks" : []}`
		body := strings.NewReader(newPlaylist)
		request, _ := http.NewRequest(http.MethodPost, "/playlist", body)
		response := httptest.NewRecorder()

		TestHandler.ServeHTTP(response, request)

		if len(TestService.PlaylistStore) != 1 {
			t.Errorf("Playlist was not stored: want %d, got %d", 1, len(TestService.PlaylistStore))
		}
	})
}

func CheckResBody(t *testing.T, res1, res2 Playlist) {
	if reflect.DeepEqual(res1, res2) == false {
		t.Errorf("got %q, want %q", res1, res2)
	}
}
func CheckStatusCodes(t *testing.T, code1, code2 int) {
	if code1 != code2 {
		t.Errorf("got %d, want %d", code1, code2)
	}
}
