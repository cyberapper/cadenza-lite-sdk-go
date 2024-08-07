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

func TestEventMarketDataMarketDataKlineWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.MarketData.MarketDataKline(context.TODO(), cadenzasdk.EventMarketDataMarketDataKlineParams{
		MarketDataKline: cadenzasdk.MarketDataKlineParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.MarketDataKlineEventTypeCadenzaTaskQuote),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.KlineParam{
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.KlineExchangeTypeBinance),
				Symbol:            cadenzasdk.F("symbol"),
				Interval:          cadenzasdk.F(cadenzasdk.KlineInterval1s),
				Candles: cadenzasdk.F([]cadenzasdk.BalanceEntryParam{{
					Asset:    cadenzasdk.F("BTC"),
					Free:     cadenzasdk.F(1.000000),
					Locked:   cadenzasdk.F(0.000000),
					Borrowed: cadenzasdk.F(3.000000),
					Net:      cadenzasdk.F(-2.000000),
					Total:    cadenzasdk.F(1.000000),
				}, {
					Asset:    cadenzasdk.F("BTC"),
					Free:     cadenzasdk.F(1.000000),
					Locked:   cadenzasdk.F(0.000000),
					Borrowed: cadenzasdk.F(3.000000),
					Net:      cadenzasdk.F(-2.000000),
					Total:    cadenzasdk.F(1.000000),
				}, {
					Asset:    cadenzasdk.F("BTC"),
					Free:     cadenzasdk.F(1.000000),
					Locked:   cadenzasdk.F(0.000000),
					Borrowed: cadenzasdk.F(3.000000),
					Net:      cadenzasdk.F(-2.000000),
					Total:    cadenzasdk.F(1.000000),
				}}),
			}),
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

func TestEventMarketDataMarketDataOrderBookWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.MarketData.MarketDataOrderBook(context.TODO(), cadenzasdk.EventMarketDataMarketDataOrderBookParams{
		MarketDataOrderBook: cadenzasdk.MarketDataOrderBookParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.MarketDataOrderBookEventTypeCadenzaTaskQuote),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.OrderbookParam{
				Asks:              cadenzasdk.F([][]float64{{0.000000, 0.000000}, {0.000000, 0.000000}, {0.000000, 0.000000}}),
				Bids:              cadenzasdk.F([][]float64{{0.000000, 0.000000}, {0.000000, 0.000000}, {0.000000, 0.000000}}),
				ExchangeType:      cadenzasdk.F("exchangeType"),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Symbol:            cadenzasdk.F("symbol"),
				Level:             cadenzasdk.F(int64(0)),
			}),
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
