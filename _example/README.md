# Go SDK examples

These examples target the **current public Go SDK**. Their authoritative source is
`codegen/demos/gate/go/` in GateAPIv4; generation copies it into `_example/`.
Change this source and regenerate rather than editing generated files.

- `spot` demonstrates a limit buy, quote-balance validation and cancellation of the
  newly created order if it remains open.
- `futures` demonstrates leverage, a collateral-shortfall transfer, cancellation
  limited to `BTC_USDT`, and a market IOC order. Decimal-string sizes are preserved;
  negative sizes do not produce negative collateral amounts.
- `margin` uses `/margin/uni/currency_pairs`, `/margin/uni/loans` and
  `/margin/uni/borrowable` to borrow and immediately repay. Removed P2P lending
  methods are not emulated. It requires an exclusively used demo account, existing
  collateral and a small quote balance to pay accrued interest. Do not run another
  trader against that account concurrently: repay-all operates on pair/currency,
  not a unique loan ID. Existing debt is rejected before borrowing.

One public SDK client owns authentication and a 10-second HTTP timeout; the complete
run has a one-minute deadline. Writes are never retried automatically. If a transfer,
order or borrow times out, inspect its history before rerunning. A repayment failure
reports the affected market/currency for manual recovery; a successful borrow is not
transactionally rolled back by the API.

## Build and verify without trading

From the generated **public SDK root**, run:

```sh
# _example is excluded by Go's ./... pattern; install its existing decimal dependency explicitly.
go get github.com/shopspring/decimal@v1.4.0
# Compile and test both the SDK and the example directory. Tests use only local HTTP mocks.
go vet ./... ./_example
go test ./... ./_example
go test -race ./_example
go build -o /tmp/gateapi-demo ./_example
```

The mock tests cover decimal quantities, signed futures size, positive collateral,
transfer shortfalls, contract/pair cancellation filters, quote/base units, current
borrow/repay request fields, existing debt, empty results, invalid input, write
failure, cancellation and deadlines. They do not prove live account permissions or
exchange availability. Do not treat these examples as a production trading strategy.

## Run deliberately

Running the binary performs real account operations. Existing flags are unchanged;
Go flags must precede the positional demo name. No new configuration is required.

```sh
# Futures testnet: fund the account first. Replace placeholders with testnet credentials.
/tmp/gateapi-demo -k <api-key> -s <api-secret> -u fx-api-testnet.gateio.ws futures
# Spot and margin default to the public production base URL; use a dedicated account.
/tmp/gateapi-demo -k <api-key> -s <api-secret> spot
/tmp/gateapi-demo -k <api-key> -s <api-secret> margin
```
