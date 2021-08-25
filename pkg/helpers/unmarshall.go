package helpers

import (
	"encoding/json"
	"net/http"
)

func UnmarshallBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
