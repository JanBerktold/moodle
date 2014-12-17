package moodle

// https://github.com/PuerkitoBio/goquery
import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"net/url"
)

func validateLogin(body io.Reader) (err error) {
	
}

func (client *Client) Login(username, password string) (err error) {

	values := url.Values{}
	values.Set("username", username)
	values.Set("password", password)

	req, _ := http.NewRequest("POST", client.BaseDomain+LoginURL, bytes.NewBufferString(values.Encode()))
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if err := AssertHttpRequest(resp); err != nil {
		return err
	}

	defer resp.Body.Close()
	return validateLogin(resp.Body)
}

func (client *Client) Logout() (err error) {
	return nil
}
