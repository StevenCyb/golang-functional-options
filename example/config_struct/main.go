package main

import (
	"fmt"
	"net/http"
)

type ILogger interface{}

type Config struct {
	BaseURL string
	Header  map[string]string
	Logger  ILogger
}

type Client struct {
	baseURL    string
	header     map[string]string
	logger     ILogger
	baseClient *http.Client
}

func NewWithConfig(config *Config) *Client {
	return &Client{
		baseURL:    config.BaseURL,
		header:     config.Header,
		logger:     config.Logger,
		baseClient: &http.Client{},
	}
}

func main() {
	config := &Config{
		BaseURL: "https://api.example.com",
		Header:  map[string]string{"Authorization": "Bearer token"},
		Logger:  nil,
	}

	client := NewWithConfig(config)
	fmt.Printf("Client: %+v\n", client)
}
