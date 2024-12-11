package kafka

import (
	"time"
	"warehouse-service/database/models"
)

type Request struct {
	Id          int         `json:"id"`
	Status      int         `json:"status"`
	Type        int         `json:"type"`
	UpdatedDate time.Time   `json:"updated_date"`
	Summary     string      `json:"summary"`
	Item        RequestItem `json:"item"`
}

type RequestItem struct {
	Id   int            `json:"id"`
	Item models.Product `json:"item"`
}
