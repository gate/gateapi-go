package main

import (
	"context"
	gateapi "github.com/gate/gateapi-go/v7"
	"net/http"
)

// demoGetPosition follows the generated variant's signature without supplying an internal user override.
func demoGetPosition(ctx context.Context, client *gateapi.APIClient, settle, contract string) (gateapi.Position, *http.Response, error) {
	return client.FuturesApi.GetPosition(ctx, settle, contract)
}
