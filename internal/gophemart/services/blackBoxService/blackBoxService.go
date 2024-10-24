package blackboxservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/config"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

var mapStatuses = map[string]int{
	"STATUS_PROCESSING": entities.STATUS_PROCESSING,
	"STATUS_INVALID":    entities.STATUS_INVALID,
	"STATUS_PROCESSED":  entities.STATUS_PROCESSED,
}

func FromStringStatusToInt(status string) int {
	return mapStatuses[status]
}

var client = resty.New()

func GetOrderStatus(number string) (orderStatusResponse, error) {
	var order orderStatusResponse

	url := fmt.Sprintf("http://%s/api/orders/%s", config.GetAccrualHost(), number)

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return order, err
	}

	if resp.StatusCode() == http.StatusNoContent {
		return order, nil
	}

	err = json.NewDecoder(resp.RawBody()).Decode(&order)

	if err != nil {
		return order, err
	}

	return order, err
}

type orderStatusResponse struct {
	Status string `json:"status"`
}
