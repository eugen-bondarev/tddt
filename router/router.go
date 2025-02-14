package router

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewHttpError(message string, code int) *HttpError {
	return &HttpError{
		Message: message,
		Code:    code,
	}
}

func (e *HttpError) Error() string {
	return e.Message
}

type Ctx interface {
	GetBody(any) error
}

type Handler func(Ctx) (any, error)
