package broker

import (
	"encoding/json"
	"net/http"

	"github.com/andreykont/KVStorage/api"
)

func (b *Broker) PutStorageValue(w http.ResponseWriter, r *http.Request, key api.PathKey) {
	ctx := r.Context()
	var storageValue api.PutStorageValueJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&storageValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = b.SomeEphemeralStorage.Put(ctx, key, storageValue)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
