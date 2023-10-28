package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sugrado/go-rest-api-template/internal/api"
	"github.com/sugrado/go-rest-api-template/pkg/custom_errors"
)

func getHandler(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			api.ResponseJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		u, err := s.Get(id)
		if err != nil {
			api.HandleError(err, w)
			return
		}
		api.ResponseJSON(w, http.StatusOK, u)
	}
}

func postHandler(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			api.HandleError(&custom_errors.ValidationError{Errors: "JSON Body is not valid"}, w)
			return
		}
		id, err := s.Create(user)
		if err != nil {
			api.HandleError(err, w)
			return
		}

		api.ResponseJSON(w, http.StatusOK, fmt.Sprintf("User %d successfully created!", id))
	}
}
