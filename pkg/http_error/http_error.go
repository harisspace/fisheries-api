package httpError

import (
	"net/http"
)

type BaseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BadRequest struct {
	BaseError
}

func NewBadRequest() BadRequest {
	errObj := BadRequest{BaseError{Code: http.StatusBadRequest, Message: "Bad request"}}
	return errObj
}

type Conflict struct {
	BaseError
}

func NewConflict() Conflict {
	errObj := Conflict{BaseError{Code: http.StatusConflict, Message: "Conflict"}}
	return errObj
}

type NotFound struct {
	BaseError
}

func NewNotFound() NotFound {
	errObj := NotFound{BaseError{Code: http.StatusNotFound, Message: "Not found"}}
	return errObj
}

type Unauthorized struct {
	BaseError
}

func NewUnathorized() Unauthorized {
	errObj := Unauthorized{BaseError{Code: http.StatusUnauthorized, Message: "Unauthorized"}}
	return errObj
}
