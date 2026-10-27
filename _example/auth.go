package main

import (
	"context"
	gateapi "github.com/gate/gateapi-go/v7"
)

// demoAuthentication preserves the existing APIv4 key/secret signing contract.
func demoAuthentication(ctx context.Context, config *RunConfig) (context.Context, error) {
	return context.WithValue(ctx, gateapi.ContextGateAPIV4, gateapi.GateAPIV4{
		Key: config.ApiKey, Secret: config.ApiSecret,
	}), nil
}
