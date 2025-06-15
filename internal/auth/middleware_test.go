package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleWare(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := GetUserIDFromContext(r.Context())
		if !ok || userID == "" {
			http.Error(w, "no userID in downstream request context", http.StatusInternalServerError)
			return
		}
		if userID != "db2e072d-c5da-44aa-b4e3-ebca8f5674f3" {
			http.Error(w, userID, http.StatusInternalServerError)
			w.Write([]byte(userID))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	protected := AuthMiddleWare(dummyHandler)
	validJWT := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiJkYjJlMDcyZC1jNWRhLTQ0YWEtYjRlMy1lYmNhOGY1Njc0ZjMiLCJ1c2VybmFtZSI6InRpdGxlZmlnaHQiLCJpc3MiOiJtaXh0YXBlQVBJIiwiZXhwIjoxNzQ5OTY4MTU5LCJpYXQiOjE3NDk5NjQ1NTl9.mIlRWbdqApVymZfhSkOS7rKzabxWFvNA92CcVUfzM-s"
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
