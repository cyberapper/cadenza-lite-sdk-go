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

func TestTradingOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Trading.Order.New(context.TODO(), cadenzalite.TradingOrderNewParams{
		ExchangeAccountID:      cadenzalite.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Leverage:               cadenzalite.F(int64(0)),
		OrderSide:              cadenzalite.F(cadenzalite.TradingOrderNewParamsOrderSideBuy),
		OrderType:              cadenzalite.F(cadenzalite.TradingOrderNewParamsOrderTypeMarket),
		PositionID:             cadenzalite.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Price:                  cadenzalite.F(0.000000),
		PriceSlippageTolerance: cadenzalite.F(0.000000),
		Priority:               cadenzalite.F([]string{"exchange_account_id_1", "exchange_account_id_2", "exchange_account_id_3"}),
		Quantity:               cadenzalite.F(0.000000),
		QuoteID:                cadenzalite.F("quoteId"),
		QuoteQuantity:          cadenzalite.F(0.000000),
		QuoteRequestID:         cadenzalite.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RoutePolicy:            cadenzalite.F(cadenzalite.TradingOrderNewParamsRoutePolicyPriority),
		Symbol:                 cadenzalite.F("BTC/USDT"),
		TenantID:               cadenzalite.F("tenantId"),
		TimeInForce:            cadenzalite.F(cadenzalite.TradingOrderNewParamsTimeInForceDay),
		IdempotencyKey:         cadenzalite.F("my_idempotency_key"),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTradingOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Trading.Order.List(context.TODO(), cadenzalite.TradingOrderListParams{
		EndTime:           cadenzalite.F(int64(1632933600000)),
		ExchangeAccountID: cadenzalite.F("exchangeAccountId"),
		Limit:             cadenzalite.F(int64(100)),
		Offset:            cadenzalite.F(int64(0)),
		OrderID:           cadenzalite.F("orderId"),
		OrderStatus:       cadenzalite.F(cadenzalite.TradingOrderListParamsOrderStatusSubmitted),
		StartTime:         cadenzalite.F(int64(1622505600000)),
		Symbol:            cadenzalite.F("symbol"),
		TenantID:          cadenzalite.F("tenantId"),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTradingOrderCancel(t *testing.T) {
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
	_, err := client.Trading.Order.Cancel(context.TODO(), cadenzalite.TradingOrderCancelParams{
		OrderID: cadenzalite.F("orderId"),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
