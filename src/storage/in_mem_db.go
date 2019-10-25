package storage

import "sync"

type inMemoryDB struct {
	m   map[string][]byte
	lck sync.RWMutex
}

//NewInMemoryDB - creates new DB implementation to store data in mem
//All ops are concurrency safe
func NewInMemoryDB() DB {
	return &inMemoryDB{m: make(map[string][]byte)}
}

// 'Get' is the interface implementation
func (d *inMemoryDB) Get(key string) ([]byte, error) {
	d.lck.Lock()
	defer d.lck.Unlock()

	v, ok := d.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

// 'Set' is the interface implementation
func (d *inMemoryDB) Set(key string, val []byte) error {
	d.lck.Lock()
	defer d.lck.Unlock()
	d.m[key] = val
	return nil
}
