package middlewares

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"synapsis-test-be/model"
	"synapsis-test-be/utilities"
)

func AuthMiddleware(han http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string = r.Header.Get("authorization")

		if tokenString == "" {
			utilities.ErrorResponseWeb(errors.New("Invalid token provided"), 401)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			}

			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			utilities.ErrorResponseWeb(errors.New("Invalid token provided"), 401)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			var parentContext context.Context = context.Background()

			var ctx context.Context = context.WithValue(parentContext, "auth", claims)

			requestCtx := r.WithContext(ctx)
			han.ServeHTTP(w, requestCtx)

		} else {

			panic(model.ResponseFailWeb{
				Status:     "Failed",
				StatusCode: 401,
				Error:      err.Error(),
			})
		}

	})

}
