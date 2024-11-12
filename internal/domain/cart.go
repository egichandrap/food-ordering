package domain

import "time"

type CartItem struct {
	ID       string  `json:"id"`
	MenuID   string  `json:"menu_id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Cart struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Items     []CartItem `json:"items"`
	Total     float64    `json:"total"`
	CreatedAt time.Time  `json:"created_at"`
}
