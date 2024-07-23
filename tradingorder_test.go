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

func TestTradingOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Trading.Order.New(context.TODO(), cadenzasdk.TradingOrderNewParams{
		ExchangeAccountID:      cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Leverage:               cadenzasdk.F(int64(0)),
		OrderSide:              cadenzasdk.F(cadenzasdk.TradingOrderNewParamsOrderSideBuy),
		OrderType:              cadenzasdk.F(cadenzasdk.TradingOrderNewParamsOrderTypeMarket),
		PositionID:             cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Price:                  cadenzasdk.F(0.000000),
		PriceSlippageTolerance: cadenzasdk.F(0.000000),
		Priority:               cadenzasdk.F([]string{"exchange_account_id_1", "exchange_account_id_2", "exchange_account_id_3"}),
		Quantity:               cadenzasdk.F(0.000000),
		QuoteID:                cadenzasdk.F("quoteId"),
		QuoteQuantity:          cadenzasdk.F(0.000000),
		QuoteRequestID:         cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RoutePolicy:            cadenzasdk.F(cadenzasdk.TradingOrderNewParamsRoutePolicyPriority),
		Symbol:                 cadenzasdk.F("BTC/USDT"),
		TenantID:               cadenzasdk.F("tenantId"),
		TimeInForce:            cadenzasdk.F(cadenzasdk.TradingOrderNewParamsTimeInForceDay),
		IdempotencyKey:         cadenzasdk.F("my_idempotency_key"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
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
	client := cadenzasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Trading.Order.List(context.TODO(), cadenzasdk.TradingOrderListParams{
		EndTime:           cadenzasdk.F(int64(1632933600000)),
		ExchangeAccountID: cadenzasdk.F("exchangeAccountId"),
		Limit:             cadenzasdk.F(int64(100)),
		Offset:            cadenzasdk.F(int64(0)),
		OrderID:           cadenzasdk.F("orderId"),
		OrderStatus:       cadenzasdk.F(cadenzasdk.TradingOrderListParamsOrderStatusSubmitted),
		StartTime:         cadenzasdk.F(int64(1622505600000)),
		Symbol:            cadenzasdk.F("symbol"),
		TenantID:          cadenzasdk.F("tenantId"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
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
	client := cadenzasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Trading.Order.Cancel(context.TODO(), cadenzasdk.TradingOrderCancelParams{
		OrderID: cadenzasdk.F("orderId"),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
