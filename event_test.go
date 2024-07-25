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

func TestEventDropCopyExecutionReportWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.DropCopyExecutionReport(context.TODO(), cadenzasdk.EventDropCopyExecutionReportParams{
		DropCopyExecutionReport: cadenzasdk.DropCopyExecutionReportParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.ExecutionReportParam{
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
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.OrderParam{
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
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.ExchangeAccountPortfolioParam{
				Payload: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadParam{
					ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					ExchangeType:      cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadExchangeTypeBinance),
					Balances: cadenzasdk.F([]cadenzasdk.ExchangeAccountPortfolioPayloadBalanceParam{{
						Asset:  cadenzasdk.F("BTC"),
						Free:   cadenzasdk.F(1.000000),
						Locked: cadenzasdk.F(0.000000),
						Total:  cadenzasdk.F(1.000000),
					}, {
						Asset:  cadenzasdk.F("BTC"),
						Free:   cadenzasdk.F(1.000000),
						Locked: cadenzasdk.F(0.000000),
						Total:  cadenzasdk.F(1.000000),
					}, {
						Asset:  cadenzasdk.F("BTC"),
						Free:   cadenzasdk.F(1.000000),
						Locked: cadenzasdk.F(0.000000),
						Total:  cadenzasdk.F(1.000000),
					}}),
					Positions: cadenzasdk.F([]cadenzasdk.ExchangeAccountPortfolioPayloadPositionParam{{
						Amount:       cadenzasdk.F(0.000000),
						Cost:         cadenzasdk.F(0.000000),
						EntryPrice:   cadenzasdk.F(0.000000),
						PositionSide: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadPositionsPositionSideLong),
						Status:       cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadPositionsStatusOpen),
						Symbol:       cadenzasdk.F("BTC/USDT"),
					}, {
						Amount:       cadenzasdk.F(0.000000),
						Cost:         cadenzasdk.F(0.000000),
						EntryPrice:   cadenzasdk.F(0.000000),
						PositionSide: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadPositionsPositionSideLong),
						Status:       cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadPositionsStatusOpen),
						Symbol:       cadenzasdk.F("BTC/USDT"),
					}, {
						Amount:       cadenzasdk.F(0.000000),
						Cost:         cadenzasdk.F(0.000000),
						EntryPrice:   cadenzasdk.F(0.000000),
						PositionSide: cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadPositionsPositionSideLong),
						Status:       cadenzasdk.F(cadenzasdk.ExchangeAccountPortfolioPayloadPositionsStatusOpen),
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
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.QuoteParam{
				QuoteRequestID:    cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				BaseCurrency:      cadenzasdk.F("BTC"),
				QuoteCurrency:     cadenzasdk.F("USDT"),
				AskPrice:          cadenzasdk.F(42859.990000),
				AskQuantity:       cadenzasdk.F(1.000000),
				BidPrice:          cadenzasdk.F(42859.710000),
				BidQuantity:       cadenzasdk.F(1.000000),
				Timestamp:         cadenzasdk.F(int64(1632933600000)),
				ValidUntil:        cadenzasdk.F(int64(1632933600000)),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.QuoteExchangeTypeBinance),
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
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.MarketDataKlinePayloadParam{
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      cadenzasdk.F(cadenzasdk.MarketDataKlinePayloadExchangeTypeBinance),
				Symbol:            cadenzasdk.F("symbol"),
				Interval:          cadenzasdk.F(cadenzasdk.MarketDataKlinePayloadInterval1s),
				Candles: cadenzasdk.F([]cadenzasdk.OhlcvParam{{
					C: cadenzasdk.F(0.000000),
					H: cadenzasdk.F(0.000000),
					L: cadenzasdk.F(0.000000),
					O: cadenzasdk.F(0.000000),
					T: cadenzasdk.F(int64(0)),
					V: cadenzasdk.F(0.000000),
				}, {
					C: cadenzasdk.F(0.000000),
					H: cadenzasdk.F(0.000000),
					L: cadenzasdk.F(0.000000),
					O: cadenzasdk.F(0.000000),
					T: cadenzasdk.F(int64(0)),
					V: cadenzasdk.F(0.000000),
				}, {
					C: cadenzasdk.F(0.000000),
					H: cadenzasdk.F(0.000000),
					L: cadenzasdk.F(0.000000),
					O: cadenzasdk.F(0.000000),
					T: cadenzasdk.F(int64(0)),
					V: cadenzasdk.F(0.000000),
				}}),
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
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.OrderbookParam{
				Asks:              cadenzasdk.F([][]float64{{0.000000, 0.000000}, {0.000000, 0.000000}, {0.000000, 0.000000}}),
				Bids:              cadenzasdk.F([][]float64{{0.000000, 0.000000}, {0.000000, 0.000000}, {0.000000, 0.000000}}),
				ExchangeType:      cadenzasdk.F("exchangeType"),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Symbol:            cadenzasdk.F("symbol"),
				Level:             cadenzasdk.F(int64(0)),
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

func TestEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.New(context.TODO(), cadenzasdk.EventNewParams{
		GenericEvent: cadenzasdk.GenericEventParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.QuoteRequestParam{
				BaseCurrency:      cadenzasdk.F("baseCurrency"),
				QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
				OrderSide:         cadenzasdk.F("orderSide"),
				Quantity:          cadenzasdk.F(0.000000),
				QuoteQuantity:     cadenzasdk.F(0.000000),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestEventTaskCancelOrderRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.TaskCancelOrderRequestAck(context.TODO(), cadenzasdk.EventTaskCancelOrderRequestAckParams{
		TaskCancelOrderRequestAck: cadenzasdk.TaskCancelOrderRequestAckParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.CancelOrderRequestParam{
				OrderID: cadenzasdk.F("orderId"),
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

func TestEventTaskPlaceOrderRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.TaskPlaceOrderRequestAck(context.TODO(), cadenzasdk.EventTaskPlaceOrderRequestAckParams{
		TaskPlaceOrderRequestAck: cadenzasdk.TaskPlaceOrderRequestAckParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.PlaceOrderRequestParam{
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

func TestEventTaskQuoteRequestAckWithOptionalParams(t *testing.T) {
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
	_, err := client.Event.TaskQuoteRequestAck(context.TODO(), cadenzasdk.EventTaskQuoteRequestAckParams{
		TaskQuoteRequestAck: cadenzasdk.TaskQuoteRequestAckParam{
			EventID:   "eventId",
			EventType: cadenzasdk.EventEventTypeCadenzaTaskQuoteRequestAck,
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.QuoteRequestParam{
				BaseCurrency:      cadenzasdk.F("baseCurrency"),
				QuoteCurrency:     cadenzasdk.F("quoteCurrency"),
				OrderSide:         cadenzasdk.F("orderSide"),
				Quantity:          cadenzasdk.F(0.000000),
				QuoteQuantity:     cadenzasdk.F(0.000000),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
