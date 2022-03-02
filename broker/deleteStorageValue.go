package broker

import (
	"net/http"

	"github.com/andreykont/KVStorage/api"
)

func (b *Broker) DeleteStorageValue(w http.ResponseWriter, r *http.Request, key api.PathKey) {
	ctx := r.Context()
	err := b.SomeEphemeralStorage.Delete(ctx, string(key))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
