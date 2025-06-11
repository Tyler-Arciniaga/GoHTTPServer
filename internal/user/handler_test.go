package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUser_SignUp(t *testing.T) {
	store := map[string]UserDB{
		"5af22bf4-d38d-4cc4-b325-3c7e34877a51": {"5af22bf4-d38d-4cc4-b325-3c7e34877a51", "Ty", "randomhash"},
	}
	TestService := Service{UserStore: store}
	TestHandler := Handler{Service: &TestService}

	t.Run("test hash password and compare", func(t *testing.T) {
		NewUser := UserMini{"WhirrFan", "secure123"}
		bodyBytes, _ := json.Marshal(NewUser)
		body := bytes.NewReader(bodyBytes)

		req, _ := http.NewRequest(http.MethodPost, "/signup", body)
		res := httptest.NewRecorder()

		TestHandler.CreateUser(res, req)

		got := res.Code
		want := http.StatusCreated

		CheckStatusCodes(t, got, want)

		if len(store) != 2 {
			t.Errorf("User was not added to user store")
		}

		//now test compare of password
		addedUser := store["WhirrFan"]

		resFromComparePassword := TestService.ComparePasswordHash(addedUser.HashedPassword, NewUser.Password)
		if resFromComparePassword != true {
			t.Errorf("Raw password and hashed password in user store do not match\n")
		}

		resFromComparePassword = TestService.ComparePasswordHash(addedUser.HashedPassword, "wrongpassword123")
		if resFromComparePassword == true {
			t.Errorf("Incorrectly succeeded password comparison")
		}
	})

	t.Run("test login route", func(t *testing.T) {
		LoginUser := UserMini{"WhirrFan", "secure123"}
		bodyBytes, _ := json.Marshal(LoginUser)
		body := bytes.NewReader(bodyBytes)
		req, _ := http.NewRequest(http.MethodPost, "/login", body)
		res := httptest.NewRecorder()

		TestHandler.LoginUser(res, req)

		got := res.Code
		want := http.StatusOK

		CheckStatusCodes(t, got, want)
	})

	t.Run("test failed login", func(t *testing.T) {
		FailedLogin := UserMini{"WhirrFan", "wrongpassword"}
		bodyBytes, _ := json.Marshal(FailedLogin)
		body := bytes.NewReader(bodyBytes)
		req, _ := http.NewRequest(http.MethodPost, "/login", body)
		res := httptest.NewRecorder()

		TestHandler.LoginUser(res, req)

		got := res.Code
		want := http.StatusUnauthorized

		CheckStatusCodes(t, got, want)
	})
}

func CheckStatusCodes(t *testing.T, code1, code2 int) {
	if code1 != code2 {
		t.Errorf("got %d, want %d", code1, code2)
	}
}
