package playlist

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPlaylistHandler_CoreLogic(t *testing.T) {
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
