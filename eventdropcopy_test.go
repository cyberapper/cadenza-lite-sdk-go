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

func TestEventDropCopyDropCopyCancelOrderRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopy.DropCopyCancelOrderRequestAck(context.TODO(), cadenzasdk.EventDropCopyDropCopyCancelOrderRequestAckParams{
		DropCopyCancelOrderRequestAck: cadenzasdk.DropCopyCancelOrderRequestAckParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.DropCopyCancelOrderRequestAckEventTypeCadenzaTaskQuote),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.CancelOrderRequestParam{
				OrderID: cadenzasdk.F("orderId"),
			}),
			Source: cadenzasdk.F("source"),
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

func TestEventDropCopyDropCopyExecutionReportWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopy.DropCopyExecutionReport(context.TODO(), cadenzasdk.EventDropCopyDropCopyExecutionReportParams{
		DropCopyExecutionReport: cadenzasdk.DropCopyExecutionReportParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.DropCopyExecutionReportEventTypeCadenzaTaskQuote),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.ExecutionReportParam{
				BaseCurrency:  cadenzasdk.F("BTC"),
				Cost:          cadenzasdk.F(42859.990000),
				CreatedAt:     cadenzasdk.F(int64(1632933600000)),
				Filled:        cadenzasdk.F(1.000000),
				QuoteCurrency: cadenzasdk.F("USDT"),
				RoutePolicy:   cadenzasdk.F(cadenzasdk.ExecutionReportRoutePolicyPriority),
				Status:        cadenzasdk.F(cadenzasdk.ExecutionReportStatusSubmitted),
				UpdatedAt:     cadenzasdk.F(int64(1632933600000)),
				ID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Executions: cadenzasdk.F([]cadenzasdk.OrderParam{{
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					Filled:            cadenzasdk.F(0.000000),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					Quantity:          cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					BaseCurrency:      cadenzasdk.F("BTC"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					QuoteCurrency:     cadenzasdk.F("USDT"),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					TenantID:          cadenzasdk.F("tenantId"),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}, {
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					Filled:            cadenzasdk.F(0.000000),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					Quantity:          cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					BaseCurrency:      cadenzasdk.F("BTC"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					QuoteCurrency:     cadenzasdk.F("USDT"),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					TenantID:          cadenzasdk.F("tenantId"),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}, {
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					Filled:            cadenzasdk.F(0.000000),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					Quantity:          cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					BaseCurrency:      cadenzasdk.F("BTC"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					QuoteCurrency:     cadenzasdk.F("USDT"),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					TenantID:          cadenzasdk.F("tenantId"),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}}),
				Fees: cadenzasdk.F([]cadenzasdk.ExecutionReportFeeParam{{
					Asset:    cadenzasdk.F("asset"),
					Quantity: cadenzasdk.F(0.000000),
				}, {
					Asset:    cadenzasdk.F("asset"),
					Quantity: cadenzasdk.F(0.000000),
				}, {
					Asset:    cadenzasdk.F("asset"),
					Quantity: cadenzasdk.F(0.000000),
				}}),
				Order: cadenzasdk.F(cadenzasdk.OrderParam{
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					Filled:            cadenzasdk.F(0.000000),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					Quantity:          cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					BaseCurrency:      cadenzasdk.F("BTC"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					QuoteCurrency:     cadenzasdk.F("USDT"),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					TenantID:          cadenzasdk.F("tenantId"),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}),
			}),
			Source: cadenzasdk.F("source"),
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

func TestEventDropCopyDropCopyOrderWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopy.DropCopyOrder(context.TODO(), cadenzasdk.EventDropCopyDropCopyOrderParams{
		DropCopyOrder: cadenzasdk.DropCopyOrderParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.DropCopyOrderEventTypeCadenzaTaskQuote),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.OrderParam{
				Cost:              cadenzasdk.F(0.000000),
				CreatedAt:         cadenzasdk.F(int64(1703052635110)),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
				Filled:            cadenzasdk.F(0.000000),
				OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
				OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
				Quantity:          cadenzasdk.F(0.000000),
				Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
				Symbol:            cadenzasdk.F("BTC/USDT"),
				TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
				UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
				BaseCurrency:      cadenzasdk.F("BTC"),
				Fee:               cadenzasdk.F(0.000000),
				FeeCurrency:       cadenzasdk.F("USDT"),
				OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Price:             cadenzasdk.F(0.000000),
				QuoteCurrency:     cadenzasdk.F("USDT"),
				QuoteQuantity:     cadenzasdk.F(0.000000),
				TenantID:          cadenzasdk.F("tenantId"),
				UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			Source: cadenzasdk.F("source"),
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

func TestEventDropCopyDropCopyPlaceOrderRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopy.DropCopyPlaceOrderRequestAck(context.TODO(), cadenzasdk.EventDropCopyDropCopyPlaceOrderRequestAckParams{
		DropCopyPlaceOrderRequestAck: cadenzasdk.DropCopyPlaceOrderRequestAckParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.DropCopyPlaceOrderRequestAckEventTypeCadenzaTaskQuote),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.PlaceOrderRequestParam{
				RoutePolicy:            cadenzasdk.F(cadenzasdk.PlaceOrderRequestRoutePolicyPriority),
				ExchangeAccountID:      cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Leverage:               cadenzasdk.F(int64(0)),
				OrderSide:              cadenzasdk.F(cadenzasdk.PlaceOrderRequestOrderSideBuy),
				OrderType:              cadenzasdk.F(cadenzasdk.PlaceOrderRequestOrderTypeMarket),
				PositionID:             cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Price:                  cadenzasdk.F(0.000000),
				PriceSlippageTolerance: cadenzasdk.F(0.000000),
				Priority:               cadenzasdk.F([]string{"exchange_account_id_1", "exchange_account_id_2", "exchange_account_id_3"}),
				Quantity:               cadenzasdk.F(0.000000),
				QuoteID:                cadenzasdk.F("quoteId"),
				QuoteQuantity:          cadenzasdk.F(0.000000),
				QuoteRequestID:         cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Symbol:                 cadenzasdk.F("BTC/USDT"),
				TenantID:               cadenzasdk.F("tenantId"),
				TimeInForce:            cadenzasdk.F(cadenzasdk.PlaceOrderRequestTimeInForceDay),
			}),
			Source: cadenzasdk.F("source"),
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

func TestEventDropCopyDropCopyPortfolioWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopy.DropCopyPortfolio(context.TODO(), cadenzasdk.EventDropCopyDropCopyPortfolioParams{
		DropCopyPortfolio: cadenzasdk.DropCopyPortfolioParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.DropCopyPortfolioEventTypeCadenzaTaskQuote),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioParam{
				Credit: cadenzasdk.F(cadenzasdk.ExchangeAccountCreditParam{
					AccountType:       cadenzasdk.F(cadenzasdk.ExchangeAccountCreditAccountTypeSpot),
					Credit:            cadenzasdk.F(10000.000000),
					Currency:          cadenzasdk.F("USDT"),
					ExchangeAccountID: cadenzasdk.F("018e41a1-cebc-7b49-a729-ae2c1c41e297"),
					ExchangeType:      cadenzasdk.F(cadenzasdk.ExchangeAccountCreditExchangeTypeBinance),
					Leverage:          cadenzasdk.F(int64(1)),
					Margin:            cadenzasdk.F(5000.000000),
					MarginLevel:       cadenzasdk.F(0.890000),
					MarginLoan:        cadenzasdk.F(3000.000000),
					MarginRequirement: cadenzasdk.F(1500.000000),
					MarginUsage:       cadenzasdk.F(0.500000),
					MaxRiskExposure:   cadenzasdk.F(5000000.000000),
					RiskExposure:      cadenzasdk.F(5677517.760000),
					RiskExposureRate:  cadenzasdk.F(0.890000),
				}),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioExchangeTypeBinance),
				UpdatedAt:         cadenzasdk.F(int64(1632933600000)),
				Balances: cadenzasdk.F([]cadenzasdk.BalanceEntryParam{{
					Asset:    cadenzasdk.F("BTC"),
					Borrowed: cadenzasdk.F(3.000000),
					Free:     cadenzasdk.F(1.000000),
					Locked:   cadenzasdk.F(0.000000),
					Net:      cadenzasdk.F(-2.000000),
					Total:    cadenzasdk.F(1.000000),
				}, {
					Asset:    cadenzasdk.F("BTC"),
					Borrowed: cadenzasdk.F(3.000000),
					Free:     cadenzasdk.F(1.000000),
					Locked:   cadenzasdk.F(0.000000),
					Net:      cadenzasdk.F(-2.000000),
					Total:    cadenzasdk.F(1.000000),
				}, {
					Asset:    cadenzasdk.F("BTC"),
					Borrowed: cadenzasdk.F(3.000000),
					Free:     cadenzasdk.F(1.000000),
					Locked:   cadenzasdk.F(0.000000),
					Net:      cadenzasdk.F(-2.000000),
					Total:    cadenzasdk.F(1.000000),
				}}),
				Positions: cadenzasdk.F([]cadenzasdk.PositionEntryParam{{
					Amount:       cadenzasdk.F(0.000000),
					PositionSide: cadenzasdk.F(cadenzasdk.PositionEntryPositionSideLong),
					Status:       cadenzasdk.F(cadenzasdk.PositionEntryStatusOpen),
					Symbol:       cadenzasdk.F("BTC/USDT"),
					Cost:         cadenzasdk.F(0.000000),
					EntryPrice:   cadenzasdk.F(0.000000),
				}, {
					Amount:       cadenzasdk.F(0.000000),
					PositionSide: cadenzasdk.F(cadenzasdk.PositionEntryPositionSideLong),
					Status:       cadenzasdk.F(cadenzasdk.PositionEntryStatusOpen),
					Symbol:       cadenzasdk.F("BTC/USDT"),
					Cost:         cadenzasdk.F(0.000000),
					EntryPrice:   cadenzasdk.F(0.000000),
				}, {
					Amount:       cadenzasdk.F(0.000000),
					PositionSide: cadenzasdk.F(cadenzasdk.PositionEntryPositionSideLong),
					Status:       cadenzasdk.F(cadenzasdk.PositionEntryStatusOpen),
					Symbol:       cadenzasdk.F("BTC/USDT"),
					Cost:         cadenzasdk.F(0.000000),
					EntryPrice:   cadenzasdk.F(0.000000),
				}}),
			}),
			Source: cadenzasdk.F("source"),
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

func TestEventDropCopyDropCopyQuoteWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopy.DropCopyQuote(context.TODO(), cadenzasdk.EventDropCopyDropCopyQuoteParams{
		DropCopyQuote: cadenzasdk.DropCopyQuoteParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.DropCopyQuoteEventTypeCadenzaTaskQuote),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.QuoteParam{
				BaseCurrency:      cadenzasdk.F("BTC"),
				QuoteCurrency:     cadenzasdk.F("USDT"),
				QuoteRequestID:    cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Timestamp:         cadenzasdk.F(int64(1632933600000)),
				ValidUntil:        cadenzasdk.F(int64(1632933600000)),
				AskPrice:          cadenzasdk.F(42859.990000),
				AskQuantity:       cadenzasdk.F(1.000000),
				BidPrice:          cadenzasdk.F(42859.710000),
				BidQuantity:       cadenzasdk.F(1.000000),
				CreatedAt:         cadenzasdk.F(int64(1632933600000)),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.QuoteExchangeTypeBinance),
				ExpiredAt:         cadenzasdk.F(int64(1632933600000)),
			}),
			Source: cadenzasdk.F("source"),
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

func TestEventDropCopyDropCopyQuoteRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopy.DropCopyQuoteRequestAck(context.TODO(), cadenzasdk.EventDropCopyDropCopyQuoteRequestAckParams{
		DropCopyRequestAck: cadenzasdk.DropCopyRequestAckParam{
			EventID:   cadenzasdk.F("eventId"),
			EventType: cadenzasdk.F(cadenzasdk.DropCopyRequestAckEventTypeCadenzaTaskQuote),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.QuoteRequestParam{
				BaseCurrency:      cadenzasdk.F("baseCurrency"),
				OrderSide:         cadenzasdk.F("orderSide"),
				QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:          cadenzasdk.F(0.000000),
				QuoteQuantity:     cadenzasdk.F(0.000000),
			}),
			Source: cadenzasdk.F("source"),
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
