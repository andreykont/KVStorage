package broker

import (
	"context"

	"github.com/andreykont/KVStorage/pkg/kvstorage"
)

type Broker struct {
	Storage KVStorage
}

type KVStorage interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Put(ctx context.Context, key string, val interface{}) error
	Delete(ctx context.Context, key string) error
}

func NewBroker() *Broker {
	storage := kvstorage.NewKVStorage()
	return &Broker{Storage: storage}
}
