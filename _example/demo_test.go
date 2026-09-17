package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gate/gateapi-go/v7"
)

// fixture validates an ordered HTTP exchange, including query filters and JSON wire types.
type fixture struct {
	method, path, query, response string
	body                          map[string]interface{}
	status                        int
}

func mockClient(t *testing.T, steps []fixture) (*gateapi.APIClient, func()) {
	t.Helper()
	index := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if index >= len(steps) {
			t.Errorf("unexpected request %s %s", r.Method, r.URL)
			http.Error(w, "unexpected", 500)
			return
		}
		step := steps[index]
		index++
		if r.Method != step.method || r.URL.Path != step.path || r.URL.RawQuery != step.query {
			t.Errorf("request %d got %s %s; want %s %s?%s", index, r.Method, r.URL, step.method, step.path, step.query)
		}
		if step.body != nil {
			var got map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
				t.Error(err)
			}
			for key, value := range step.body {
				if got[key] != value {
					t.Errorf("body[%s]=%v; want %v", key, got[key], value)
				}
			}
		}
		w.Header().Set("Content-Type", "application/json")
		status := step.status
		if status == 0 {
			status = 200
		}
		w.WriteHeader(status)
		if status != 204 {
			_, err := w.Write([]byte(step.response))
			if err != nil {
				t.Error(err)
			}
		}
	}))
	cfg := gateapi.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.HTTPClient = server.Client()
	cfg.HTTPClient.Timeout = time.Second
	return gateapi.NewAPIClient(cfg), func() {
		server.Close()
		if index != len(steps) {
			t.Errorf("received %d/%d requests", index, len(steps))
		}
	}
}

func TestFuturesAmounts(t *testing.T) {
	// Short positions must keep the sign only on size, never on collateral; fractional minima stay exact.
	size, margin, err := futuresAmounts("-0.5", "10.5", "2", "1", "2")
	if err != nil || size.String() != "-10.5" || margin.String() != "11.55" {
		t.Fatalf("size=%s margin=%s err=%v", size, margin, err)
	}
	for _, input := range []struct{ position, minimum, price, multiplier, leverage string }{
		{"bad", "1", "1", "1", "1"}, {"0", "bad", "1", "1", "1"}, {"0", "0", "1", "1", "1"}, {"0", "1", "0", "1", "1"}, {"0", "1", "1", "1", "0"},
	} {
		if _, _, err := futuresAmounts(input.position, input.minimum, input.price, input.multiplier, input.leverage); err == nil {
			t.Errorf("accepted invalid input %+v", input)
		}
	}
}

func futuresReads(available string) []fixture {
	return []fixture{
		{method: "GET", path: "/futures/usdt/positions/BTC_USDT", response: `{"size":"-0.5"}`},
		{method: "GET", path: "/futures/usdt/contracts/BTC_USDT", response: `{"order_size_min":"10.5","quanto_multiplier":"1"}`},
		{method: "GET", path: "/futures/usdt/tickers", query: "contract=BTC_USDT", response: `[{"last":"2"}]`},
		{method: "GET", path: "/futures/usdt/accounts", response: `{"available":"` + available + `"}`},
	}
}

func TestFuturesWireAndTransferShortfall(t *testing.T) {
	steps := append(futuresReads("1"), []fixture{
		{method: "POST", path: "/wallet/transfers", body: map[string]interface{}{"from": "spot", "to": "futures", "settle": "usdt", "amount": "6.7"}, response: `{"tx_id":1}`},
		{method: "POST", path: "/futures/usdt/positions/BTC_USDT/leverage", query: "leverage=3", response: `{}`},
		{method: "DELETE", path: "/futures/usdt/orders", query: "contract=BTC_USDT", response: `[]`},
		{method: "POST", path: "/futures/usdt/orders", body: map[string]interface{}{"contract": "BTC_USDT", "size": "-10.5", "price": "0", "tif": "ioc"}, response: `{"id":42,"status":"open","size":"-10.5"}`},
		{method: "DELETE", path: "/futures/usdt/orders/42", response: `{"id":42,"status":"finished"}`},
	}...)
	client, done := mockClient(t, steps)
	defer done()
	if err := FuturesDemo(context.Background(), client, false); err != nil {
		t.Fatal(err)
	}
}

