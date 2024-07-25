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
		PlaceOrderRequest: cadenzasdk.PlaceOrderRequestParam{
			QuoteRequestID:         cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ExchangeAccountID:      cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Leverage:               cadenzasdk.F(int64(0)),
			OrderSide:              cadenzasdk.F(cadenzasdk.PlaceOrderRequestOrderSideBuy),
			OrderType:              cadenzasdk.F(cadenzasdk.PlaceOrderRequestOrderTypeMarket),
			PositionID:             cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Price:                  cadenzasdk.F(0.000000),
			PriceSlippageTolerance: cadenzasdk.F(0.000000),
			Quantity:               cadenzasdk.F(0.000000),
			QuoteQuantity:          cadenzasdk.F(0.000000),
			Symbol:                 cadenzasdk.F("BTC/USDT"),
			TimeInForce:            cadenzasdk.F(cadenzasdk.PlaceOrderRequestTimeInForceDay),
			RoutePolicy:            cadenzasdk.F(cadenzasdk.PlaceOrderRequestRoutePolicyPriority),
			Priority:               cadenzasdk.F([]string{"exchange_account_id_1", "exchange_account_id_2", "exchange_account_id_3"}),
			QuoteID:                cadenzasdk.F("quoteId"),
			TenantID:               cadenzasdk.F("tenantId"),
		},
		IdempotencyKey: cadenzasdk.F("my_idempotency_key"),
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
		CancelOrderRequest: cadenzasdk.CancelOrderRequestParam{
			OrderID: cadenzasdk.F("orderId"),
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
