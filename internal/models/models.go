package models

import "time"

type User struct {
	ID        int64
	Provider  string
	Subject   string
	Email     string
	Name      string
	CreatedAt time.Time
}

type StayCard struct {
	ID                int64
	UserID            int64
	Token             string
	AccommodationName string
	GuestName         string
	Subtitle          string
	CheckInAt         time.Time
	CheckOutAt        time.Time
	ValidFrom         time.Time
	ValidUntil        time.Time
	DeleteAfter       time.Time
	Address           string
	MapsURL           string
	EntryType         string
	EntryInstructions string
	KeyboxCode        string
	WifiSSID          string
	WifiPassword      string
	HouseInfo         string
	ContactName       string
	ContactPhone      string
	ContactWhatsapp   string
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (s StayCard) IsPubliclyVisible(now time.Time) bool {
	return s.IsActive && (now.Equal(s.ValidFrom) || now.After(s.ValidFrom)) && now.Before(s.ValidUntil)
}
