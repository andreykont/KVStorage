package kvstorage

import (
	"context"
	"fmt"
)

func NewKVStorage() *Storage {
	return &Storage{
		data: map[string]interface{}{},
	}
}
func (s *Storage) Get(ctx context.Context, key string) (interface{}, error) {
	s.dataLock.RLock()
	defer s.dataLock.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("Element is not found")
	}
	return val, nil
}

func (s *Storage) Put(ctx context.Context, key string, val interface{}) error {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	s.data[key] = val
	return nil
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	s.dataLock.Lock()
	defer s.dataLock.Unlock()
	delete(s.data, key)
	return nil
}
