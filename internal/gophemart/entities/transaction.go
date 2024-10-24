package entities

import "time"

type Transaction struct {
	Number      int       `json:"order"`
	Amount      float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
