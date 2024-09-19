package utils

type (
	ResponseMeta struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Error      string      `json:"error,omitempty"`
		Data       interface{} `json:"data,omitempty"`
		Meta       Meta        `json:"meta,omitempty"`
	}

	ResponseData struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Error      string      `json:"error,omitempty"`
		Data       interface{} `json:"data,omitempty"`
	}

	ResponseValidator struct {
		StatusCode int      `json:"status_code"`
		Message    string   `json:"message"`
		Error      []string `json:"error,omitempty"`
	}

	Meta struct {
		Pagination Pagination `json:"pagination"`
	}

	Pagination struct {
		Total       int   `json:"total"`
		Count       int   `json:"count"`
		PerPage     int   `json:"per_page"`
		CurrentPage int   `json:"current_page"`
		TotalPages  int   `json:"total_pages"`
		Links       Links `json:"links"`
	}

	Links struct {
		Next string `json:"next"`
	}
)
