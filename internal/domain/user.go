package domain

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	PhoneNumber int64  `json:"phone_number"`
	Password    string `json:"password"`
}
