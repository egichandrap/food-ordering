package domain

type Payment struct {
	ID            string  `json:"id"`
	OrderID       string  `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentStatus string  `json:"payment_status"`
	PaymentMethod string  `json:"payment_method"`
}
