package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/challenge/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

// ValidateUser checks for a token and validates it
// before allowing the method to execute
func ValidateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse authorization header
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			http.Error(w, "Malformed token", http.StatusUnauthorized)
			return
		}

		// validate token
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(authHeader[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetConfig().JwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
