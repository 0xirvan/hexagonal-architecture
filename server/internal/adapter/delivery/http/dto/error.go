package dto

// ErrorResponse represents a standardized error response payload
// ref: https://tools.ietf.org/html/rfc7807
type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}
