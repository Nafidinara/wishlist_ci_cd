package utils

import (
	"net/http"

	"go-wishlist-api/constant"

)

func ConvertResponseCode(err error) int {
	switch err {
	case constant.ErrInsertDatabase:
		return http.StatusInternalServerError
	case constant.ErrEmptyInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
