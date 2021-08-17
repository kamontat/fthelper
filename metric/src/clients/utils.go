package clients

import (
	"encoding/json"
	"net/http"
)

func toJson(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
