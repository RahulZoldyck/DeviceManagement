package models

import "time"

type Device struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"created_at"`
}
