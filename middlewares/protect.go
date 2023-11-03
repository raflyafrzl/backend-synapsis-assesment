package middlewares

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"synapsis-test-be/utilities"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string = r.Header.Get("authorization")

		if tokenString == "" {
			utilities.ErrorResponseWeb(errors.New("Invalid token provided"), 401)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// Handle unsupported signing method
				return nil, errors.New("Unsupported signing method")
			}

			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			utilities.ErrorResponseWeb(errors.New("Invalid token provided"), 401)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var parentContext context.Context = context.Background()
			var ctx context.Context = context.WithValue(parentContext, "auth", claims)

			next.ServeHTTP(w, r.WithContext(ctx))
			return
		} else {
			utilities.ErrorResponseWeb(errors.New("Invalid token provided"), 401)
			return
		}
	})
}
