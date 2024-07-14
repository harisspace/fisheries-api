package utils

import (
	"net/http"

	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/labstack/echo/v4"
)

type Result struct {
	Data  interface{}
	Error interface{}
}

type PaginationMeta struct {
	Page      int `json:"page"`
	Quantity  int `json:"quantity"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

type PaginationResult struct {
	Data  interface{}
	Error interface{}
	Meta  PaginationMeta
}

type ResponseWrapper struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

type PaginationResponseWrapper struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data"`
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Meta    PaginationMeta `json:"meta"`
}

func ResponseSuccess(data interface{}, message string, statusCode int, c echo.Context) error {
	success := false

	// Pretend status code < 400 is success
	if statusCode < http.StatusBadRequest {
		success = true
	}

	result := ResponseWrapper{
		Success: success,
		Data:    data,
		Message: message,
		Code:    statusCode,
	}

	return c.JSON(statusCode, result)
}

func PaginationResponseSuccess(data interface{}, message string, statusCode int, meta PaginationMeta, c echo.Context) error {
	success := false

	// Pretend status code < 400 is success
	if statusCode < http.StatusBadRequest {
		success = true
	}

	result := PaginationResponseWrapper{
		Success: success,
		Data:    data,
		Message: message,
		Code:    statusCode,
		Meta:    meta,
	}

	return c.JSON(statusCode, result)
}

func ResponseError(err interface{}, c echo.Context) error {
	errorObj := getErrorStatusCode(err)

	result := ResponseWrapper{
		Success: false,
		Data:    errorObj.Data,
		Message: errorObj.Message,
		Code:    errorObj.Code,
	}

	return c.JSON(result.Code, result)
}

func getErrorStatusCode(err interface{}) httpError.BaseError {
	errData := httpError.BaseError{}

	switch obj := err.(type) {
	case httpError.BadRequest:
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Conflict:
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Unauthorized:
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.NotFound:
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	default:
		errData.Code = http.StatusConflict
		return errData
	}
}
