// stockQuote_test.go

package main

import (
	"testing"
)

func TestCurrentMarketData(t *testing.T) {
	symbol, price, name := currentMarketData("AAPL")

	// Verify the results
	expectedSymbol := "AAPL"
	if symbol != expectedSymbol {
		t.Errorf("Expected symbol '%s', got '%s'", expectedSymbol, symbol)
	}

	// Verify the results
	expectedPrice := 0.0 // You can update this with the expected price value
	if price != expectedPrice {
		t.Errorf("Expected price '%f', got '%f'", expectedPrice, price)
	}

	// Verify the results
	expectedName := "" // You can update this with the expected name value
	if name != expectedName {
		t.Errorf("Expected name '%s', got '%s'", expectedName, name)
	}
}
