package resty

import "github.com/go-resty/resty/v2"

var client *resty.Client

func init() {
	client = resty.New()
}

func SetClient(c *resty.Client) {
	client = c
}

func Get(url string) (resp *resty.Response, err error) {
	return client.R().Get(url)
}
