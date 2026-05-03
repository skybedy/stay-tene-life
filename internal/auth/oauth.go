package auth

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
)

func Start(provider string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, provider+" OAuth not configured yet. Add credentials in .env.", http.StatusNotImplemented)
	}
}
func Callback(db *sql.DB, store *sessions.CookieStore, provider string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, provider+" callback placeholder.", http.StatusNotImplemented)
	}
}
