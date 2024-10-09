package dto

import "time"

type Order struct {
	Number    string    `json:"number"`
	CreatedAt time.Time `json:"created_at"`
}
