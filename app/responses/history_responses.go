package responses

import "time"

type HistoryResponse struct {
	Id     uint      `json:"id"`
	User   string    `json:"user"`
	Client string    `json:"client"`
	Date   time.Time `json:"date"`
}
