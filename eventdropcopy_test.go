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
			EventType: cadenzasdk.F(cadenzasdk.DropCopyExecutionReportEventTypeCadenzaTaskQuoteRequestAck),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.ExecutionReportParam{
				ClOrdID:       cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				BaseCurrency:  cadenzasdk.F("BTC"),
				QuoteCurrency: cadenzasdk.F("USDT"),
				RoutePolicy:   cadenzasdk.F(cadenzasdk.ExecutionReportRoutePolicyPriority),
				Order: cadenzasdk.F(cadenzasdk.OrderParam{
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Filled:            cadenzasdk.F(0.000000),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					Quantity:          cadenzasdk.F(0.000000),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					TenantID:          cadenzasdk.F("tenantId"),
				}),
				Filled: cadenzasdk.F(1.000000),
				Cost:   cadenzasdk.F(42859.990000),
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
				Status:    cadenzasdk.F(cadenzasdk.ExecutionReportStatusSubmitted),
				CreatedAt: cadenzasdk.F(int64(1632933600000)),
				UpdatedAt: cadenzasdk.F(int64(1632933600000)),
				Executions: cadenzasdk.F([]cadenzasdk.OrderParam{{
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Filled:            cadenzasdk.F(0.000000),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					Quantity:          cadenzasdk.F(0.000000),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					TenantID:          cadenzasdk.F("tenantId"),
				}, {
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Filled:            cadenzasdk.F(0.000000),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					Quantity:          cadenzasdk.F(0.000000),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					TenantID:          cadenzasdk.F("tenantId"),
				}, {
					Cost:              cadenzasdk.F(0.000000),
					CreatedAt:         cadenzasdk.F(int64(1703052635110)),
					ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Filled:            cadenzasdk.F(0.000000),
					OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
					OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
					PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Price:             cadenzasdk.F(0.000000),
					Quantity:          cadenzasdk.F(0.000000),
					QuoteQuantity:     cadenzasdk.F(0.000000),
					Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
					Symbol:            cadenzasdk.F("BTC/USDT"),
					TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
					UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
					UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Fee:               cadenzasdk.F(0.000000),
					FeeCurrency:       cadenzasdk.F("USDT"),
					TenantID:          cadenzasdk.F("tenantId"),
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
			EventType: cadenzasdk.F(cadenzasdk.DropCopyOrderEventTypeCadenzaTaskQuoteRequestAck),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.OrderParam{
				Cost:              cadenzasdk.F(0.000000),
				CreatedAt:         cadenzasdk.F(int64(1703052635110)),
				ExchangeType:      cadenzasdk.F(cadenzasdk.OrderExchangeTypeBinance),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Filled:            cadenzasdk.F(0.000000),
				OrderID:           cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				OrderSide:         cadenzasdk.F(cadenzasdk.OrderOrderSideBuy),
				OrderType:         cadenzasdk.F(cadenzasdk.OrderOrderTypeMarket),
				PositionID:        cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Price:             cadenzasdk.F(0.000000),
				Quantity:          cadenzasdk.F(0.000000),
				QuoteQuantity:     cadenzasdk.F(0.000000),
				Status:            cadenzasdk.F(cadenzasdk.OrderStatusSubmitted),
				Symbol:            cadenzasdk.F("BTC/USDT"),
				TimeInForce:       cadenzasdk.F(cadenzasdk.OrderTimeInForceDay),
				UpdatedAt:         cadenzasdk.F(int64(1703052635111)),
				UserID:            cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Fee:               cadenzasdk.F(0.000000),
				FeeCurrency:       cadenzasdk.F("USDT"),
				TenantID:          cadenzasdk.F("tenantId"),
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
			EventType: cadenzasdk.F(cadenzasdk.DropCopyPortfolioEventTypeCadenzaTaskQuoteRequestAck),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioParam{
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioExchangeTypeBinance),
				Balances: cadenzasdk.F([]cadenzasdk.ExchangeAccountPortfolioBalanceParam{{
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
				Positions: cadenzasdk.F([]cadenzasdk.ExchangeAccountPortfolioPositionParam{{
					Amount:       cadenzasdk.F(0.000000),
					Cost:         cadenzasdk.F(0.000000),
					EntryPrice:   cadenzasdk.F(0.000000),
					PositionSide: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPositionsPositionSideLong),
					Status:       cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPositionsStatusOpen),
					Symbol:       cadenzasdk.F("BTC/USDT"),
				}, {
					Amount:       cadenzasdk.F(0.000000),
					Cost:         cadenzasdk.F(0.000000),
					EntryPrice:   cadenzasdk.F(0.000000),
					PositionSide: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPositionsPositionSideLong),
					Status:       cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPositionsStatusOpen),
					Symbol:       cadenzasdk.F("BTC/USDT"),
				}, {
					Amount:       cadenzasdk.F(0.000000),
					Cost:         cadenzasdk.F(0.000000),
					EntryPrice:   cadenzasdk.F(0.000000),
					PositionSide: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPositionsPositionSideLong),
					Status:       cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPositionsStatusOpen),
					Symbol:       cadenzasdk.F("BTC/USDT"),
				}}),
				Credit: cadenzasdk.F(cadenzasdk.ExchangeAccountCreditParam{
					ExchangeAccountID: cadenzasdk.F("018e41a1-cebc-7b49-a729-ae2c1c41e297"),
					ExchangeType:      cadenzasdk.F(cadenzasdk.ExchangeAccountCreditExchangeTypeBinance),
					AccountType:       cadenzasdk.F(cadenzasdk.ExchangeAccountCreditAccountTypeSpot),
					Currency:          cadenzasdk.F("USDT"),
					Leverage:          cadenzasdk.F(int64(1)),
					Credit:            cadenzasdk.F(10000.000000),
					Margin:            cadenzasdk.F(5000.000000),
					MarginLoan:        cadenzasdk.F(3000.000000),
					MarginRequirement: cadenzasdk.F(1500.000000),
					MarginUsage:       cadenzasdk.F(0.500000),
					MarginLevel:       cadenzasdk.F(0.890000),
					RiskExposure:      cadenzasdk.F(5677517.760000),
					MaxRiskExposure:   cadenzasdk.F(5000000.000000),
					RiskExposureRate:  cadenzasdk.F(0.890000),
				}),
				UpdatedAt: cadenzasdk.F(int64(1632933600000)),
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
			EventType: cadenzasdk.F(cadenzasdk.DropCopyQuoteEventTypeCadenzaTaskQuoteRequestAck),
			Source:    cadenzasdk.F("source"),
			Timestamp: cadenzasdk.F(int64(1632933600000)),
			Payload: cadenzasdk.F(cadenzasdk.QuoteParam{
				QuoteRequestID:    cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				BaseCurrency:      cadenzasdk.F("BTC"),
				QuoteCurrency:     cadenzasdk.F("USDT"),
				AskPrice:          cadenzasdk.F(42859.990000),
				AskQuantity:       cadenzasdk.F(1.000000),
				BidPrice:          cadenzasdk.F(42859.710000),
				BidQuantity:       cadenzasdk.F(1.000000),
				Timestamp:         cadenzasdk.F(int64(1632933600000)),
				CreatedAt:         cadenzasdk.F(int64(1632933600000)),
				ValidUntil:        cadenzasdk.F(int64(1632933600000)),
				ExpiredAt:         cadenzasdk.F(int64(1632933600000)),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.QuoteExchangeTypeBinance),
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
