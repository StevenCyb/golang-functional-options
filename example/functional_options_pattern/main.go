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

type Option func(*Client)

func New(baseURL string, opts ...Option) *Client {
	client := &Client{
		baseURL:    baseURL,
		header:     map[string]string{},
		baseClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithHeader(header map[string]string) Option {
	return func(c *Client) {
		c.header = header
	}
}

func WithLogger(logger ILogger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func main() {
	client := New("https://api.example.com",
		WithHeader(map[string]string{"Authorization": "Bearer token"}),
		WithLogger(nil),
	)

	fmt.Printf("Client: %+v\n", client)
}
