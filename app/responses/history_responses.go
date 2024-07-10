package responses

import (
	"github.com/golang-module/carbon/v2"
)

type HistoryResponse struct {
	Id     uint            `json:"id"`
	User   string          `json:"user"`
	Client string          `json:"client"`
	Date   carbon.DateTime `json:"date"`
}
