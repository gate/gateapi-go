package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/gate/gateapi-go/v7"
)

// MarginDemo uses the current isolated-margin borrow/repay contract, replacing removed P2P loan APIs.
func MarginDemo(ctx context.Context, client *gateapi.APIClient) error {
	pair, currency := "BTC_USDT", "USDT"
	market, _, err := client.MarginUniApi.GetUniCurrencyPair(ctx, pair)
	if err != nil {
		return fmt.Errorf("get margin market: %v", err)
	}
	if market.Status != "enabled" {
		return fmt.Errorf("margin market %s is not enabled", pair)
	}
	amount, err := positiveAmount("minimum quote borrow amount", market.QuoteMinBorrowAmount)
	if err != nil {
		return err
	}
	// Repay-all is scoped to a pair/currency, not a loan ID. Refuse pre-existing debt to protect it.
	loans, _, err := client.MarginUniApi.ListUniLoans(ctx, &gateapi.ListUniLoansOpts{
		CurrencyPair: optional.NewString(pair), Currency: optional.NewString(currency), Limit: optional.NewInt32(1),
	})
	if err != nil {
		return fmt.Errorf("check existing margin debt: %v", err)
	}
	if len(loans) != 0 {
		return fmt.Errorf("use an isolated demo account without existing %s debt in %s", currency, pair)
	}
	borrowable, _, err := client.MarginUniApi.GetUniBorrowable(ctx, currency, pair)
	if err != nil {
		return fmt.Errorf("get margin borrowing capacity: %v", err)
	}
	capacity, err := parseAmount("borrowable amount", borrowable.Borrowable)
	if err != nil {
		return err
	}
	if capacity.LessThan(amount) {
		return fmt.Errorf("insufficient margin collateral; fund the demo account first")
	}
	request := gateapi.CreateUniLoan{CurrencyPair: pair, Currency: currency, Amount: amount.String(), Type: "borrow"}
	if _, err = client.MarginUniApi.CreateUniLoan(ctx, request); err != nil {
		// A transport error can follow a successful borrow; do not automatically repeat a financial write.
		return fmt.Errorf("borrow outcome requires checking %s/%s loan history before retrying: %v", pair, currency, err)
	}
	// Keep borrowed funds available for repayment; the old sell/lend flow used removed loan APIs.
	request.Type, request.RepaidAll = "repay", true
	if _, err = client.MarginUniApi.CreateUniLoan(ctx, request); err != nil {
		return fmt.Errorf("repay failed; inspect and repay %s/%s debt including accrued interest: %v", pair, currency, err)
	}
	logger.Printf("borrowed and repaid %s %s in %s", amount.String(), currency, pair)
	return nil
}
