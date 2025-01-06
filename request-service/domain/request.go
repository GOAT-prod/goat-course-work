package domain

import "time"

type Request struct {
	Id          int       `json:"id"`
	Status      Status    `json:"status"`
	Type        Type      `json:"type"`
	UpdatedDate time.Time `json:"updated_date"`
	Summary     string    `json:"summary"`
	Item        Product   `json:"item"`
}
