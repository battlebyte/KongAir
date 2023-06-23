// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
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

	. "github.com/Kong/KongAir/flight-data/routes/api/models"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Request all the KongAir routes
	// (GET /routes)
	GetRoutes(ctx echo.Context, params GetRoutesParams) error
	// Request a specific KongAir route by id (LHR-SIN)
	// (GET /routes/{id})
	GetRoute(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetRoutes converts echo context to params.
func (w *ServerInterfaceWrapper) GetRoutes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRoutesParams
	// ------------- Optional query parameter "origin" -------------

	err = runtime.BindQueryParameter("form", true, false, "origin", ctx.QueryParams(), &params.Origin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter origin: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRoutes(ctx, params)
	return err
}

// GetRoute converts echo context to params.
func (w *ServerInterfaceWrapper) GetRoute(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRoute(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/routes", wrapper.GetRoutes)
	router.GET(baseURL+"/routes/:id", wrapper.GetRoute)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/5yU34/jNBDH/xVr4AGkbFPgdA95W+kEVCBA28dNhbzJJJ0jsb3jSY+oyv+ObCf9dT3u",
	"4Kmpxx7PfL+f8REq2ztr0IiH4gi+2mOv4+eTHQTDh2PrkIUwLutD+2c9sBayJvyX0SEUQEawRYYpgxq9",
	"kLnd4IXJtCFO9d1ly9TSvRNTBoyvAzHWUDyH46fN15dl18XtsiWTfXmPlcAUUpFpbLikRl8xuVQl/GJN",
	"+0isYtNeeeQDVagc2wPV6JXsUTG25AUZa8Vp23Kq6Qi9ekH5gGhUqk1pU6uL6pQmdpbFQwZC0oW65tu2",
	"6TbI4IDsU0HfrdardZTFodGOoIAf4lIGTss+WpGnMsJni/JxU08oAxuvdNelBq6rJq+qgRmNdONlc2JV",
	"042r0jwmOcJGrZymYMcXdqkqW6NflcGUQE8Mbmoo4CeU1HbshHWPguyheL6tvqFOkNXLqE5mBzzgdUAe",
	"IQOj++jtEkzoBhVIsPd3GZsXNLMew38vYzSisdzDtAugeWeNT6J+v16Hn8oaQRP11c51VMVm8vc+8Y1/",
	"69516US35wef+Drobgipn4+lUUqVQHUJhSrh15+fHrY//l5CNgdSB6fgOXCha4peHbtkPYTfvl2XRk27",
	"0sSRuSPH14wNFPBVfp75fB74PE37RxJNaZ4vbNkOVYXeN0OnFrHUB5L97YjMJsdShr7XPEYiXwf0ciJy",
	"QZEXIkS3gQVoOmr38lBr0bALOWbW8yPV0+eBV95hRQ1VKbNq2PZKq5YOaGaoN+/+Bc7PsbmNRM3JN+8W",
	"NMNknsmMD9X55RIe8NKW21fu/9F3zvcF7v5HNxs7mNnIQMab9Zs7kkcJjJW0+1N+nw25sjxMN9XqmzgT",
	"m9++/TQBIS3y4b4dS87HPzbxMUWGDAbuoIC9iPNFnmtHq7+saR808aqyYdqnfwIAAP//hNa45v0GAAA=",
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
