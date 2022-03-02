package kvstorage

import (
	"sync"
)

type Storage struct {
	data     map[string]interface{}
	dataLock sync.RWMutex
}
