package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"stay-tene-life/internal/services"
)

type Handler struct {
	T       *template.Template
	Cards   services.CardService
	BaseURL string
}

func (h Handler) PublicCard(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	card, err := h.Cards.FindByToken(r.Context(), token)
	if err != nil || !card.IsPubliclyVisible(time.Now().UTC()) {
		w.WriteHeader(http.StatusGone)
		h.T.ExecuteTemplate(w, "public_expired.gohtml", nil)
		return
	}
	h.T.ExecuteTemplate(w, "public_card.gohtml", map[string]any{"Card": card})
}

func (h Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("userID").(int64)
	cards, _ := h.Cards.ListByUser(r.Context(), uid)
	h.T.ExecuteTemplate(w, "admin_dashboard.gohtml", map[string]any{"Cards": cards, "BaseURL": h.BaseURL})
}

func (h Handler) NewCardForm(w http.ResponseWriter, r *http.Request) {
	h.T.ExecuteTemplate(w, "admin_form.gohtml", nil)
}

func parseDate(v string) time.Time { t, _ := time.Parse("2006-01-02T15:04", v); return t.UTC() }

func (h Handler) SaveCard(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value("userID").(int64)
		r.ParseForm()
		token, _ := h.Cards.GenerateToken()
		_, _ = db.ExecContext(r.Context(), `INSERT INTO stay_cards (user_id,token,accommodation_name,guest_name,subtitle,check_in_at,check_out_at,valid_from,valid_until,delete_after,address,maps_url,entry_type,entry_instructions,keybox_code,wifi_ssid,wifi_password,house_info,contact_name,contact_phone,contact_whatsapp,is_active) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, uid, token, r.FormValue("accommodation_name"), r.FormValue("guest_name"), r.FormValue("subtitle"), parseDate(r.FormValue("check_in_at")), parseDate(r.FormValue("check_out_at")), parseDate(r.FormValue("valid_from")), parseDate(r.FormValue("valid_until")), parseDate(r.FormValue("delete_after")), r.FormValue("address"), r.FormValue("maps_url"), r.FormValue("entry_type"), r.FormValue("entry_instructions"), r.FormValue("keybox_code"), r.FormValue("wifi_ssid"), r.FormValue("wifi_password"), r.FormValue("house_info"), r.FormValue("contact_name"), r.FormValue("contact_phone"), r.FormValue("contact_whatsapp"), true)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func (h Handler) ToggleCard(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		db.ExecContext(r.Context(), `UPDATE stay_cards SET is_active = NOT is_active WHERE id=?`, id)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
func (h Handler) DeleteCard(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		db.ExecContext(r.Context(), `DELETE FROM stay_cards WHERE id=?`, id)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<a href='/auth/google'>Sign in with Google</a><br><a href='/auth/apple'>Sign in with Apple</a>"))
}
