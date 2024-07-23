// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzalite_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/cadenza-lite-go"
	"github.com/stainless-sdks/cadenza-lite-go/internal/testutil"
	"github.com/stainless-sdks/cadenza-lite-go/option"
)

func TestTradingQuoteRequestForQuoteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenzalite.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Trading.Quote.RequestForQuote(context.TODO(), cadenzalite.TradingQuoteRequestForQuoteParams{
		BaseCurrency:      cadenzalite.F("baseCurrency"),
		OrderSide:         cadenzalite.F("orderSide"),
		QuoteCurrency:     cadenzalite.F("quoteCurrency"),
		ExchangeAccountID: cadenzalite.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Quantity:          cadenzalite.F(0.000000),
		QuoteQuantity:     cadenzalite.F(0.000000),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
