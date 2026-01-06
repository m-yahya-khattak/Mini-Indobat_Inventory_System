package service

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrInvalidStock = errors.New("stock cannot be negative")
	ErrInvalidPrice = errors.New("price cannot be negative")
)

