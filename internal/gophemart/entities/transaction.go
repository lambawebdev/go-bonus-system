package entities

import "time"

type Transaction struct {
	Number      int       `json:"order"`
	Amount      int       `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}
