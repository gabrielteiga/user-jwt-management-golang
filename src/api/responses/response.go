package responses

type DefaultResponse[T any] struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
}

func NewDefaultResponse[T any](message, status string, data T) *DefaultResponse[T] {
	return &DefaultResponse[T]{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

func Success[T any](message string, data T) *DefaultResponse[T] {
	return NewDefaultResponse(message, "success", data)
}

func Error[T any](message string, data T) *DefaultResponse[T] {
	return NewDefaultResponse(message, "error", data)
}
