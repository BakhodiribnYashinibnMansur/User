package serviceerror

import "encoding/json"

type ErrorResponse struct {
	Error            error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

var (
	ErrorNotFound = NewServiceError(nil, "not found", "", "US-000003")
)

func (e *ErrorResponse) ErrorResult() string {
	return e.Message
}

func (e *ErrorResponse) Unwrap() error { return e.Error }

func (e *ErrorResponse) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewServiceError(err error, message, developerMessage, code string) *ErrorResponse {
	return &ErrorResponse{
		Error:            err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func systemError(err error) *ErrorResponse {
	return NewServiceError(err, "internal system error", err.Error(), "US-000000")
}
