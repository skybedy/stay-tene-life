package services

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"strings"
	"time"

	"stay-tene-life/internal/models"
)

type CardService struct{ DB *sql.DB }

func (s CardService) GenerateToken() (string, error) {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return strings.TrimRight(base64.RawURLEncoding.EncodeToString(b), "="), nil
}

func (s CardService) ListByUser(ctx context.Context, uid int64) ([]models.StayCard, error) {
	rows, err := s.DB.QueryContext(ctx, `SELECT id,user_id,token,accommodation_name,guest_name,subtitle,check_in_at,check_out_at,valid_from,valid_until,delete_after,address,maps_url,entry_type,entry_instructions,keybox_code,wifi_ssid,wifi_password,house_info,contact_name,contact_phone,contact_whatsapp,is_active,created_at,updated_at FROM stay_cards WHERE user_id=? ORDER BY created_at DESC`, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []models.StayCard
	for rows.Next() {
		var c models.StayCard
		rows.Scan(&c.ID, &c.UserID, &c.Token, &c.AccommodationName, &c.GuestName, &c.Subtitle, &c.CheckInAt, &c.CheckOutAt, &c.ValidFrom, &c.ValidUntil, &c.DeleteAfter, &c.Address, &c.MapsURL, &c.EntryType, &c.EntryInstructions, &c.KeyboxCode, &c.WifiSSID, &c.WifiPassword, &c.HouseInfo, &c.ContactName, &c.ContactPhone, &c.ContactWhatsapp, &c.IsActive, &c.CreatedAt, &c.UpdatedAt)
		out = append(out, c)
	}
	return out, rows.Err()
}

func (s CardService) FindByToken(ctx context.Context, token string) (*models.StayCard, error) {
	var c models.StayCard
	err := s.DB.QueryRowContext(ctx, `SELECT id,user_id,token,accommodation_name,guest_name,subtitle,check_in_at,check_out_at,valid_from,valid_until,delete_after,address,maps_url,entry_type,entry_instructions,keybox_code,wifi_ssid,wifi_password,house_info,contact_name,contact_phone,contact_whatsapp,is_active,created_at,updated_at FROM stay_cards WHERE token=?`, token).Scan(&c.ID, &c.UserID, &c.Token, &c.AccommodationName, &c.GuestName, &c.Subtitle, &c.CheckInAt, &c.CheckOutAt, &c.ValidFrom, &c.ValidUntil, &c.DeleteAfter, &c.Address, &c.MapsURL, &c.EntryType, &c.EntryInstructions, &c.KeyboxCode, &c.WifiSSID, &c.WifiPassword, &c.HouseInfo, &c.ContactName, &c.ContactPhone, &c.ContactWhatsapp, &c.IsActive, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s CardService) CleanupExpired(ctx context.Context) (int64, error) {
	res, err := s.DB.ExecContext(ctx, `DELETE FROM stay_cards WHERE delete_after < ?`, time.Now().UTC())
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
