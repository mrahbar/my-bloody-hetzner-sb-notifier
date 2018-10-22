package client

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	crawler := &Client{
		&http.Client{Timeout: 10 * time.Second}	,
	}

	return crawler
}

func (c *Client) DoRequest(url string, target interface{}) error {
	r, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
