package resty

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func init() {
	client = resty.New()
}

func SetClient(c *resty.Client) {
	client = c
}

func GetHTTPClient() *http.Client {
	return client.GetClient()
}

func Get(url string) (resp *resty.Response, err error) {
	return client.R().Get(url)
}
