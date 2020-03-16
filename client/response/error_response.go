package response

type ErrorResponse []struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
