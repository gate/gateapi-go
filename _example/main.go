package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gate/gateapi-go/v7"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(flag.CommandLine.Output(), "", log.LstdFlags)

// run owns the public SDK client and deadlines; demos reuse its transport and authentication.
func run(config *RunConfig, demos []string) error {
	for _, demo := range demos {
		if demo != "spot" && demo != "margin" && demo != "futures" {
			return fmt.Errorf("unknown demo %q: use spot, margin or futures", demo)
		}
	}
	cfg := gateapi.NewConfiguration()
	cfg.BasePath = config.BaseUrl
	cfg.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	defer cfg.HTTPClient.CloseIdleConnections()
	client := gateapi.NewAPIClient(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	// Public request signing is driven by the SDK context, not its legacy Key/Secret fields.
	ctx = context.WithValue(ctx, gateapi.ContextGateAPIV4, gateapi.GateAPIV4{Key: config.ApiKey, Secret: config.ApiSecret})
	for _, demo := range demos {
		var err error
		switch demo {
		case "spot":
			err = SpotDemo(ctx, client)
		case "margin":
			err = MarginDemo(ctx, client)
		case "futures":
			err = FuturesDemo(ctx, client, config.UseTestNet)
		}
		if err != nil {
			return fmt.Errorf("%s demo stopped: %v", demo, err)
		}
	}
	return nil
}

func main() {
	var key, secret, baseURL string
	flag.StringVar(&key, "k", "", "Gate APIv4 key")
	flag.StringVar(&secret, "s", "", "Gate APIv4 secret")
	flag.StringVar(&baseURL, "u", "", "API base URL")
	flag.Parse()
	if key == "" || secret == "" || flag.NArg() == 0 {
		logger.Printf("Usage: %s -k <api-key> -s <api-secret> [-u <base-url>] <spot|margin|futures>", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	config, err := NewRunConfig(key, secret, &baseURL)
	if err == nil {
		err = run(config, flag.Args())
	}
	if err != nil {
		// run returns first so its deferred cancellation and transport cleanup always execute.
		logger.Print(err)
		os.Exit(1)
	}
}
