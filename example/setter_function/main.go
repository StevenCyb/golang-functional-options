package main

import (
	"fmt"
	"net/http"
)

type ILogger interface{}

type Client struct {
	baseURL    string
	header     map[string]string
	logger     ILogger
	baseClient *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     map[string]string{},
		baseClient: &http.Client{},
	}
}

func (c *Client) SetHeader(header map[string]string) *Client {
	c.header = header
	return c
}

func (c *Client) SetLogger(logger ILogger) *Client {
	c.logger = logger
	return c
}

func main() {
	client := New("https://api.example.com").
		SetHeader(map[string]string{"Authorization": "Bearer token"}).
		SetLogger(nil)

	fmt.Printf("Client: %+v\n", client)
}
