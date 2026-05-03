package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

func RequireAuth(store *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sess, _ := store.Get(r, "stay_session")
			uid, ok := sess.Values["user_id"].(int64)
			if !ok || uid == 0 {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			ctx := context.WithValue(r.Context(), "userID", uid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
