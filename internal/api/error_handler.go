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
	switch {
	case errors.Is(err, &custom_errors.NotFoundError{}):
		httpStatusCode = http.StatusNotFound
		break
	case errors.Is(err, &custom_errors.BusinessError{}):
		httpStatusCode = http.StatusBadRequest
		break
	case errors.Is(err, &custom_errors.ValidationError{}):
		httpStatusCode = http.StatusBadRequest
		break
	case errors.Is(err, &custom_errors.NotAuthorizedError{}):
		httpStatusCode = http.StatusUnauthorized
		break
	case errors.Is(err, &custom_errors.ForbiddenError{}):
		httpStatusCode = http.StatusForbidden
		break
	default:
		ResponseJSON(w, http.StatusInternalServerError, &errorResponseObject{"Internal Server Error"})
		return
	}

	ResponseJSON(w, httpStatusCode, &errorResponseObject{err.Error()})
}
