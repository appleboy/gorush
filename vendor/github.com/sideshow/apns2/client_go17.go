// +build go1.7

package apns2

import (
	"context"
	"net/http"
)

// A Context carries a deadline, a cancellation signal, and other values across
// API boundaries.
//
// Context's methods may be called by multiple goroutines simultaneously.
type Context interface {
	context.Context
}

func (c *Client) requestWithContext(ctx Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return c.HTTPClient.Do(req)
}
