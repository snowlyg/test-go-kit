package service

import (
	"context"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kataras/iris/v12"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) *iris.Application {
	app := iris.Default()
	app.Use(commonMiddleware)

	app.Get("/status", iris.FromStd(httptransport.NewServer(
		endpoints.StatusEndpoint,
		decodeStatusRequest,
		encodeResponse,
	)))

	app.Get("/get", iris.FromStd(httptransport.NewServer(
		endpoints.GetEndpoint,
		decodeGetRequest,
		encodeResponse,
	)))
	app.Post("/validate", iris.FromStd(httptransport.NewServer(
		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	)))

	return app
}

func commonMiddleware(ctx iris.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.Next()
}
