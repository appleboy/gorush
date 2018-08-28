// +build go1.6,!go1.7

package apns2

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
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
		return ctxhttp.Do(ctx, c.HTTPClient, req)
	}
	return c.HTTPClient.Do(req)
}
