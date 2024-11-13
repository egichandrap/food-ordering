package domain

import "time"

type Order struct {
	ID         string    `json:"id"`
	CertID     string    `json:"cert_id"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
