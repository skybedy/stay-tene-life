package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"stay-tene-life/internal/auth"
	"stay-tene-life/internal/config"
	"stay-tene-life/internal/db"
	"stay-tene-life/internal/handlers"
	mw "stay-tene-life/internal/middleware"
	"stay-tene-life/internal/services"
)

func main() {
	cfg := config.Load()
	database, err := db.Open(cfg.DBDSN)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Migrate(context.Background(), database, "migrations"); err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.ParseGlob("web/templates/**/*.gohtml"))
	cards := services.CardService{DB: database}
	h := handlers.Handler{T: t, Cards: cards, BaseURL: cfg.BaseURL}
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))

	go func() {
		ticker := time.NewTicker(cfg.CleanupEvery)
		for range ticker.C {
			n, err := cards.CleanupExpired(context.Background())
			if err == nil && n > 0 {
				log.Printf("cleanup deleted=%d", n)
			}
		}
	}()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/login", handlers.LoginPage)
	r.Get("/auth/google", auth.Start("google"))
	r.Get("/auth/apple", auth.Start("apple"))
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Get("/c/{token}", h.PublicCard)
	r.Route("/admin", func(ar chi.Router) {
		ar.Use(mw.RequireAuth(store))
		ar.Get("/", h.Dashboard)
		ar.Get("/cards/new", h.NewCardForm)
		ar.Post("/cards", h.SaveCard(database))
		ar.Post("/cards/{id}/toggle", h.ToggleCard(database))
		ar.Post("/cards/{id}/delete", h.DeleteCard(database))
	})
	log.Printf("listening on %s", cfg.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, r))
}
