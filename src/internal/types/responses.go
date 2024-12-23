package types

type ErrorResponse struct {
	Message    string `json:"msg"`
	StatusCode int    `json:"statusCode"`
}
