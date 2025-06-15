package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Extract entire auth header and check if Bearer is present
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("No authorization header and/or bearer token found in request"))
			return
		}

		//Extract JWT from Bearer in auth header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("No JWT provided"))
			return
		}

		//Parse JWT token using custom JWTClaims type and secret key for JWT signature
		token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (any, error) {
			return []byte("secretKeyChangeLater"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error parsing JWT claims"))
			return
		}

		//Extract JWT claims and check validity of token
		claims, ok := token.Claims.(*JWTClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid JWT"))
			return
		}

		//Inject userID into request context for future handlers/business logic
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		r = r.WithContext(ctx)

		//Call next handler's (business logic) ServeHTTP() w/ new request context
		next.ServeHTTP(w, r)
	})
}
