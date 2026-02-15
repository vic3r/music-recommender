package domain

import "errors"

var (
	ErrDimensionMismatch = errors.New("embedding dimension does not match store dimension")
	ErrNotFound         = errors.New("song not found")
)
