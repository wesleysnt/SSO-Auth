package responses

type UserClientLogResponse struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	ClientId uint   `json:"client_id"`
	User     string `json:"user"`
	Client   string `json:"client"`
}
