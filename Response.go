package dyflexis

import (
	"encoding/json"

	"github.com/leapforce-libraries/go_dyflexis/types"
)

type Response struct {
	Status struct {
		ServerTime string  `json:"serverTime"`
		Pt         float64 `json:"pt"`
		Qn         int64   `json:"qn"`
		Hnd        string  `json:"hnd"`
		Src        string  `json:"src"`
		Dev        bool    `json:"dev"`
	} `json:"api"`
	Response struct {
		Data    json.RawMessage      `json:"data"`
		Expires types.DateTimeString `json:"expires"`
	} `json:"response"`
	OrderBy string `json:"order_by"`
	Error   string `json:"error"`
}
