package entity

import "errors"

var (
	ErrNotFound     = errors.New("entity not found")
	ErrInvalidEnt   = errors.New("invalid entity")
	ErrInvalidEmail = errors.New("invalid email")
	ErrCantDelete   = errors.New("cannot be deleted")
)
