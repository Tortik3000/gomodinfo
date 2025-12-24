package vcs

import "errors"

var (
	ErrInvalidRepoReference = errors.New("invalid repository reference")
	ErrUnsupportedHost      = errors.New("unsupported host")
	ErrNotFound             = errors.New("repository or go.mod not found")
	ErrNetwork              = errors.New("network error")
	ErrDecodingContent      = errors.New("failed to decode content")
)
