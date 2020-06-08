package moodle

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	webService     = "/webservice/rest/server.php"
	defaultService = "moodle_mobile_app"
)

type Client struct {
	wwwroot string
	service string
	token   string
}

type RequestOptions struct {
	Method   string
	Function string
	Data     url.Values
}

func (c Client) Request(o RequestOptions) (*http.Response, error) {
	h := http.DefaultClient

	u := c.wwwroot + webService

	data := o.Data
	data.Add("wstoken", c.token)
	data.Add("wsfunction", o.Function)
	data.Add("moodlewsrestformat", "json")

	var body strings.Reader

	if o.Method == "GET" {
		u = u + "?" + data.Encode()
		data = nil
	}

	if o.Method == "POST" {
		body = *strings.NewReader(data.Encode())
	}

	req, err := http.NewRequest(o.Method, u, &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := h.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewClient(wwwroot, token string) *Client {
	c := &Client{
		wwwroot,
		defaultService,
		token,
	}

	return c
}
