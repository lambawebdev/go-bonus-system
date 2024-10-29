package entities

import "time"

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAd time.Time `json:"updated_at"`
	Accrual   float64   `json:"accrual,omitempty"`
}

const (
	StatusNew        = 0
	StatusProcessing = 1
	StatusInvalid    = 2
	StatusProcessed  = 3
)

func TransformStatusToString(status int) string {
	mapStatuses := map[int]string{
		0: "NEW",
		1: "PROCESSING",
		2: "INVALID",
		3: "PROCESSED",
	}

	return mapStatuses[status]
}
