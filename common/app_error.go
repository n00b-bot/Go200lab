package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)

}

func NewFullErrorResponse(status int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: status,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(e error) *AppError {
	return NewErrorResponse(e, "something went wrong with DB", e.Error(), "DB_ERROR")
}

func ErrInvalidRequest(e error) *AppError {
	return NewErrorResponse(e, "invalid request", e.Error(), "ErrInvalidRequest")
}

func ErrInternal(e error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, e, "something went wrong with Server", e.Error(), "SV_ERROR")
}

func ErrCannotListEntity(entity string, e error) *AppError {
	return NewCustomError(
		e,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}
func ErrCannotDeleteEntity(entity string, e error) *AppError {
	return NewCustomError(
		e,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}
func ErrCannotUpdateEntity(entity string, e error) *AppError {
	return NewCustomError(
		e,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, e error) *AppError {
	return NewCustomError(
		e,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrEntityNotFound(entity string, e error) *AppError {
	return NewCustomError(
		e,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrEntityDeleted(entity string, e error) *AppError {
	return NewCustomError(
		e,
		fmt.Sprintf("%s already deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}
