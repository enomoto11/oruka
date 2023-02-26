package user

import (
	"net/http"
	"oruka/api/utils/error"
)

func CreateGetUserError() error.ApiError {
	var error error.ApiError

	error.Message = "cannot find any users"
	error.Status = http.StatusBadRequest

	return error
}
