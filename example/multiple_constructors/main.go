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

func NewWithBaseURLAndHeaders(baseURL string, header map[string]string) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     header,
		baseClient: &http.Client{},
	}
}

func NewWithBaseURLHeadersAndLogger(baseURL string, header map[string]string, logger ILogger) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     header,
		logger:     logger,
		baseClient: &http.Client{},
	}
}

func main() {
	header := map[string]string{"Authorization": "Bearer token"}
	client := NewWithBaseURLHeadersAndLogger("https://api.example.com", header, nil)

	fmt.Printf("Client: %+v\n", client)
}
