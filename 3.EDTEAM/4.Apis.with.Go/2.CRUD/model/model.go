package model

import "errors"

var (
	ErrPersonCanNotBeNil    = errors.New("Person cannot be nil")
	ErrIDPersonDoesNotExist = errors.New("Person does not exist")
)
