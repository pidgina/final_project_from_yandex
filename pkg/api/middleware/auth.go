package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"proj/pkg/api/handlers"
	"proj/pkg/api/service"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(pass string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if len(pass) <= 0 {
				next.ServeHTTP(w, r)
				return
			}

			var claims service.Claims

			cookie, err := r.Cookie("token")
			if err != nil {
				handlers.SendErrorJSON(w, "Необходимо авторизоваться", http.StatusUnauthorized)
				return
			}
			jwtToken, err := jwt.ParseWithClaims(cookie.Value, &claims, func(t *jwt.Token) (any, error) {
				return []byte(service.Secret), nil
			})
			if err != nil || !jwtToken.Valid {
				handlers.SendErrorJSON(w, "Необходимо авторизоваться", http.StatusUnauthorized)
				return
			}
			hashedPassword := sha256.Sum256([]byte(pass))
			hashStringPassword := hex.EncodeToString(hashedPassword[:])

			if claims.PasswordHash != hashStringPassword {
				handlers.SendErrorJSON(w, "Необходимо авторизоваться", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)

		})
	}

}
