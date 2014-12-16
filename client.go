package moodle

import (
	"net/http"
)

type Client struct {
	http.Client
	BaseDomain string
}

func NewClient(baseDomain string) *Client {
	return &Client{
		BaseDomain: baseDomain,
	}
}
