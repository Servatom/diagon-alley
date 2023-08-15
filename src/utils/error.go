package utils

import "errors"

type ExceptionError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ExceptionError) Error() string {
	return e.Code
}

var ErrUnkown = ExceptionError{
	Code:    "DA-E000",
	Message: "Unknown error",
}

var ProductNotFound = ExceptionError{
	Code:    "DA-E001",
	Message: "Product not found",
}


func HandleError(err error) error {
	switch {
	case errors.Is(err, ProductNotFound):
		return ProductNotFound
	default:
		return ErrUnkown
	}
}