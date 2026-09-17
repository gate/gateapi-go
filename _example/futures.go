package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/antihax/optional"
	"github.com/gate/gateapi-go/v7"
)

// FuturesDemo demonstrates current string-sized futures orders using one governed SDK client.
func FuturesDemo(ctx context.Context, client *gateapi.APIClient, useTestNet bool) error {
	settle, contract, leverage := "usdt", "BTC_USDT", "3"
	position, _, err := client.FuturesApi.GetPosition(ctx, settle, contract)
	if err != nil {
		return fmt.Errorf("get position: %v", err)
	}
	futuresContract, _, err := client.FuturesApi.GetFuturesContract(ctx, settle, contract)
	if err != nil {
		return fmt.Errorf("get futures contract: %v", err)
	}
	tickers, _, err := client.FuturesApi.ListFuturesTickers(ctx, settle, &gateapi.ListFuturesTickersOpts{Contract: optional.NewString(contract)})
	if err != nil {
		return fmt.Errorf("list futures tickers: %v", err)
	}
	if len(tickers) == 0 {
		return fmt.Errorf("no ticker for %s", contract)
	}
	size, margin, err := futuresAmounts(position.Size, futuresContract.OrderSizeMin, tickers[0].Last, futuresContract.QuantoMultiplier, leverage)
	if err != nil {
		return err
	}
	account, _, err := client.FuturesApi.ListFuturesAccounts(ctx, settle)
	if err != nil {
		return fmt.Errorf("get futures account: %v", err)
	}
	available, err := parseAmount("available collateral", account.Available)
	if err != nil {
		return err
	}
	if available.IsNegative() {
		return fmt.Errorf("available collateral must not be negative")
	}
	// Validate all read data first; only transfer the shortfall, with an explicit settlement currency.
	if margin.GreaterThan(available) {
		if useTestNet {
			return fmt.Errorf("testnet collateral insufficient; fund the account before running the demo")
		}
		_, _, err = client.WalletApi.Transfer(ctx, gateapi.Transfer{
			Currency: strings.ToUpper(settle), From: "spot", To: "futures", Settle: settle,
			Amount: margin.Sub(available).String(),
		})
		if err != nil {
			return fmt.Errorf("transfer collateral (check transfer history before retrying): %v", err)
		}
	}
	if _, _, err = client.FuturesApi.UpdatePositionLeverage(ctx, settle, contract, leverage, nil); err != nil {
		return fmt.Errorf("update leverage: %v", err)
	}
	// Contract is now an optional query field; omitting it would widen cancellation to every contract.
	if _, _, err = client.FuturesApi.CancelFuturesOrders(ctx, settle, &gateapi.CancelFuturesOrdersOpts{Contract: optional.NewString(contract)}); err != nil {
		return fmt.Errorf("cancel contract orders: %v", err)
	}
	order := gateapi.FuturesOrder{Contract: contract, Size: size.String(), Price: "0", Tif: "ioc"}
	created, _, err := client.FuturesApi.CreateFuturesOrder(ctx, settle, order, nil)
	if err != nil {
		return fmt.Errorf("create futures order (check order history before retrying): %v", err)
	}
	logger.Printf("futures order %d status=%s size=%s", created.Id, created.Status, created.Size)
	if created.Status == "open" {
		id := strconv.FormatInt(created.Id, 10)
		// Cancel a remaining open order immediately; a failed status read must not prevent cleanup.
		if _, _, err = client.FuturesApi.CancelFuturesOrder(ctx, settle, id, nil); err != nil {
			return fmt.Errorf("cancel remaining futures order %s manually if needed: %v", id, err)
		}
	}
	return nil
}
