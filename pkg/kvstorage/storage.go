package kvstorage

import (
	"context"
)

func NewKVStorage() *Storage {
	return &Storage{}
}
func (s *Storage) Get(ctx context.Context, key string) (interface{}, error) {
	s.dataLock.RLock()
	defer s.dataLock.RUnlock()
	return "fd", nil
}

func (s *Storage) Put(ctx context.Context, key string, val interface{}) error {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	return nil
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	return nil
}
