package model

import "errors"

// ErrNotFound error if symbol not found
var ErrNotFound = errors.New("not found")

// ErrIsExist error is symbol already exists
var ErrIsExist = errors.New("already exists")
