package entity

import "time"

// EntityNameEntity — namuna entity. Loyihaga mos nomga o'zgartiring.
type EntityNameEntity struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
