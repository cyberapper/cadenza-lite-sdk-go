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
		DropCopyOrder: cadenzasdk.DropCopyOrderParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskPlaceOrderRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.PlaceOrderRequestParam{
				Cost:              int64(0),
				CreatedAt:         int64(1703052635110),
				ExchangeType:      "BINANCE",
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Filled:            int64(0),
				OrderID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				OrderSide:         cadenzasdk.F(cadenzasdk.PlaceOrderRequestOrderSideBuy),
				OrderType:         cadenzasdk.F(cadenzasdk.PlaceOrderRequestOrderTypeMarket),
				PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Price:             cadenzasdk.F(0.000000),
				Quantity:          cadenzasdk.F(0.000000),
				QuoteQuantity:     cadenzasdk.F(0.000000),
				Status:            "SUBMITTED",
				Symbol:            cadenzasdk.F("BTC/USDT"),
				TimeInForce:       cadenzasdk.F(cadenzasdk.PlaceOrderRequestTimeInForceDay),
				UpdatedAt:         int64(1703052635111),
				UserID:            "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Fee:               int64(0),
				FeeCurrency:       "USDT",
				TenantID:          cadenzasdk.F("tenantId"),
			},
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
		DropCopyPortfolio: cadenzasdk.DropCopyPortfolioParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskPlaceOrderRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.PlaceOrderRequestParam{
				Payload: map[string]interface{}{
					"exchangeAccountId": "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					"exchangeType":      "BINANCE",
					"balances": map[string]interface{}{
						"0": map[string]interface{}{
							"asset":  "BTC",
							"free":   int64(1),
							"locked": int64(0),
							"total":  int64(1),
						},
						"1": map[string]interface{}{
							"asset":  "BTC",
							"free":   int64(1),
							"locked": int64(0),
							"total":  int64(1),
						},
						"2": map[string]interface{}{
							"asset":  "BTC",
							"free":   int64(1),
							"locked": int64(0),
							"total":  int64(1),
						},
					},
					"positions": map[string]interface{}{
						"0": map[string]interface{}{
							"amount":       int64(0),
							"cost":         int64(0),
							"entryPrice":   int64(0),
							"positionSide": "LONG",
							"status":       "OPEN",
							"symbol":       "BTC/USDT",
						},
						"1": map[string]interface{}{
							"amount":       int64(0),
							"cost":         int64(0),
							"entryPrice":   int64(0),
							"positionSide": "LONG",
							"status":       "OPEN",
							"symbol":       "BTC/USDT",
						},
						"2": map[string]interface{}{
							"amount":       int64(0),
							"cost":         int64(0),
							"entryPrice":   int64(0),
							"positionSide": "LONG",
							"status":       "OPEN",
							"symbol":       "BTC/USDT",
						},
					},
					"credit": map[string]interface{}{
						"exchangeAccountId": "018e41a1-cebc-7b49-a729-ae2c1c41e297",
						"exchangeType":      "BINANCE",
						"accountType":       "SPOT",
						"currency":          "USDT",
						"leverage":          int64(1),
						"credit":            int64(10000),
						"margin":            int64(5000),
						"marginLoan":        int64(3000),
						"marginRequirement": int64(1500),
						"marginUsage":       0.500000,
						"marginLevel":       0.890000,
						"riskExposure":      5677517.760000,
						"maxRiskExposure":   int64(5000000),
						"riskExposureRate":  0.890000,
					},
					"updatedAt": int64(1632933600000),
				},
			},
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
		DropCopyQuote: cadenzasdk.DropCopyQuoteParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskPlaceOrderRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.PlaceOrderRequestParam{
				QuoteRequestID:    cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				BaseCurrency:      "BTC",
				QuoteCurrency:     "USDT",
				AskPrice:          42859.990000,
				AskQuantity:       int64(1),
				BidPrice:          42859.710000,
				BidQuantity:       int64(1),
				Timestamp:         int64(1632933600000),
				ValidUntil:        int64(1632933600000),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      "BINANCE",
			},
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
		MarketDataKline: cadenzasdk.MarketDataKlineParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskPlaceOrderRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.PlaceOrderRequestParam{
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      "BINANCE",
				Symbol:            cadenzasdk.F("symbol"),
				Interval:          "1s",
				Candles: map[string]interface{}{
					"0": map[string]interface{}{
						"c": int64(0),
						"h": int64(0),
						"l": int64(0),
						"o": int64(0),
						"t": int64(0),
						"v": int64(0),
					},
					"1": map[string]interface{}{
						"c": int64(0),
						"h": int64(0),
						"l": int64(0),
						"o": int64(0),
						"t": int64(0),
						"v": int64(0),
					},
					"2": map[string]interface{}{
						"c": int64(0),
						"h": int64(0),
						"l": int64(0),
						"o": int64(0),
						"t": int64(0),
						"v": int64(0),
					},
				},
			},
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
		MarketDataOrderBook: cadenzasdk.MarketDataOrderBookParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskPlaceOrderRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.PlaceOrderRequestParam{
				Asks: map[string]interface{}{
					"0": map[string]interface{}{
						"0": int64(0),
						"1": int64(0),
					},
					"1": map[string]interface{}{
						"0": int64(0),
						"1": int64(0),
					},
					"2": map[string]interface{}{
						"0": int64(0),
						"1": int64(0),
					},
				},
				Bids: map[string]interface{}{
					"0": map[string]interface{}{
						"0": int64(0),
						"1": int64(0),
					},
					"1": map[string]interface{}{
						"0": int64(0),
						"1": int64(0),
					},
					"2": map[string]interface{}{
						"0": int64(0),
						"1": int64(0),
					},
				},
				ExchangeType:      "exchangeType",
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Symbol:            cadenzasdk.F("symbol"),
				Level:             int64(0),
			},
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
