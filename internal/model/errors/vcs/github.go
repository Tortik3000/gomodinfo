package vcs

import (
	"errors"
)

var (
	// ErrInvalidRepoReference When the provided repository URL is empty or malformed.
	ErrInvalidRepoReference = errors.New("invalid repository reference")

	// ErrUnsupportedHost When the repository host is not supported (e.g., GitLab, Bitbucket, etc.).
	ErrUnsupportedHost = errors.New("unsupported repository host")

	// ErrNotFound When the go.mod file is missing in the repository.
	ErrNotFound = errors.New("go.mod file or repository not found")

	// ErrNetwork When a network or connection problem occurs.
	ErrNetwork = errors.New("network or connection error")

	// ErrDecodingContent When decoding the file content fails.
	ErrDecodingContent = errors.New("failed to decode file content")
)
