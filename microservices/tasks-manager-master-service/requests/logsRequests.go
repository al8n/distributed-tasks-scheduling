package requests

type GetLogsRequest struct {
	Start int64 `json:"start,omitempty"`
	After int64 `json:"after,omitempty"`
	Page int32 `json:"page"`
	Limit int32 `json:"limit"`
	Field string `json:"field,omitempty"`
}
