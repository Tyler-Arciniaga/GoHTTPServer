package user

import (
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserStore map[string]UserDB
}

func (s *Service) AppendUser(u UserMini) int {
	_, ok := s.UserStore[u.Username]
	if ok {
		return http.StatusConflict
	} else {
		id := uuid.New().String()
		hashedPassword, e := s.HashPassword(u.Password)
		if e != nil {
			return http.StatusConflict
		}
		NewU := UserDB{id, u.Username, hashedPassword}
		s.UserStore[u.Username] = NewU
		return http.StatusCreated
	}
}

func (s *Service) LoginUser(u UserMini) int {
	userDB, ok := s.UserStore[u.Username]
	if ok {
		if s.ComparePasswordHash(userDB.HashedPassword, u.Password) {
			//eventually return JWT token
			return http.StatusOK
		} else {
			return http.StatusUnauthorized
		}
	} else {
		return http.StatusUnauthorized
	}
}

func (s *Service) HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	return string(hash), err
}

func (s *Service) ComparePasswordHash(h, p string) bool {
	e := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	if e == nil {
		return true
	} else {
		return false
	}
}
