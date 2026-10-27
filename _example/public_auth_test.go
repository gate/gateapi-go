package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunUsesSDKSigning(t *testing.T) {
	// Exercise the CLI's shared-client path: merely setting legacy configuration keys does not sign requests.
	signed := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var body string
		switch r.URL.Path {
		case "/spot/currency_pairs/GT_USDT":
			body = `{"id":"GT_USDT","quote":"USDT","min_base_amount":"1"}`
		case "/spot/tickers":
			body = `[{"last":"1"}]`
		case "/spot/accounts":
			signed = r.Header.Get("KEY") == "synthetic-key" && r.Header.Get("SIGN") != "" && r.Header.Get("Timestamp") != ""
			body = `[]`
		default:
			t.Errorf("unexpected write or endpoint: %s", r.URL.Path)
			w.WriteHeader(500)
		}
		if _, err := w.Write([]byte(body)); err != nil {
			t.Error(err)
		}
	}))
	config := &RunConfig{ApiKey: "synthetic-key", ApiSecret: "synthetic-secret", BaseUrl: server.URL}
	err := run(config, []string{"spot"})
	server.Close()
	if err == nil || !strings.Contains(err.Error(), "insufficient") {
		t.Fatalf("expected balance rejection, got %v", err)
	}
	if !signed {
		t.Fatal("protected request was not signed by the SDK")
	}
}
