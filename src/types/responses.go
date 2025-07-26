package types

// ErrorResp is a standardized error response
type ErrorResp struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SuccessResp is a standardized sucessful response
type SuccessResp struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
