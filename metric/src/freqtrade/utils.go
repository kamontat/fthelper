package freqtrade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func buildPath(url *url.URL, version string, path string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("/api/%s/%s", version, path))
}

func toJson(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
