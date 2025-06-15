package user

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (s *Service) LoginUser(u UserMini) (int, string) {
	userDB, ok := s.UserStore[u.Username]
	if ok {
		if s.ComparePasswordHash(userDB.HashedPassword, u.Password) {
			//eventually return JWT token
			jwt, e := s.GenerateJWT(userDB)
			if e != nil {
				return http.StatusInternalServerError, jwt
			} else {
				return http.StatusOK, jwt
			}
		} else {
			return http.StatusUnauthorized, ""
		}
	} else {
		return http.StatusUnauthorized, ""
	}
}

func (s *Service) GenerateJWT(u UserDB) (string, error) {
	claims := JWTClaims{
		UserID:   u.UUID,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mixtapeAPI",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sampleSecret := "secretKeyChangeLater"

	tokenString, err := token.SignedString([]byte(sampleSecret))
	return tokenString, err

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
