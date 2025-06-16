package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleWare(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := GetUsernameFromContext(r.Context())
		if !ok || userID == "" {
			http.Error(w, "no userID in downstream request context", http.StatusInternalServerError)
			return
		}
		if userID != "TitleFightLover" {
			http.Error(w, userID, http.StatusInternalServerError)
			w.Write([]byte(userID))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	protected := AuthMiddleWare(dummyHandler)
	validJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJiZjc4NTFlNS0zZGQ0LTRiMDgtYWYyYi1hYjY0YjdmMjZjY2MiLCJ1c2VybmFtZSI6IklMb3ZlVGl0bGVGaWdodCIsImlzcyI6Im1peHRhcGVBUEkiLCJleHAiOjE3NTAwNDE3NTgsImlhdCI6MTc1MDAzODE1OH0.MC8o2mZZHJFTc9lVI_2jS0S2ECsdnLPXG-FZgVUX2qU"
	t.Run("test middleware accepts valid JWT", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+validJWT)
		res := httptest.NewRecorder()

		protected.ServeHTTP(res, req)

		if res.Code == http.StatusUnauthorized {
			t.Errorf("middleware incorrectly rejected valid JWT")
		}

		if res.Code == http.StatusInternalServerError {
			t.Error(res.Body)
			t.Errorf("middleware did not correctly inject userID into downstream context")
		}

		if res.Code != http.StatusOK {
			t.Errorf("NO!")
		}
	})
}
