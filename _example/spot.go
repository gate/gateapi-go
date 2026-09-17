package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/gate/gateapi-go/v7"
	"github.com/shopspring/decimal"
)

// SpotDemo sizes a limit buy in base units and checks its cost against the quote balance.
func SpotDemo(ctx context.Context, client *gateapi.APIClient) error {
	pair, _, err := client.SpotApi.GetCurrencyPair(ctx, "GT_USDT")
	if err != nil {
		return fmt.Errorf("get spot pair: %v", err)
	}
	if pair.AmountPrecision < 0 || pair.AmountPrecision > 18 {
		return fmt.Errorf("unsupported base amount precision")
	}
	minimum, err := positiveAmount("minimum base amount", pair.MinBaseAmount)
	if err != nil {
		return err
	}
	tickers, _, err := client.SpotApi.ListTickers(ctx, &gateapi.ListTickersOpts{CurrencyPair: optional.NewString(pair.Id)})
	if err != nil {
		return fmt.Errorf("get spot ticker: %v", err)
	}
	if len(tickers) == 0 {
		return fmt.Errorf("no spot ticker for %s", pair.Id)
	}
	price, err := positiveAmount("spot price", tickers[0].Last)
	if err != nil {
		return err
	}
	amount := minimum.Mul(decimal.NewFromInt(2))
	if pair.MinQuoteAmount != "" {
		minimumQuote, err := positiveAmount("minimum quote amount", pair.MinQuoteAmount)
		if err != nil {
			return err
		}
		if amount.Mul(price).LessThan(minimumQuote) {
			// Round up to base precision so the resulting order meets the quote minimum.
			amount = minimumQuote.DivRound(price, pair.AmountPrecision)
			if amount.Mul(price).LessThan(minimumQuote) {
				amount = amount.Add(decimal.New(1, -pair.AmountPrecision))
			}
		}
	}
	accounts, _, err := client.SpotApi.ListSpotAccounts(ctx, &gateapi.ListSpotAccountsOpts{Currency: optional.NewString(pair.Quote)})
	if err != nil {
		return fmt.Errorf("get spot balance: %v", err)
	}
	available := decimal.Zero
	for _, account := range accounts {
		if account.Currency == pair.Quote {
			available, err = parseAmount("quote balance", account.Available)
			if err != nil {
				return err
			}
			break
		}
	}
	if available.LessThan(amount.Mul(price)) {
		return fmt.Errorf("insufficient %s balance for spot order", pair.Quote)
	}
	created, _, err := client.SpotApi.CreateOrder(ctx, gateapi.Order{
		CurrencyPair: pair.Id, Type: "limit", Account: "spot", Side: "buy",
		Amount: amount.String(), Price: price.String(), TimeInForce: "gtc",
	}, nil)
	if err != nil {
		return fmt.Errorf("create spot order (check order history before retrying): %v", err)
	}
	logger.Printf("spot order %s status=%s", created.Id, created.Status)
	if created.Status == "open" {
		// Preserve the pair and account when cancelling only the order created by this demo.
		if _, _, err = client.SpotApi.CancelOrder(ctx, created.Id, pair.Id, &gateapi.CancelOrderOpts{Account: optional.NewString("spot")}); err != nil {
			return fmt.Errorf("cancel remaining spot order %s manually if needed: %v", created.Id, err)
		}
	}
	return nil
}
