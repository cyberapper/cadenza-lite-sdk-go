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

func TestEventDropCopyOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopyOrder(context.TODO(), cadenzasdk.EventDropCopyOrderParams{
		EventID:   cadenzasdk.F("eventId"),
		EventType: cadenzasdk.F(cadenzasdk.EventDropCopyOrderParamsEventTypeCadenzaDropCopyOrder),
		Timestamp: cadenzasdk.F(int64(1632933600000)),
		Payload: cadenzasdk.F[cadenzasdk.OrderUnionParam](cadenzasdk.QuoteRequestParam{
			BaseCurrency:      cadenzasdk.F("baseCurrency"),
			QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
			OrderSide:         cadenzasdk.F("orderSide"),
			Quantity:          cadenzasdk.F(0.000000),
			QuoteQuantity:     cadenzasdk.F(0.000000),
			ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		Source: cadenzasdk.F("source"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventDropCopyPortfolioWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopyPortfolio(context.TODO(), cadenzasdk.EventDropCopyPortfolioParams{
		EventID:   cadenzasdk.F("eventId"),
		EventType: cadenzasdk.F(cadenzasdk.EventDropCopyPortfolioParamsEventTypeCadenzaDropCopyPortfolio),
		Timestamp: cadenzasdk.F(int64(1632933600000)),
		Payload: cadenzasdk.F[cadenzasdk.ExchangeAccountPortfolioUnionParam](cadenzasdk.QuoteRequestParam{
			BaseCurrency:      cadenzasdk.F("baseCurrency"),
			QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
			OrderSide:         cadenzasdk.F("orderSide"),
			Quantity:          cadenzasdk.F(0.000000),
			QuoteQuantity:     cadenzasdk.F(0.000000),
			ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		Source: cadenzasdk.F("source"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventDropCopyQuoteWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopyQuote(context.TODO(), cadenzasdk.EventDropCopyQuoteParams{
		EventID:   cadenzasdk.F("eventId"),
		EventType: cadenzasdk.F(cadenzasdk.EventDropCopyQuoteParamsEventTypeCadenzaDropCopyQuote),
		Timestamp: cadenzasdk.F(int64(1632933600000)),
		Payload: cadenzasdk.F[cadenzasdk.QuoteUnionParam](cadenzasdk.QuoteRequestParam{
			BaseCurrency:      cadenzasdk.F("baseCurrency"),
			QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
			OrderSide:         cadenzasdk.F("orderSide"),
			Quantity:          cadenzasdk.F(0.000000),
			QuoteQuantity:     cadenzasdk.F(0.000000),
			ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		Source: cadenzasdk.F("source"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventMarketDataKlineWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.MarketDataKline(context.TODO(), cadenzasdk.EventMarketDataKlineParams{
		EventID:   cadenzasdk.F("eventId"),
		EventType: cadenzasdk.F(cadenzasdk.EventMarketDataKlineParamsEventTypeCadenzaMarketDataKline),
		Timestamp: cadenzasdk.F(int64(1632933600000)),
		Payload: cadenzasdk.F[cadenzasdk.EventMarketDataKlineParamsPayloadUnion](cadenzasdk.QuoteRequestParam{
			BaseCurrency:      cadenzasdk.F("baseCurrency"),
			QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
			OrderSide:         cadenzasdk.F("orderSide"),
			Quantity:          cadenzasdk.F(0.000000),
			QuoteQuantity:     cadenzasdk.F(0.000000),
			ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		Source: cadenzasdk.F("source"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEventMarketDataOrderBookWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.MarketDataOrderBook(context.TODO(), cadenzasdk.EventMarketDataOrderBookParams{
		EventID:   cadenzasdk.F("eventId"),
		EventType: cadenzasdk.F(cadenzasdk.EventMarketDataOrderBookParamsEventTypeCadenzaMarketDataOrderBook),
		Timestamp: cadenzasdk.F(int64(1632933600000)),
		Payload: cadenzasdk.F[cadenzasdk.OrderbookUnionParam](cadenzasdk.QuoteRequestParam{
			BaseCurrency:      cadenzasdk.F("baseCurrency"),
			QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
			OrderSide:         cadenzasdk.F("orderSide"),
			Quantity:          cadenzasdk.F(0.000000),
			QuoteQuantity:     cadenzasdk.F(0.000000),
			ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		Source: cadenzasdk.F("source"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
