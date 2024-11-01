package middlewares

import (
	"context"
	"net/http"
)

type UserData struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

var UserContext = &struct{}{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// probably after jwt verification
		user := &UserData{
			Id:    1,
			Email: "ramdom@example.com",
		}

		ctx := context.WithValue(r.Context(), UserContext, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
