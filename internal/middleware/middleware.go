package middleware

import (
	"net/http"

	"Socialist/internal/session"
)

func LoginOnly(store *session.Storage) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if store.IsNew(r) || store.GetUserID(r) == "" {
				http.Error(w, "Not allowed", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
