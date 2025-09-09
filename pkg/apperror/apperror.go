package apperror

import "fmt"

type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func New(code, message string, status int, err error) *AppError {
	return &AppError{Code: code, Message: message, HTTPStatus: status, Err: err}
}

func BadRequest(message string, err error) *AppError   { return New("bad_request", message, 400, err) }
func Unauthorized(message string, err error) *AppError { return New("unauthorized", message, 401, err) }
func Forbidden(message string, err error) *AppError    { return New("forbidden", message, 403, err) }
func NotFound(message string, err error) *AppError     { return New("not_found", message, 404, err) }
func Conflict(message string, err error) *AppError     { return New("conflict", message, 409, err) }
func Internal(message string, err error) *AppError     { return New("internal_error", message, 500, err) }
