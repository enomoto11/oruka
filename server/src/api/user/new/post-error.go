package new

import (
	"net/http"
	"oruka/api/utils/error"
)

func UserAlreadyExistsError() error.ApiError {
	error := error.ApiError{
		Message: "user of the manavis id already exists",
		Status:  http.StatusBadRequest,
	}

	return error
}

func UserRegisterationError() error.ApiError {
	error := error.ApiError{
		Message: "invalid user paramaters",
		Status:  http.StatusBadRequest,
	}

	return error
}