func TestFuturesRejectsInsufficientTestnetFunds(t *testing.T) {
	client, done := mockClient(t, futuresReads("0"))
	defer done()
	if err := FuturesDemo(context.Background(), client, true); err == nil {
		t.Fatal("expected funding error before writes")
	}
}

func TestFuturesRejectsEmptyTicker(t *testing.T) {
	steps := futuresReads("1")[:3]
	steps[2].response = `[]`
	client, done := mockClient(t, steps)
	defer done()
	if err := FuturesDemo(context.Background(), client, false); err == nil {
		t.Fatal("expected empty ticker error")
	}
}

func TestFuturesDoesNotRetryUnknownWriteOutcome(t *testing.T) {
	steps := append(futuresReads("100"), fixture{method: "POST", path: "/futures/usdt/positions/BTC_USDT/leverage", query: "leverage=3", response: `{}`}, fixture{method: "DELETE", path: "/futures/usdt/orders", query: "contract=BTC_USDT", response: `[]`}, fixture{method: "POST", path: "/futures/usdt/orders", status: 500, response: `{"label":"INTERNAL_SERVER_ERROR","message":"failed"}`})
	client, done := mockClient(t, steps)
	defer done()
	if err := FuturesDemo(context.Background(), client, false); err == nil || !strings.Contains(err.Error(), "before retrying") {
		t.Fatalf("expected unknown-outcome error, got %v", err)
	}
}

func spotReads(balance string) []fixture {
	return []fixture{
		{method: "GET", path: "/spot/currency_pairs/GT_USDT", response: `{"id":"GT_USDT","quote":"USDT","min_base_amount":"0.1","min_quote_amount":"1","amount_precision":2}`},
		{method: "GET", path: "/spot/tickers", query: "currency_pair=GT_USDT", response: `[{"last":"10"}]`},
		{method: "GET", path: "/spot/accounts", query: "currency=USDT", response: `[{"currency":"USDT","available":"` + balance + `"}]`},
	}
}

func TestSpotChecksQuoteCost(t *testing.T) {
	// 0.2 base units cost 2 USDT: a balance of 1 must fail despite being greater than 0.2.
	client, done := mockClient(t, spotReads("1"))
	defer done()
	if err := SpotDemo(context.Background(), client); err == nil {
		t.Fatal("expected insufficient quote balance")
	}
}

func TestSpotWireAndScopedCancel(t *testing.T) {
	steps := append(spotReads("10"), fixture{method: "POST", path: "/spot/orders", body: map[string]interface{}{"amount": "0.2", "price": "10", "account": "spot"}, response: `{"id":"spot-42","status":"open"}`}, fixture{method: "DELETE", path: "/spot/orders/spot-42", query: "account=spot&currency_pair=GT_USDT", response: `{"id":"spot-42","status":"cancelled"}`})
	client, done := mockClient(t, steps)
	defer done()
	if err := SpotDemo(context.Background(), client); err != nil {
		t.Fatal(err)
	}
}

func marginReads(loans string) []fixture {
	return []fixture{
		{method: "GET", path: "/margin/uni/currency_pairs/BTC_USDT", response: `{"status":"enabled","quote_min_borrow_amount":"1.25"}`},
		{method: "GET", path: "/margin/uni/loans", query: "currency=USDT&currency_pair=BTC_USDT&limit=1", response: loans},
	}
}

func TestMarginProtectsExistingDebt(t *testing.T) {
	client, done := mockClient(t, marginReads(`[{"currency":"USDT"}]`))
	defer done()
	if err := MarginDemo(context.Background(), client); err == nil {
		t.Fatal("expected existing debt rejection")
	}
}

