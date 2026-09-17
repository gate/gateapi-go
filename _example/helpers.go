package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// parseAmount keeps API decimal strings exact and rejects malformed upstream values.
func parseAmount(name, value string) (decimal.Decimal, error) {
	if len(value) == 0 || len(value) > 128 {
		return decimal.Zero, fmt.Errorf("invalid %s length", name)
	}
	amount, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero, fmt.Errorf("invalid %s: %v", name, err)
	}
	// Bound exponent-driven arithmetic before multiplying amounts supplied by the API.
	if amount.Exponent() < -30 || amount.Exponent() > 30 {
		return decimal.Zero, fmt.Errorf("unsupported %s precision", name)
	}
	return amount, nil
}

// positiveAmount validates divisors, prices and order minima before any write occurs.
func positiveAmount(name, value string) (decimal.Decimal, error) {
	amount, err := parseAmount(name, value)
	if err != nil {
		return decimal.Zero, err
	}
	if !amount.IsPositive() {
		return decimal.Zero, fmt.Errorf("%s must be positive", name)
	}
	return amount, nil
}

// futuresAmounts preserves fractional contract sizes and keeps collateral positive for shorts.
func futuresAmounts(position, minimum, price, multiplier, leverage string) (decimal.Decimal, decimal.Decimal, error) {
	positionSize, err := parseAmount("position size", position)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	minSize, err := positiveAmount("minimum order size", minimum)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	last, err := positiveAmount("last price", price)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	quanto, err := positiveAmount("quanto multiplier", multiplier)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	lev, err := positiveAmount("leverage", leverage)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	size := decimal.NewFromInt(10)
	if size.LessThan(minSize) {
		size = minSize
	}
	margin := size.Mul(last).Mul(quanto).DivRound(lev, 8).Mul(decimal.New(11, -1)).RoundCeil(8)
	if !margin.IsPositive() {
		return decimal.Zero, decimal.Zero, fmt.Errorf("calculated margin must be positive")
	}
	if positionSize.IsNegative() {
		size = size.Neg()
	}
	return size, margin, nil
}
