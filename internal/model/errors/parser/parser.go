package parser

import "errors"

var (
	// ErrEmptyGoMod When the provided go.mod content is empty or nil.
	ErrEmptyGoMod = errors.New("go.mod file is empty")

	// ErrInvalidGoModSyntax When the go.mod file cannot be parsed due to invalid syntax or corruption.
	ErrInvalidGoModSyntax = errors.New("invalid go.mod syntax")

	// ErrMissingModuleDirective When the 'module' directive is missing.
	ErrMissingModuleDirective = errors.New("missing module directive in go.mod")

	// ErrMissingGoVersion When the 'go' version directive is missing.
	ErrMissingGoVersion = errors.New("missing Go version directive in go.mod")
)
