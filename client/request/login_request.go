package request

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