func TestMarginBorrowRepayWire(t *testing.T) {
	steps := append(marginReads(`[]`), fixture{method: "GET", path: "/margin/uni/borrowable", query: "currency=USDT&currency_pair=BTC_USDT", response: `{"borrowable":"10"}`}, fixture{method: "POST", path: "/margin/uni/loans", body: map[string]interface{}{"type": "borrow", "amount": "1.25", "currency": "USDT", "currency_pair": "BTC_USDT"}, status: 204}, fixture{method: "POST", path: "/margin/uni/loans", body: map[string]interface{}{"type": "repay", "amount": "1.25", "currency": "USDT", "currency_pair": "BTC_USDT", "repaid_all": true}, status: 204})
	client, done := mockClient(t, steps)
	defer done()
	if err := MarginDemo(context.Background(), client); err != nil {
		t.Fatal(err)
	}
}

func TestDemoCancellation(t *testing.T) {
	client, done := mockClient(t, nil)
	defer done()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := SpotDemo(ctx, client); err == nil {
		t.Fatal("cancelled context must stop requests")
	}
}

func TestNilEndpoint(t *testing.T) {
	config, err := NewRunConfig("synthetic-key", "synthetic-secret", nil)
	if err != nil || config.BaseUrl != "https://api.gateio.ws/api/v4" {
		t.Fatalf("config=%+v err=%v", config, err)
	}
}

func TestMarginRepayFailureReportsOutstandingDebt(t *testing.T) {
	// A failed repayment must be visible and must not cause a second borrow or an automatic write retry.
	steps := append(marginReads(`[]`), fixture{method: "GET", path: "/margin/uni/borrowable", query: "currency=USDT&currency_pair=BTC_USDT", response: `{"borrowable":"10"}`}, fixture{method: "POST", path: "/margin/uni/loans", status: 204}, fixture{method: "POST", path: "/margin/uni/loans", status: 500, response: `{"label":"INTERNAL_SERVER_ERROR","message":"failed"}`})
	client, done := mockClient(t, steps)
	defer done()
	if err := MarginDemo(context.Background(), client); err == nil || !strings.Contains(err.Error(), "including accrued interest") {
		t.Fatalf("expected actionable debt error, got %v", err)
	}
}

func TestRequestDeadline(t *testing.T) {
	// Block until the client cancels, proving the propagated deadline reaches the HTTP request.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { <-r.Context().Done() }))
	defer server.Close()
	cfg := gateapi.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.HTTPClient = server.Client()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if err := SpotDemo(ctx, gateapi.NewAPIClient(cfg)); err == nil {
		t.Fatal("expected deadline error")
	}
}

func TestAmountBounds(t *testing.T) {
	for _, value := range []string{"", "1e100000", "1e-100000", strings.Repeat("1", 129)} {
		if _, err := parseAmount("test amount", value); err == nil {
			t.Errorf("accepted unbounded amount %q", value)
		}
	}
}

func TestSpotQuoteMinimumRoundsUp(t *testing.T) {
	// A quote minimum must become a base quantity at the declared precision, never a quote-sized order.
	steps := spotReads("10")
	steps[0].response = `{"id":"GT_USDT","quote":"USDT","min_base_amount":"0.1","min_quote_amount":"1","amount_precision":2}`
	steps[1].response = `[{"last":"3"}]`
	steps = append(steps, fixture{method: "POST", path: "/spot/orders", body: map[string]interface{}{"amount": "0.34", "price": "3"}, response: `{"id":"spot-43","status":"closed"}`})
	client, done := mockClient(t, steps)
	defer done()
	if err := SpotDemo(context.Background(), client); err != nil {
		t.Fatal(err)
	}
}

func TestMarginRejectsInsufficientCapacity(t *testing.T) {
	steps := append(marginReads(`[]`), fixture{method: "GET", path: "/margin/uni/borrowable", query: "currency=USDT&currency_pair=BTC_USDT", response: `{"borrowable":"0.5"}`})
	client, done := mockClient(t, steps)
	defer done()
	if err := MarginDemo(context.Background(), client); err == nil {
		t.Fatal("expected insufficient capacity before borrowing")
	}
}

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
