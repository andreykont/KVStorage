package broker

import (
	"encoding/json"
	"net/http"

	"github.com/andreykont/KVStorage/api"
)

func (b *Broker) GetStorageValue(w http.ResponseWriter, r *http.Request, key api.PathKey) {
	ctx := r.Context()
	value, err := b.SomeEphemeralStorage.Get(ctx, string(key))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_ = json.NewEncoder(w).Encode(api.ValueTransfer{
		Value: value.(string),
	})
}
