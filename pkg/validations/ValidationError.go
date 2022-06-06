package validations

type ValidationErrors struct {
	Key       string `json:"key"`
	ErrorType string `json:"error_type"`
}
