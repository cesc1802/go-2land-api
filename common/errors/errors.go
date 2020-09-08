package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	ErrType   gin.ErrorType
	ErrorCode int
	OrgError  error
	Msg       string
}

func (a AppError) Error() string {
	if a.Msg == "" {
		return a.OrgError.Error()
	}
	return a.Msg
}

type Option func(o *AppError)

func WithErrorCode(httpCode int) Option {
	return func(o *AppError) {
		o.ErrorCode = httpCode
	}
}

func WithOriginError(err error) Option {
	return func(o *AppError) {
		o.OrgError = err
	}
}

func WithMessage(msg string) Option {
	return func(o *AppError) {
		o.Msg = msg
	}
}
func NewAppError(opts ...Option) AppError {

	var appError AppError
	for _, o := range opts {
		o(&appError)
	}

	return appError
}

func ErrWithMessage(err error, msg string) AppError {
	return AppError{
		ErrorCode: http.StatusBadRequest,
		ErrType: 400,
		OrgError:  err,
		Msg:       msg,
	}
}

func ErrDataNotFound(err error, msg string) AppError {
	return AppError{
		ErrorCode: http.StatusNotFound,
		ErrType:   404,
		OrgError:  err,
		Msg:       msg,
	}
}

func ErrDb(err error) AppError {
	return AppError{
		ErrorCode: http.StatusInternalServerError,
		ErrType: 500,
		OrgError:  err,
	}
}
