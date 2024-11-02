package blackboxservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/config"
	"github.com/lambawebdev/go-bonus-system/internal/gophemart/entities"
)

const SleepTimeSeconds = 60

var mapStatuses = map[string]int{
	"PROCESSING": entities.StatusProcessing,
	"INVALID":    entities.StatusInvalid,
	"PROCESSED":  entities.StatusProcessed,
}

func FromStringStatusToInt(status string) int {
	return mapStatuses[status]
}

var client = resty.New()

func GetOrderAccrual(number string) (OrderAccrual, error) {
	var order OrderAccrual

	url := fmt.Sprintf("%s/api/orders/%s", config.GetAccrualHost(), number)

	resp, err := client.R().
		SetHeader("Accept", "text/plain").
		Get(url)

	if err != nil {
		return order, err
	}

	if resp.StatusCode() == http.StatusTooManyRequests {
		fmt.Printf("429, Sleep for %d\n seconds", SleepTimeSeconds)
		time.Sleep(SleepTimeSeconds * time.Second)
	}

	if resp.StatusCode() == http.StatusNoContent {
		return order, nil
	}

	err = json.Unmarshal([]byte(resp.String()), &order)

	if err != nil {
		return order, err
	}

	return order, err
}

type OrderAccrual struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
	OrderID int
}
