package utils

type (
	StandardResponse struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Error      string      `json:"error,omitempty"`
		Data       interface{} `json:"data,omitempty"`
	}

	ErrorResponse struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
		Error      string `json:"error,omitempty"`
	}
	ValidatorResponse struct {
		StatusCode int      `json:"status_code"`
		Message    string   `json:"message"`
		Error      []string `json:"error,omitempty"`
	}
)
