// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// Объект для передачи value в json структуре
type ValueTransfer struct {
	Value string `json:"value"`
}

// PathKey defines model for PathKey.
type PathKey string

// PutStorageValueJSONBody defines parameters for PutStorageValue.
type PutStorageValueJSONBody ValueTransfer

// PutStorageValueJSONRequestBody defines body for PutStorageValue for application/json ContentType.
type PutStorageValueJSONRequestBody PutStorageValueJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Удаляет значение из хранилища
	// (DELETE /storage/{key})
	DeleteStorageValue(w http.ResponseWriter, r *http.Request, key PathKey)
	// Отдает значение из хранилища
	// (GET /storage/{key})
	GetStorageValue(w http.ResponseWriter, r *http.Request, key PathKey)
	// Создает или переписывает значение в хранилище
	// (PUT /storage/{key})
	PutStorageValue(w http.ResponseWriter, r *http.Request, key PathKey)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// DeleteStorageValue operation middleware
func (siw *ServerInterfaceWrapper) DeleteStorageValue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "key" -------------
	var key PathKey

	err = runtime.BindStyledParameter("simple", false, "key", chi.URLParam(r, "key"), &key)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter key: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteStorageValue(w, r, key)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetStorageValue operation middleware
func (siw *ServerInterfaceWrapper) GetStorageValue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "key" -------------
	var key PathKey

	err = runtime.BindStyledParameter("simple", false, "key", chi.URLParam(r, "key"), &key)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter key: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetStorageValue(w, r, key)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PutStorageValue operation middleware
func (siw *ServerInterfaceWrapper) PutStorageValue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "key" -------------
	var key PathKey

	err = runtime.BindStyledParameter("simple", false, "key", chi.URLParam(r, "key"), &key)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter key: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutStorageValue(w, r, key)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/storage/{key}", wrapper.DeleteStorageValue)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/storage/{key}", wrapper.GetStorageValue)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/storage/{key}", wrapper.PutStorageValue)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xUXUvbbhT/KuH8/5dZk6pjkssxGGM3wsZu1IusfbTRNnl88qRQSqBWRDfHhN1sMHDI",
	"vkAtllZr61c45xuN87RZrUa2weZNaZLz8ns55zShFNVkFIpQx+A1QfrKrwktlHla8XXlpWjw37KISyqQ",
	"OohC8AC/4BUdW3iJQ/pIB9gBGwL+IH1dARtCvybAg23RABuU2EkCJcrgaZUIG+JSRdR8LqobksNirYJw",
	"E9I0zT6a7m/8aiJeKz+MN4TKwXCCZ/Qee3hJbQvPcciArrFHLezhOXboAAdWnWtY2LW24ii0aJfa1KI9",
	"TqE9DgQbpIqkUDoQpqlJyMN2k8fqNGzdzsKit1uipCcUgnAjyoF7aqB1cUC7VkVrmWGmFnbwDMfUpiOL",
	"9s3jCAc4xAG9wx5eWWuQCf0I+zgy1HomprcGYIMOdJVBgA11oeJJP7dQLLiQ2hBJEfoyAA8WC25hkRn7",
	"umLIOrGOlL8pnOa2aKQTzFWhDX8WxWfsL8rgwTPz/tUk3Bhj6syGZbUJ/yuxAR7858xGypmFONkwpess",
	"ZSyjMJ5IvuC6OXJ9Zw9xaIiOLdqjXWPuIT8yrSV3KSfr87w8Fo4mPx28wPNpLezmyGwcjpNazVeNG+3p",
	"GHs8Xv3bdQfYv1ulw8A2hb6r33Oh/6l4pSjUIjSNfSmrQcm0dnjq+d1s5fL6THfOmV84M8u3TZm6sP8A",
	"DpxQ23jwp/rLJEf/leRv6r+TiFg/jcqNB5T+29xl6/Cp6xj6V6wQtXCM3eyk1Kcc5y9v+nt7d9Pix7kh",
	"n9hT3kQ65ps7pkMc4Blesv7zHp7iGPszF41PsyN9zbeQjrB7r8v3zAk3EaqeeXcL3lccMxYc0gcc0RFe",
	"gA2JqoIHDt/B1P51hmVGtfszke+15zjFhScFt+AWit6yu1ycVFtPfwQAAP//Of9b60QHAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
