package moodle

import (
	"errors"
	"net/http"
	"strconv"
)

// Default checking whether http request suceed
// If the request has been successful, then nil is returned
func AssertHttpRequest(resp *http.Response) (err error) {

	if resp.StatusCode != 200 {
		return errors.New(strconv.Itoa(resp.StatusCode) + ": " + http.StatusText(resp.StatusCode))
	}

	return nil
}
