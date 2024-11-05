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

type DepositReq struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

type ExchangeReq struct {
	FromCurrency string   `json:"from_currency"`
	ToCurrency   string   `json:"to_currency"`
	Amount       float32  `json:"amount"`
	Rate         Currency `json:"rate"`
}

// WithdrawReq is an alias for Withdrawal request
type WithdrawReq DepositReq

// Rates is an alias for rates requests
type Rates Currency
