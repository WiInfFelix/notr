package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)

	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err.Error())
		return err
	}

	return nil
}
