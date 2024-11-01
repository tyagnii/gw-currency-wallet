package models

type User struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Currency struct {
	USD float32 `json:"USD"`
	RUB float32 `json:"RUB"`
	EUR float32 `json:"EUR"`
}

type Wallet struct {
	ID      int
	UUID    string
	Message string   `json:"message"`
	Balance Currency `json:"balance"`
}

type Rates struct {
	Rates Currency `json:"rates"`
}
