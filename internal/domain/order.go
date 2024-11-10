package domain

type Order struct {
	ID         string  `json:"id"`
	MenuID     string  `json:"menu_id"`
	Quantity   string  `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
