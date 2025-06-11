package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTracksHandler_GetTrack(t *testing.T) {
	store := map[int]Track{
		1: {"How Time Stretches", "Whirr", "Feels Like You", 1},
		2: {"Mellow", "Whirr", "Feels Like You", 1},
	}
	TestService := Service{TrackStore: store}
	TestHandler := Handler{Service: &TestService}

	t.Run("good staus code on voting for getting a track", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tracks/2", nil)
		res := httptest.NewRecorder()

		TestHandler.GetTrack(res, req)

		got := res.Code
		want := http.StatusOK

		var gotTrack Track
		json.Unmarshal(res.Body.Bytes(), &gotTrack)
		wantTrack := store[2]

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
		if !reflect.DeepEqual(gotTrack, wantTrack) {
			t.Errorf("got %q, want %q", gotTrack, wantTrack)
		}
	})

	t.Run("404 status code when getting non existant track", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/tracks/12", nil)
		res := httptest.NewRecorder()

		TestHandler.GetTrack(res, req)

		got := res.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestTracksHandler_PostVotes(t *testing.T) {
	store := map[int]Track{
		1: {"How Time Stretches", "Whirr", "Feels Like You", 1},
		2: {"Mellow", "Whirr", "Feels Like You", 1},
	}
	TestService := Service{TrackStore: store}
	TestHandler := Handler{Service: &TestService}

	t.Run("increment track's vote when Post /tracks/{id} route is hit", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/tracks/1", nil)
		res := httptest.NewRecorder()

		TestHandler.AddTrackVote(res, req)

		got := TestHandler.Service.TrackStore[1].Votes
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
