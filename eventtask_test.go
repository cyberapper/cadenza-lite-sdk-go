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

func TestEventTaskTaskCancelOrderRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.Task.TaskCancelOrderRequestAck(context.TODO(), cadenzasdk.EventTaskTaskCancelOrderRequestAckParams{
		TaskCancelOrderRequestAck: cadenzasdk.TaskCancelOrderRequestAckParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.TaskCancelOrderRequestAckEventTypeCadenzaTaskQuoteRequestAck),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.CancelOrderRequestParam{
				OrderID: cadenzasdk.F("orderId"),
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

func TestEventTaskTaskPlaceOrderRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.Task.TaskPlaceOrderRequestAck(context.TODO(), cadenzasdk.EventTaskTaskPlaceOrderRequestAckParams{
		TaskPlaceOrderRequestAck: cadenzasdk.TaskPlaceOrderRequestAckParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.TaskPlaceOrderRequestAckEventTypeCadenzaTaskQuoteRequestAck),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.PlaceOrderRequestParam{
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

func TestEventTaskTaskQuoteRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.Task.TaskQuoteRequestAck(context.TODO(), cadenzasdk.EventTaskTaskQuoteRequestAckParams{
		TaskQuoteRequestAck: cadenzasdk.TaskQuoteRequestAckParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.TaskQuoteRequestAckEventTypeCadenzaTaskQuoteRequestAck),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.QuoteRequestParam{
				BaseCurrency:      cadenzasdk.F("baseCurrency"),
				QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
				OrderSide:         cadenzasdk.F("orderSide"),
				Quantity:          cadenzasdk.F(0.000000),
				QuoteQuantity:     cadenzasdk.F(0.000000),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
