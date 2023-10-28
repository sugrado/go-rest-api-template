package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sugrado/go-rest-api-template/pkg/logger"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Logger().Error(fmt.Sprintf("Error when returning response: %v", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.Logger().Error(fmt.Sprintf("Error when returning response: %v", err))
		return
	}
}
