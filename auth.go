package moodle

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func (client *Client) Login(username, password string) (err error) {

	values := url.Values{}
	values.Set("username", username)
	values.Set("password", password)

	req, _ := http.NewRequest("POST", client.BaseDomain+LoginURL, bytes.NewBufferString(values.Encode()))
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.Status != 200 {
		return errors.New(strconv.Itoa(resp.StatusCode) + ": " + http.StatusText(resp.StatusCode))
	}

	return nil
}

func (client *Client) Logout() (err error) {
	return nil
}
