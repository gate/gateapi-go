# Go SDK examples

These examples target the **SDK variant generated alongside them**. Their authoritative source is
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

One SDK client owns authentication and a 10-second HTTP timeout; the complete
run has a one-minute deadline. Writes are never retried automatically. If a transfer,
order or borrow times out, inspect its history before rerunning. A repayment failure
reports the affected market/currency for manual recovery; a successful borrow is not
transactionally rolled back by the API.

## Build and verify without trading

From the generated SDK root, compile the library first and bind the example to
that exact local SDK for verification. The example has its own dependency module;
the library module and its Go 1.12 declaration remain unchanged. The temporary
`replace` is restored on exit and must never be committed or published.

```sh
# Check the library without adding example-only dependencies.
go vet ./... && go test ./...
(
  cd _example
  cp go.mod go.mod.verify-backup
  trap 'mv go.mod.verify-backup go.mod' EXIT
  # Validate this unpublished candidate, rather than fetching an older remote SDK.
  GOWORK=off go mod edit -replace="$(awk '/^module / {print $2}' ../go.mod)=../"
  GOWORK=off go vet ./... && GOWORK=off go test -race ./... && GOWORK=off go build -o /tmp/gateapi-demo .
)
```

The public and filtered variants use API key/secret signing. The complete internal
variant imports its own generated module and requires the trusted gateway's
`ContextAPIKey` through `runWithContext`; its mock test shows that integration.
The internal CLI rejects public key/secret credentials before making any request.
It does not reinterpret an API key as an internal user identity.

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
