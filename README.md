# Golang Functional Options Pattern

The Functional Options Pattern in Go (Golang) is a design pattern that enables flexible and readable object or function configuration. By defining functions (options) that modify the attributes or behavior of a struct during its initialization, this pattern allows customization without the need for numerous constructors. It simplifies handling optional parameters, avoids bulky struct definitions, and keeps code clean and maintainable. This approach promotes a declarative and extensible configuration style.

## Traditional Constructor Method

A common way to initialize an object in Go is through a constructor function that explicitly defines required and optional parameters. Here's an example:

```go
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
```
Usage Example:
```go
client := New("https://api.example.com", map[string]string{"Authorization": "Bearer token"}, myLogger)
```

This method provides clarity by explicitly listing parameters. However, as the number of optional parameters grows, constructors can become cumbersome. Managing defaults and introducing new options may require additional constructors, leading to verbosity and reduced flexibility.

This approach works well for simple configurations with a small number of parameters. However, it becomes less practical for complex setups or frequent changes, where maintaining multiple constructors can make the codebase harder to manage.

## Multiple Constructors for Each Configuration Variant

To handle varying configurations, developers often create separate constructors for each combination of parameters. Here's an example:

```go
type Client struct {
	baseURL    string
	header     map[string]string
	logger     ILogger
	baseClient *http.Client
}

// Constructor with only baseURL
func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     map[string]string{},
		baseClient: &http.Client{},
	}
}

// Constructor with baseURL and headers
func NewWithBaseURLAndHeaders(baseURL string, header map[string]string) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     header,
		baseClient: &http.Client{},
	}
}

// Constructor with baseURL, headers, and logger
func NewWithBaseURLHeadersAndLogger(baseURL string, header map[string]string, logger ILogger) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     header,
		logger:     logger,
		baseClient: &http.Client{},
	}
}
```
Usage Example:
```go
// With base url only.
client := New("https://api.example.com")

// With base url and header.
client := NewWithBaseURLAndHeaders("https://api.example.com", map[string]string{"Authorization": "Bearer token"})

// With all attributes.
client := NewWithBaseURLHeadersAndLogger("https://api.example.com", map[string]string{"Authorization": "Bearer token"}, myLogger)
```

This method accommodates various configurations but can lead to constructor bloat as the number of variants increases.

This approach is suitable for predictable configuration sets but quickly becomes unwieldy for more dynamic or extensible configurations. It may lead to a cluttered codebase and reduced maintainability.

## Using a Custom Config Struct

A `Config` struct centralizes all configuration options, which can then be passed to a single constructor. Here's an example:

```go
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
```
Usage Example:
```go
config := &Config{
	BaseURL: "https://api.example.com",
	Header:  map[string]string{"Authorization": "Bearer token"},
	Logger:  myLogger,
}
client := NewWithConfig(config)
```

Using a Config struct simplifies the constructor and improves maintainability.

This approach is effective for centralizing configuration. However, it may lack the expressiveness and flexibility of other patterns, particularly when adding dynamic or conditional configurations.

## Setter Function Pattern

The Setter Function Pattern initializes an object using a basic constructor, followed by setter methods for additional configuration.

```go
type Client struct {
	baseURL    string
	header     map[string]string
	logger     ILogger
	baseClient *http.Client
}

// Basic constructor
func New(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		header:     map[string]string{},
		baseClient: &http.Client{},
	}
}

// Setter for headers
func (c *Client) SetHeader(header map[string]string) *Client {
	c.header = header
	return c
}

// Setter for logger
func (c *Client) SetLogger(logger ILogger) *Client {
	c.logger = logger
	return c
}
```
Usage Example:
```go
// On instantiation
client := New("https://api.example.com").
	SetHeader(map[string]string{"Authorization": "Bearer token"}).
	SetLogger(myLogger)
// Split
client := New("https://api.example.com")
client.SetHeader(map[string]string{"Authorization": "Bearer token"})
client.SetLogger(myLogger)
```

This method allows incremental configuration by chaining setter calls.

Setter methods are useful for incremental configurations and a fluent API style. However, they may complicate validation and initialization logic if setters are misused or called in the wrong order.

## Functional Options Pattern
The Functional Options Pattern provides a clean and flexible way to configure objects by passing functions (options) to a single constructor.

```go
type ILogger interface{}

type Client struct {
	baseURL    string
	header     map[string]string
	logger     ILogger
	baseClient *http.Client
}

// Option type
type Option func(*Client)

// Basic constructor with functional options
func New(baseURL string, opts ...Option) *Client {
	client := &Client{
		baseURL:    baseURL,
		header:     map[string]string{},
		baseClient: &http.Client{},
	}

	// Apply each option
	for _, opt := range opts {
		opt(client)
	}
	return client
}

// Option to set headers
func WithHeader(header map[string]string) Option {
	return func(c *Client) {
		c.header = header
	}
}

// Option to set logger
func WithLogger(logger ILogger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}
```
Usage Example:
```go
client := New("https://api.example.com",
	WithHeader(map[string]string{"Authorization": "Bearer token"}),
	WithLogger(myLogger),
)
```

The pattern enables developers to configure an object in a highly customizable and expressive way.

This approach is ideal for complex configurations with many optional parameters. It is extensible, avoids constructor bloat, and supports a clean API. However, it can add complexity to debugging and understanding code due to the indirection introduced by options.