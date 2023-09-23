package users

import (
	"encoding/json"
	"fmt"
	"github.com/sugrado/tama-server/internal/api"
	"net/http"
	"strconv"
)

var (
	user User
)

func getHandler(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		myArr := []string{"adana", "ankara", "istanbul"}
		fmt.Println(myArr[65])
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			api.ResponseJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		u, err := s.Get(id)
		if err != nil {
			api.ResponseJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		api.ResponseJSON(w, http.StatusOK, u)
	}
}

func postHandler(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&user)
		id, err := s.Create(user)
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong...")
			return
		}

		json.NewEncoder(w).Encode(fmt.Sprintf("User %d successfully created!", id))
	}
}
