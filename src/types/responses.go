package types

// ErrorResp is a standardized error response
type ErrorResp struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
