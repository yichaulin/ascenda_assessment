package client

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	restyClient *resty.Client
}

func New() Client {
	return Client{
		restyClient: resty.New(),
	}
}

func (c *Client) SetTimeout(duration time.Duration) {
	c.restyClient.SetTimeout(duration)
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.restyClient.GetClient()
}

func (c *Client) Get(url string) (resp *resty.Response, err error) {
	return c.restyClient.R().Get(url)
}
