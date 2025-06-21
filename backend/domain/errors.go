package domain

import "errors"

var (
	ErrConflict       = errors.New("value already exists")
	ErrNotFound       = errors.New("not found")
	ErrNotImplemented = errors.New("not implemented")
)
