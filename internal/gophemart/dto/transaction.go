package dto

type Transaction struct {
	Number  string  `json:"order"`
	Amount  float64 `json:"sum"`
	OrderID int
}
