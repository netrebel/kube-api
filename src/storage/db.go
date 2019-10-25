package storage

import "errors"

var (
	//ErrNotFound - To be returned when not found
	ErrNotFound = errors.New("not found")
)

// DB is the interface to a simple key/value store
type DB interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}
