package go_versions

import "errors"

var (
	// ErrInvalidModulePath When the module path is invalid or cannot be escaped.
	ErrInvalidModulePath = errors.New("invalid module path")

	// ErrProxyUnavailable When the Go proxy (proxy.golang.org) cannot be reached due to network issues.
	ErrProxyUnavailable = errors.New("unable to reach Go proxy")

	// ErrModuleNotFound When the Go proxy returns a non-200 status (e.g., 404 Not Found, 410 Gone).
	ErrModuleNotFound = errors.New("module not found on Go proxy")

	// ErrInvalidProxyResponse When the Go proxy returns unexpected data or the JSON response cannot be decoded.
	ErrInvalidProxyResponse = errors.New("invalid or malformed response from Go proxy")
)
