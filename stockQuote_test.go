package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStockQuote(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		if r.URL.Path != "/stock/aapl/quote" {
			t.Errorf("Expected request URL '/stock/aapl/quote', got '%s'", r.URL.Path)
		}

		// Send a mock response
		response := `{"symbol":"AAPL","price":135.25}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer mockServer.Close()

	// Replace the original base URL with the mock server URL
	baseURL = mockServer.URL

	// Perform the test
	quote, err := getStockQuote("AAPL")
	if err != nil {
		t.Errorf("Error while getting stock quote: %s", err)
	}

	// Verify the results
	expected := StockQuote{Symbol: "AAPL", Price: 135.25}
	if quote != expected {
		t.Errorf("Expected quote %+v, got %+v", expected, quote)
	}
}
