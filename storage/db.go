// Package storage contains interface functions to update the in-memory database
package storage

import "errors"

//ErrNotFound is the error returned when not found (404)
var ErrNotFound = errors.New("not found")

// DB is the interface to a simple key/value store
type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}
