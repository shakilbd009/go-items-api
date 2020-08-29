package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err rest_errors.RestErr) {
	RespondJSON(w, err.Status(), err)
}
