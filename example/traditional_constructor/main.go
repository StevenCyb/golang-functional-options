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

func New(baseURL string, header map[string]string, logger ILogger) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     header,
		logger:     logger,
		baseClient: &http.Client{},
	}
}

func main() {
	header := map[string]string{"Authorization": "Bearer token"}
	client := New("https://api.example.com", header, nil)

	fmt.Printf("Client: %+v\n", client)
}
