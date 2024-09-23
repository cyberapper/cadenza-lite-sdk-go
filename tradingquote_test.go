// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cyberapper/cadenza-lite-sdk-go"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/testutil"
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
)

func TestTradingQuotePostWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenzasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Trading.Quote.Post(context.TODO(), cadenzasdk.TradingQuotePostParams{
		QuoteRequest: cadenzasdk.QuoteRequestParam{
			BaseCurrency:      cadenzasdk.F("baseCurrency"),
			OrderSide:         cadenzasdk.F("orderSide"),
			QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
			ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Quantity:          cadenzasdk.F(0.000000),
			QuoteQuantity:     cadenzasdk.F(0.000000),
		},
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTradingQuoteRequestForQuoteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenzasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Trading.Quote.RequestForQuote(context.TODO(), cadenzasdk.TradingQuoteRequestForQuoteParams{
		QuoteRequest: cadenzasdk.QuoteRequestParam{
			BaseCurrency:      cadenzasdk.F("baseCurrency"),
			OrderSide:         cadenzasdk.F("orderSide"),
			QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
			ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Quantity:          cadenzasdk.F(0.000000),
			QuoteQuantity:     cadenzasdk.F(0.000000),
		},
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
