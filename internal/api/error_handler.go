package api

import (
	"errors"
	"github.com/sugrado/go-rest-api-template/pkg/custom_errors"
	"net/http"
)

type errorResponseObject struct {
	Message string `json:"message"`
}

func HandleError(err error, w http.ResponseWriter) {
	var httpStatusCode int
	var notFoundError *custom_errors.NotFoundError
	var businessError *custom_errors.BusinessError
	var validationError *custom_errors.ValidationError
	var notAuthorizedError *custom_errors.NotAuthorizedError
	var forbiddenError *custom_errors.ForbiddenError
	switch {
	case errors.As(err, &notFoundError):
		httpStatusCode = http.StatusNotFound
		break
	case errors.As(err, &businessError):
		httpStatusCode = http.StatusBadRequest
		break
	case errors.As(err, &validationError):
		httpStatusCode = http.StatusBadRequest
		break
	case errors.As(err, &notAuthorizedError):
		httpStatusCode = http.StatusUnauthorized
		break
	case errors.As(err, &forbiddenError):
		httpStatusCode = http.StatusForbidden
		break
	default:
		ResponseJSON(w, http.StatusInternalServerError, &errorResponseObject{"Internal Server Error"})
		return
	}

	ResponseJSON(w, httpStatusCode, &errorResponseObject{err.Error()})
}
