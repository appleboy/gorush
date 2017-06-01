package fcm

import (
	"errors"
	"net/http"
)

// Option configurates Client with defined option.
type Option func(*Client) error

// WithEndpoint returns Option to configure FCM Endpoint.
func WithEndpoint(endpoint string) Option {
	return func(c *Client) error {
		if endpoint == "" {
			return errors.New("invalid endpoint")
		}
		c.endpoint = endpoint
		return nil
	}
}

// WithHTTPClient returns Option to configure HTTP Client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		c.client = httpClient
		return nil
	}
}
