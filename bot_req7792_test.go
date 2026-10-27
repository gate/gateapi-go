package gateapi

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/antihax/optional"
)

// These are wire regressions against the generated SDK, not a second HTTP client.
func TestREQ7792BotFailureFacts(t *testing.T) {
	// A failed write may have no data or only an identity; never require invented state.
	for _, body := range []string{`{"code":400,"label":"RECOVERY_REQUIRED","message":"verify"}`, `{"code":400,"label":"RECOVERY_REQUIRED","message":"verify","data":{"strategy_id":"known-id","operation_status":"pending","recovery_required":true}}`} {
		var result AiHubCreateSuccessResponse
		if err := json.Unmarshal([]byte(body), &result); err != nil {
			t.Fatal(err)
		}
		if result.Code != 400 || result.Label != "RECOVERY_REQUIRED" || result.Data.Status != "" {
			t.Fatalf("lost or invented facts: %+v", result)
		}
		if strings.Contains(body, "known-id") && result.Data.StrategyId != "known-id" {
			t.Fatal("partial identity lost")
		}
	}
}

// The wire peer checks method, body and one request per explicit SDK invocation.
func TestREQ7792BotWire(t *testing.T) {
	calls := 0
	peer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		raw, err := ioutil.ReadAll(io.LimitReader(r.Body, 8193))
		if err != nil {
			t.Error(err)
		}
		var body map[string]json.RawMessage
		if len(raw) > 0 {
			if err := json.Unmarshal(raw, &body); err != nil {
				t.Error(err)
			}
		}
		switch r.URL.Path {
		case "/bot/strategy/recommend":
			if r.URL.Query().Get("exclude_recommendation_id") != "smart_position|USDT|123|a" {
				t.Error("lost exclusion identity")
			}
		case "/bot/smart-position/create":
			if !strings.Contains(string(raw), `"close_flag":0`) {
				t.Error("lost explicit close_flag=0")
			}
		case "/bot/futures-combo/create":
			if body["create_params"] == nil {
				t.Error("lost combo parameters")
			}
		case "/bot/portfolio/stop":
			if string(body["strategy_id"]) != `"9223372036854775807"` {
				t.Error("lost exact strategy ID")
			}
			if r.URL.Query().Get("mode") == "limit" {
				if body["stop_options"] == nil {
					t.Error("lost limit options")
				}
			} else if body["stop_options"] != nil {
				t.Error("absent stop options must not become {}")
			}
		default:
			t.Errorf("unexpected route %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"code":400,"label":"RECOVERY_REQUIRED","message":"verify","trace_id":"fixture-trace","data":{"strategy_id":"9223372036854775807","strategy_type":"smart_position","status":"starting","operation_status":"pending","recovery_required":false}}`)
	}))
	defer peer.Close()
	cfg := NewConfiguration()
	cfg.BasePath = peer.URL
	cfg.HTTPClient = peer.Client()
	api := NewAPIClient(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, _, err := api.BotApi.GetAIHubStrategyRecommend(ctx, &GetAIHubStrategyRecommendOpts{ExcludeRecommendationId: optional.NewString("smart_position|USDT|123|a")})
	if err != nil {
		t.Fatal(err)
	}
	zero := int32(0)
	result, _, err := api.BotApi.PostAIHubSmartPositionCreate(ctx, AiHubRebalanceCreateRequest{StrategyType: "smart_position", CreateParams: AiHubRebalanceCreateParams{Period: "1h", CloseFlag: &zero}}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Label != "RECOVERY_REQUIRED" || result.Data.StrategyId != "9223372036854775807" || result.Data.RecoveryRequired == nil || *result.Data.RecoveryRequired {
		t.Fatalf("lost partial/error facts: %#v", result)
	}
	_, _, err = api.BotApi.PostAIHubFuturesComboCreate(ctx, AiHubFuturesComboCreateRequest{StrategyType: "futures_combo"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = api.BotApi.PostAIHubPortfolioStop(ctx, AiHubPortfolioStopRequest{StrategyId: "9223372036854775807", StrategyType: StrategyType("futures_combo")}, nil)
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(AiHubPortfolioStopRequest{StrategyId: "1", StrategyType: StrategyType("smart_position"), StopMode: "limit", StopOptions: &AiHubRebalanceStopOptions{Price: "10", Expiration: 60}})
	if err != nil || !strings.Contains(string(encoded), `"stop_options":{"price":"10","expiration":60`) {
		t.Fatalf("lost options: %s %v", encoded, err)
	}
	if calls != 4 {
		t.Fatalf("writes must not retry: %d", calls)
	}
}

// Portfolio projections preserve observed false, while missing progress stays
// unverified. String monetary values and saved configuration must survive decoding.
func TestREQ7792PortfolioProjections(t *testing.T) {
	peer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/bot/portfolio/running":
			_, _ = io.WriteString(w, `{"code":200,"data":{"items":[{"strategy_id":"9223372036854775807","strategy_type":"futures_combo","assets":["BTC","ETH"],"markets":["BTC_USDT","ETH_USDT"]}],"total":1,"page":1,"page_size":10}}`)
		case "/bot/portfolio/detail":
			if r.URL.Query().Get("strategy_id") == "unverified" {
				_, _ = io.WriteString(w, `{"code":200,"data":{"strategy_id":"unverified","positions_status":"unverified"}}`)
				return
			}
			_, _ = io.WriteString(w, `{"code":200,"data":{"strategy_id":"9223372036854775807","strategy_type":"futures_combo","status":"stopping","positions_status":"partial","positions":[{"market":"BTC_USDT","amount":"0.000000000000000001","entry_price":"100.000000000000000001"}],"configuration":{"stop_loss_price":{"price":"10.000000000000000001","price_type":"mark","is_gte":false}},"stop_progress":{"trading_finished":false,"funds_returned":false,"account_released":false,"refund_amount":"0.000000000000000001","refund_currency":"USDT"}}}`)
		default:
			t.Errorf("unexpected route %s", r.URL.Path)
			w.WriteHeader(404)
		}
	}))
	defer peer.Close()
	cfg := NewConfiguration()
	cfg.BasePath = peer.URL
	cfg.HTTPClient = peer.Client()
	api := NewAPIClient(cfg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	detail, _, err := api.BotApi.GetAIHubPortfolioDetail(ctx, "9223372036854775807", "futures_combo", nil)
	if err != nil {
		t.Fatal(err)
	}
	if detail.Data.StopProgress == nil || detail.Data.StopProgress.FundsReturned == nil || *detail.Data.StopProgress.FundsReturned {
		t.Fatal("lost explicit false progress")
	}
	raw, err := json.Marshal(detail)
	if err != nil {
		t.Fatal(err)
	}
	for _, fact := range []string{`"funds_returned":false`, `"trading_finished":false`, `"account_released":false`, `100.000000000000000001`, `10.000000000000000001`, `"positions_status":"partial"`} {
		if !strings.Contains(string(raw), fact) {
			t.Errorf("lost %s in %s", fact, raw)
		}
	}
	unknown, _, err := api.BotApi.GetAIHubPortfolioDetail(ctx, "unverified", "smart_position", nil)
	if err != nil {
		t.Fatal(err)
	}
	raw, err = json.Marshal(unknown)
	if err != nil {
		t.Fatal(err)
	}
	if unknown.Data.StopProgress != nil || strings.Contains(string(raw), `"stop_progress"`) {
		t.Fatalf("absent progress invented: %s", raw)
	}
	running, _, err := api.BotApi.GetAIHubPortfolioRunning(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(running.Data.Items) != 1 || len(running.Data.Items[0].Assets) != 2 || len(running.Data.Items[0].Markets) != 2 {
		t.Fatal("lost combination fields")
	}
}
