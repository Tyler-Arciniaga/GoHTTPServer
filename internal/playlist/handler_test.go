package playlist

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPlaylistHandler_GET(t *testing.T) {
	store := map[string]Playlist{
		"Playlist1":   {"Playlist1", "Tyler", "2016", []Track{}},
		"Chill-Vibes": {"Chill-Vibes", "Derek", "2020", []Track{}},
	}
	TestService := &Service{PlaylistStore: store}
	TestHandler := &Handler{Service: TestService}
	t.Run("test good get request and response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/playlist/Playlist1", nil)
		response := httptest.NewRecorder()

		TestHandler.GetSinglePlaylist(response, request)

		var PlaylistRes Playlist
		json.Unmarshal(response.Body.Bytes(), &PlaylistRes)
		//gotBody := response.Body
		wantBody := store["Playlist1"]

		gotCode := response.Code
		wantCode := 200

		AssertEqualPlaylists(t, PlaylistRes, wantBody)
		CheckStatusCodes(t, gotCode, wantCode)
	})

	t.Run("test good get request on different req", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/playlist/Chill-Vibes", nil)
		response := httptest.NewRecorder()

		TestHandler.GetSinglePlaylist(response, request)

		var PlaylistRes Playlist
		json.Unmarshal(response.Body.Bytes(), &PlaylistRes)
		//gotBody := response.Body.String()
		wantBody := store["Chill-Vibes"]

		gotCode := response.Code
		wantCode := 200

		AssertEqualPlaylists(t, PlaylistRes, wantBody)
		CheckStatusCodes(t, gotCode, wantCode)
	})

	t.Run("test bad get request and 404 response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/playlist/Random", nil)
		response := httptest.NewRecorder()

		TestHandler.GetSinglePlaylist(response, request)

		gotCode := response.Code
		wantCode := 404
		CheckStatusCodes(t, gotCode, wantCode)
	})

	t.Run("test get request for all playlists in store", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/playlist", nil)
		response := httptest.NewRecorder()

		TestHandler.GetAllPlaylists(response, request)

		got := response.Code
		want := http.StatusOK

		wantBody, _ := json.Marshal(TestService.PlaylistStore)
		gotBody := response.Body.Bytes()

		CheckStatusCodes(t, got, want)
		if !(bytes.EqualFold(wantBody, gotBody)) {
			GenericErrorLog(t, got, want)
		}
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

		TestHandler.PostPlaylist(response, request)

		got := response.Code
		want := http.StatusCreated

		CheckStatusCodes(t, got, want)
	})

	t.Run("test post method correct stores a new playlist object", func(t *testing.T) {
		newPlaylist := Playlist{"Playlist3", "Dad", "2023", []Track{}}
		bodyBytes, _ := json.Marshal(newPlaylist)
		body := bytes.NewReader(bodyBytes)
		request, _ := http.NewRequest(http.MethodPost, "/playlist", body)
		response := httptest.NewRecorder()

		TestHandler.PostPlaylist(response, request)

		CheckStatusCodes(t, response.Code, http.StatusCreated)

		if len(TestService.PlaylistStore) != 2 {
			t.Errorf("Playlist was not stored: want %d, got %d", 2, len(TestService.PlaylistStore))
		}
		AssertEqualPlaylists(t, TestService.PlaylistStore["Playlist3"], newPlaylist)
	})

	t.Run("test post method failing to store a duplicate playlist object", func(t *testing.T) {
		newPlaylist := Playlist{"Playlist4", "Bodhi", "2021", []Track{}}
		bodyBytes, _ := json.Marshal(newPlaylist)
		body := bytes.NewReader(bodyBytes)

		firstReq, _ := http.NewRequest(http.MethodPost, "/playlist", body)
		response1 := httptest.NewRecorder()
		TestHandler.PostPlaylist(response1, firstReq)

		if len(TestService.PlaylistStore) != 3 {
			t.Errorf("Playlist was not stored: want %d, got %d", 3, len(TestService.PlaylistStore))
		}
		AssertEqualPlaylists(t, TestService.PlaylistStore["Playlist4"], newPlaylist)

		body = bytes.NewReader(bodyBytes)
		request2, _ := http.NewRequest(http.MethodPost, "/playlist", body)
		response2 := httptest.NewRecorder()

		TestHandler.PostPlaylist(response2, request2)

		if len(TestService.PlaylistStore) != 3 {
			t.Errorf("Duplicates both stored: want %d, got %d", 3, len(TestService.PlaylistStore))
		}

		CheckStatusCodes(t, response1.Code, http.StatusCreated)
		CheckStatusCodes(t, response2.Code, http.StatusConflict)

	})
}
func TestPlaylistHandler_POST_Tracks(t *testing.T) {
	store := map[string]Playlist{
		"ShoeGaze":    {"ShoeGaze", "Tyler", "2016", []Track{}},
		"Chill-Vibes": {"Chill-Vibes", "Derek", "2020", []Track{}},
	}
	TestService := &Service{PlaylistStore: store}
	TestHandler := &Handler{Service: TestService}

	TestSong := Track{"Vividly", "Whirr", "Feels Like You", 1}

	t.Run("test adding songs to existing playlist", func(t *testing.T) {
		bodyBytes, _ := json.Marshal(TestSong)
		body := bytes.NewReader(bodyBytes)
		request, _ := http.NewRequest(http.MethodGet, "/playlist/ShoeGaze/tracks", body)
		response := httptest.NewRecorder()

		TestHandler.PostPlaylistTrack(response, request)

		gotCode := response.Code
		wantCode := http.StatusCreated

		gotRes := TestService.PlaylistStore["ShoeGaze"].Tracks[0]
		wantRes := TestSong

		CheckStatusCodes(t, gotCode, wantCode)
		if !(reflect.DeepEqual(gotRes, wantRes)) {
			GenericErrorLog(t, gotRes, wantRes)
		}
	})

	t.Run("test adding song to non existent playlist", func(t *testing.T) {
		bodyBytes, _ := json.Marshal(TestSong)
		body := bytes.NewReader(bodyBytes)
		request, _ := http.NewRequest(http.MethodGet, "/playlist/NotHere/tracks", body)
		response := httptest.NewRecorder()

		TestHandler.PostPlaylistTrack(response, request)

		gotCode := response.Code
		wantCode := http.StatusBadRequest

		CheckStatusCodes(t, gotCode, wantCode)
	})
}

func AssertEqualPlaylists(t *testing.T, res1, res2 Playlist) {
	if reflect.DeepEqual(res1, res2) == false {
		t.Errorf("got %q, want %q", res1, res2)
	}
}
func CheckStatusCodes(t *testing.T, code1, code2 int) {
	if code1 != code2 {
		t.Errorf("got %d, want %d", code1, code2)
	}
}
func GenericErrorLog(t *testing.T, got, want any) {
	t.Errorf("got %q, want %q", got, want)
}
