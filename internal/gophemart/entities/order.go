package entities

import "time"

type Order struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAd time.Time `json:"updated_at"`
}

const (
	STATUS_NEW        = 0
	STATUS_PROCESSING = 1
	STATUS_INVALID    = 2
	STATUS_PROCESSED  = 3
)

func TransformStatusToString(status int) string {
	mapStatuses := map[int]string{
		0: "STATUS_NEW",
		1: "STATUS_PROCESSING",
		2: "STATUS_INVALID",
		3: "STATUS_PROCESSED",
	}

	return mapStatuses[status]
}
