// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk

import (
	"context"
	"net/http"

	"github.com/cyberapper/cadenza-lite-sdk-go/internal/apijson"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/param"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/requestconfig"
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
)

// WebhookService contains methods and other services that help with interacting
// with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// PubSub Event Handler
func (r *WebhookService) Pubsub(ctx context.Context, body WebhookPubsubParams, opts ...option.RequestOption) (res *WebhookPubsubResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EventParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// The type of the event (e.g., order, executionReport, portfolio, orderBook).
	EventType param.Field[string] `json:"eventType,required"`
	// The actual data of the event, which varies based on the event type.
	Payload param.Field[EventPayloadUnionParam] `json:"payload,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64] `json:"timestamp,required"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The actual data of the event, which varies based on the event type.
type EventPayloadParam struct {
	// Base currency is the currency you want to buy or sell
	BaseCurrency param.Field[string] `json:"baseCurrency"`
	// Quote currency is the currency you want to pay or receive, and the price of the
	// base currency is quoted in the quote currency
	QuoteCurrency param.Field[string] `json:"quoteCurrency"`
	// Order side, BUY or SELL
	OrderSide param.Field[string] `json:"orderSide"`
	// Amount of the base currency
	Quantity param.Field[float64] `json:"quantity"`
	// Amount of the quote currency
	QuoteQuantity param.Field[float64] `json:"quoteQuantity"`
	// The identifier for the exchange account
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Quote request ID
	QuoteRequestID param.Field[string] `json:"quoteRequestId" format:"uuid"`
	// Levarage
	Leverage param.Field[int64] `json:"leverage"`
	// Order type
	OrderType param.Field[EventPayloadOrderType] `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID param.Field[string] `json:"positionId" format:"uuid"`
	// Price
	Price param.Field[float64] `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance param.Field[float64] `json:"priceSlippageTolerance"`
	// Symbol
	Symbol param.Field[string] `json:"symbol"`
	// Time in force
	TimeInForce param.Field[EventPayloadTimeInForce] `json:"timeInForce"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy param.Field[EventPayloadRoutePolicy] `json:"routePolicy"`
	Priority    param.Field[interface{}]             `json:"priority,required"`
	// Quote ID used by exchange for RFQ, e.g. WINTERMUTE need this field to execute
	// QUOTED order
	QuoteID param.Field[string] `json:"quoteId"`
	// Tenant ID
	TenantID param.Field[string] `json:"tenantId"`
	// Order ID
	OrderID param.Field[string] `json:"orderId"`
	// Ask price
	AskPrice param.Field[float64] `json:"askPrice"`
	// Ask quantity
	AskQuantity param.Field[float64] `json:"askQuantity"`
	// Bid price
	BidPrice param.Field[float64] `json:"bidPrice"`
	// Bid quantity
	BidQuantity param.Field[float64] `json:"bidQuantity"`
	// Create time of the quote
	Timestamp param.Field[int64] `json:"timestamp"`
	// Expiration time of the quote
	ValidUntil      param.Field[int64]       `json:"validUntil"`
	OrderCandidates param.Field[interface{}] `json:"orderCandidates,required"`
	// The total cost of this order.
	Cost param.Field[float64] `json:"cost"`
	// Created timestamp
	CreatedAt param.Field[int64] `json:"createdAt"`
	// Exchange type
	ExchangeType param.Field[string] `json:"exchangeType"`
	// The quantity of this order that has been filled.
	Filled param.Field[float64] `json:"filled"`
	// Order status
	Status param.Field[EventPayloadStatus] `json:"status"`
	// Last updated timestamp
	UpdatedAt param.Field[int64] `json:"updatedAt"`
	// User ID
	UserID param.Field[string] `json:"userId" format:"uuid"`
	// Fee
	Fee param.Field[float64] `json:"fee"`
	// Fee currency
	FeeCurrency param.Field[string] `json:"feeCurrency"`
	// Order request ID, Client Order ID
	ClOrdID    param.Field[string]      `json:"clOrdId" format:"uuid"`
	Order      param.Field[OrderParam]  `json:"order"`
	Fees       param.Field[interface{}] `json:"fees,required"`
	Executions param.Field[interface{}] `json:"executions,required"`
	Payload    param.Field[interface{}] `json:"payload,required"`
	Asks       param.Field[interface{}] `json:"asks,required"`
	Bids       param.Field[interface{}] `json:"bids,required"`
	// Order book level
	Level    param.Field[int64]                `json:"level"`
	Interval param.Field[EventPayloadInterval] `json:"interval"`
	Candles  param.Field[interface{}]          `json:"candles,required"`
}

func (r EventPayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventPayloadParam) implementsEventPayloadUnionParam() {}

// The actual data of the event, which varies based on the event type.
//
// Satisfied by [EventPayloadQuoteRequestParam],
// [EventPayloadPlaceOrderRequestParam], [EventPayloadCancelOrderRequestParam],
// [QuoteWithOrderCandidatesParam], [OrderParam], [ExecutionReportParam],
// [EventPayloadExchangeAccountPortfolioParam], [OrderbookParam],
// [EventPayloadKlineParam], [EventPayloadParam].
type EventPayloadUnionParam interface {
	implementsEventPayloadUnionParam()
}

type EventPayloadQuoteRequestParam struct {
	// Base currency is the currency you want to buy or sell
	BaseCurrency param.Field[string] `json:"baseCurrency,required"`
	// Order side, BUY or SELL
	OrderSide param.Field[string] `json:"orderSide,required"`
	// Quote currency is the currency you want to pay or receive, and the price of the
	// base currency is quoted in the quote currency
	QuoteCurrency param.Field[string] `json:"quoteCurrency,required"`
	// The identifier for the exchange account
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Amount of the base currency
	Quantity param.Field[float64] `json:"quantity"`
	// Amount of the quote currency
	QuoteQuantity param.Field[float64] `json:"quoteQuantity"`
}

func (r EventPayloadQuoteRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventPayloadQuoteRequestParam) implementsEventPayloadUnionParam() {}

type EventPayloadPlaceOrderRequestParam struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Levarage
	Leverage param.Field[int64] `json:"leverage"`
	// Order side
	OrderSide param.Field[EventPayloadPlaceOrderRequestOrderSide] `json:"orderSide"`
	// Order type
	OrderType param.Field[EventPayloadPlaceOrderRequestOrderType] `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID param.Field[string] `json:"positionId" format:"uuid"`
	// Price
	Price param.Field[float64] `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance param.Field[float64] `json:"priceSlippageTolerance"`
	// Priority list of exchange account ID in descending order
	Priority param.Field[[]string] `json:"priority"`
	// Quantity. One of quantity or quoteQuantity must be provided. If both is
	// provided, only quantity will be used.
	Quantity param.Field[float64] `json:"quantity"`
	// Quote ID used by exchange for RFQ, e.g. WINTERMUTE need this field to execute
	// QUOTED order
	QuoteID param.Field[string] `json:"quoteId"`
	// Quote Quantity
	QuoteQuantity param.Field[float64] `json:"quoteQuantity"`
	// Quote request ID
	QuoteRequestID param.Field[string] `json:"quoteRequestId" format:"uuid"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy param.Field[EventPayloadPlaceOrderRequestRoutePolicy] `json:"routePolicy"`
	// Symbol
	Symbol param.Field[string] `json:"symbol"`
	// Tenant ID
	TenantID param.Field[string] `json:"tenantId"`
	// Time in force
	TimeInForce param.Field[EventPayloadPlaceOrderRequestTimeInForce] `json:"timeInForce"`
}

func (r EventPayloadPlaceOrderRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventPayloadPlaceOrderRequestParam) implementsEventPayloadUnionParam() {}

// Order side
type EventPayloadPlaceOrderRequestOrderSide string

const (
	EventPayloadPlaceOrderRequestOrderSideBuy  EventPayloadPlaceOrderRequestOrderSide = "BUY"
	EventPayloadPlaceOrderRequestOrderSideSell EventPayloadPlaceOrderRequestOrderSide = "SELL"
)

func (r EventPayloadPlaceOrderRequestOrderSide) IsKnown() bool {
	switch r {
	case EventPayloadPlaceOrderRequestOrderSideBuy, EventPayloadPlaceOrderRequestOrderSideSell:
		return true
	}
	return false
}

// Order type
type EventPayloadPlaceOrderRequestOrderType string

const (
	EventPayloadPlaceOrderRequestOrderTypeMarket          EventPayloadPlaceOrderRequestOrderType = "MARKET"
	EventPayloadPlaceOrderRequestOrderTypeLimit           EventPayloadPlaceOrderRequestOrderType = "LIMIT"
	EventPayloadPlaceOrderRequestOrderTypeStopLoss        EventPayloadPlaceOrderRequestOrderType = "STOP_LOSS"
	EventPayloadPlaceOrderRequestOrderTypeStopLossLimit   EventPayloadPlaceOrderRequestOrderType = "STOP_LOSS_LIMIT"
	EventPayloadPlaceOrderRequestOrderTypeTakeProfit      EventPayloadPlaceOrderRequestOrderType = "TAKE_PROFIT"
	EventPayloadPlaceOrderRequestOrderTypeTakeProfitLimit EventPayloadPlaceOrderRequestOrderType = "TAKE_PROFIT_LIMIT"
	EventPayloadPlaceOrderRequestOrderTypeQuoted          EventPayloadPlaceOrderRequestOrderType = "QUOTED"
)

func (r EventPayloadPlaceOrderRequestOrderType) IsKnown() bool {
	switch r {
	case EventPayloadPlaceOrderRequestOrderTypeMarket, EventPayloadPlaceOrderRequestOrderTypeLimit, EventPayloadPlaceOrderRequestOrderTypeStopLoss, EventPayloadPlaceOrderRequestOrderTypeStopLossLimit, EventPayloadPlaceOrderRequestOrderTypeTakeProfit, EventPayloadPlaceOrderRequestOrderTypeTakeProfitLimit, EventPayloadPlaceOrderRequestOrderTypeQuoted:
		return true
	}
	return false
}

// Route policy. For PRIORITY, the order request will be routed to the exchange
// account with the highest priority. For QUOTE, the system will execute the
// execution plan based on the quote. Order request with route policy QUOTE will
// only accept two parameters, quoteRequestId and priceSlippageTolerance
type EventPayloadPlaceOrderRequestRoutePolicy string

const (
	EventPayloadPlaceOrderRequestRoutePolicyPriority EventPayloadPlaceOrderRequestRoutePolicy = "PRIORITY"
	EventPayloadPlaceOrderRequestRoutePolicyQuote    EventPayloadPlaceOrderRequestRoutePolicy = "QUOTE"
)

func (r EventPayloadPlaceOrderRequestRoutePolicy) IsKnown() bool {
	switch r {
	case EventPayloadPlaceOrderRequestRoutePolicyPriority, EventPayloadPlaceOrderRequestRoutePolicyQuote:
		return true
	}
	return false
}

// Time in force
type EventPayloadPlaceOrderRequestTimeInForce string

const (
	EventPayloadPlaceOrderRequestTimeInForceDay EventPayloadPlaceOrderRequestTimeInForce = "DAY"
	EventPayloadPlaceOrderRequestTimeInForceGtc EventPayloadPlaceOrderRequestTimeInForce = "GTC"
	EventPayloadPlaceOrderRequestTimeInForceGtx EventPayloadPlaceOrderRequestTimeInForce = "GTX"
	EventPayloadPlaceOrderRequestTimeInForceGtd EventPayloadPlaceOrderRequestTimeInForce = "GTD"
	EventPayloadPlaceOrderRequestTimeInForceOpg EventPayloadPlaceOrderRequestTimeInForce = "OPG"
	EventPayloadPlaceOrderRequestTimeInForceCls EventPayloadPlaceOrderRequestTimeInForce = "CLS"
	EventPayloadPlaceOrderRequestTimeInForceIoc EventPayloadPlaceOrderRequestTimeInForce = "IOC"
	EventPayloadPlaceOrderRequestTimeInForceFok EventPayloadPlaceOrderRequestTimeInForce = "FOK"
	EventPayloadPlaceOrderRequestTimeInForceGfa EventPayloadPlaceOrderRequestTimeInForce = "GFA"
	EventPayloadPlaceOrderRequestTimeInForceGfs EventPayloadPlaceOrderRequestTimeInForce = "GFS"
	EventPayloadPlaceOrderRequestTimeInForceGtm EventPayloadPlaceOrderRequestTimeInForce = "GTM"
	EventPayloadPlaceOrderRequestTimeInForceMoo EventPayloadPlaceOrderRequestTimeInForce = "MOO"
	EventPayloadPlaceOrderRequestTimeInForceMoc EventPayloadPlaceOrderRequestTimeInForce = "MOC"
	EventPayloadPlaceOrderRequestTimeInForceExt EventPayloadPlaceOrderRequestTimeInForce = "EXT"
)

func (r EventPayloadPlaceOrderRequestTimeInForce) IsKnown() bool {
	switch r {
	case EventPayloadPlaceOrderRequestTimeInForceDay, EventPayloadPlaceOrderRequestTimeInForceGtc, EventPayloadPlaceOrderRequestTimeInForceGtx, EventPayloadPlaceOrderRequestTimeInForceGtd, EventPayloadPlaceOrderRequestTimeInForceOpg, EventPayloadPlaceOrderRequestTimeInForceCls, EventPayloadPlaceOrderRequestTimeInForceIoc, EventPayloadPlaceOrderRequestTimeInForceFok, EventPayloadPlaceOrderRequestTimeInForceGfa, EventPayloadPlaceOrderRequestTimeInForceGfs, EventPayloadPlaceOrderRequestTimeInForceGtm, EventPayloadPlaceOrderRequestTimeInForceMoo, EventPayloadPlaceOrderRequestTimeInForceMoc, EventPayloadPlaceOrderRequestTimeInForceExt:
		return true
	}
	return false
}

type EventPayloadCancelOrderRequestParam struct {
	// Order ID
	OrderID param.Field[string] `json:"orderId,required"`
}

func (r EventPayloadCancelOrderRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventPayloadCancelOrderRequestParam) implementsEventPayloadUnionParam() {}

type EventPayloadExchangeAccountPortfolioParam struct {
	Payload param.Field[EventPayloadExchangeAccountPortfolioPayloadParam] `json:"payload"`
}

func (r EventPayloadExchangeAccountPortfolioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventPayloadExchangeAccountPortfolioParam) implementsEventPayloadUnionParam() {}

type EventPayloadExchangeAccountPortfolioPayloadParam struct {
	Balances param.Field[[]EventPayloadExchangeAccountPortfolioPayloadBalanceParam] `json:"balances,required"`
	// Exchange Account Credit Info
	Credit param.Field[ExchangeAccountCreditParam] `json:"credit,required"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId,required" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[EventPayloadExchangeAccountPortfolioPayloadExchangeType]    `json:"exchangeType,required"`
	Positions    param.Field[[]EventPayloadExchangeAccountPortfolioPayloadPositionParam] `json:"positions,required"`
	// The timestamp when the portfolio information was updated.
	UpdatedAt param.Field[int64] `json:"updatedAt,required"`
}

func (r EventPayloadExchangeAccountPortfolioPayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventPayloadExchangeAccountPortfolioPayloadBalanceParam struct {
	// Asset
	Asset param.Field[string] `json:"asset,required"`
	// Free balance
	Free param.Field[float64] `json:"free,required"`
	// Locked balance
	Locked param.Field[float64] `json:"locked,required"`
	// Total balance
	Total param.Field[float64] `json:"total,required"`
}

func (r EventPayloadExchangeAccountPortfolioPayloadBalanceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Exchange type
type EventPayloadExchangeAccountPortfolioPayloadExchangeType string

const (
	EventPayloadExchangeAccountPortfolioPayloadExchangeTypeBinance       EventPayloadExchangeAccountPortfolioPayloadExchangeType = "BINANCE"
	EventPayloadExchangeAccountPortfolioPayloadExchangeTypeBinanceMargin EventPayloadExchangeAccountPortfolioPayloadExchangeType = "BINANCE_MARGIN"
	EventPayloadExchangeAccountPortfolioPayloadExchangeTypeB2C2          EventPayloadExchangeAccountPortfolioPayloadExchangeType = "B2C2"
	EventPayloadExchangeAccountPortfolioPayloadExchangeTypeWintermute    EventPayloadExchangeAccountPortfolioPayloadExchangeType = "WINTERMUTE"
	EventPayloadExchangeAccountPortfolioPayloadExchangeTypeBlockfills    EventPayloadExchangeAccountPortfolioPayloadExchangeType = "BLOCKFILLS"
	EventPayloadExchangeAccountPortfolioPayloadExchangeTypeStonex        EventPayloadExchangeAccountPortfolioPayloadExchangeType = "STONEX"
)

func (r EventPayloadExchangeAccountPortfolioPayloadExchangeType) IsKnown() bool {
	switch r {
	case EventPayloadExchangeAccountPortfolioPayloadExchangeTypeBinance, EventPayloadExchangeAccountPortfolioPayloadExchangeTypeBinanceMargin, EventPayloadExchangeAccountPortfolioPayloadExchangeTypeB2C2, EventPayloadExchangeAccountPortfolioPayloadExchangeTypeWintermute, EventPayloadExchangeAccountPortfolioPayloadExchangeTypeBlockfills, EventPayloadExchangeAccountPortfolioPayloadExchangeTypeStonex:
		return true
	}
	return false
}

type EventPayloadExchangeAccountPortfolioPayloadPositionParam struct {
	// Amount
	Amount param.Field[float64] `json:"amount,required"`
	// Position side
	PositionSide param.Field[EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSide] `json:"positionSide,required"`
	// Status
	Status param.Field[EventPayloadExchangeAccountPortfolioPayloadPositionsStatus] `json:"status,required"`
	// Symbol
	Symbol param.Field[string] `json:"symbol,required"`
	// Cost
	Cost param.Field[float64] `json:"cost"`
	// Entry price
	EntryPrice param.Field[float64] `json:"entryPrice"`
}

func (r EventPayloadExchangeAccountPortfolioPayloadPositionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Position side
type EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSide string

const (
	EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSideLong  EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSide = "LONG"
	EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSideShort EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSide = "SHORT"
)

func (r EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSide) IsKnown() bool {
	switch r {
	case EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSideLong, EventPayloadExchangeAccountPortfolioPayloadPositionsPositionSideShort:
		return true
	}
	return false
}

// Status
type EventPayloadExchangeAccountPortfolioPayloadPositionsStatus string

const (
	EventPayloadExchangeAccountPortfolioPayloadPositionsStatusOpen EventPayloadExchangeAccountPortfolioPayloadPositionsStatus = "OPEN"
)

func (r EventPayloadExchangeAccountPortfolioPayloadPositionsStatus) IsKnown() bool {
	switch r {
	case EventPayloadExchangeAccountPortfolioPayloadPositionsStatusOpen:
		return true
	}
	return false
}

type EventPayloadKlineParam struct {
	Candles param.Field[[]OhlcvParam] `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[EventPayloadKlineExchangeType] `json:"exchangeType"`
	Interval     param.Field[EventPayloadKlineInterval]     `json:"interval"`
	Symbol       param.Field[string]                        `json:"symbol"`
}

func (r EventPayloadKlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventPayloadKlineParam) implementsEventPayloadUnionParam() {}

// Exchange type
type EventPayloadKlineExchangeType string

const (
	EventPayloadKlineExchangeTypeBinance       EventPayloadKlineExchangeType = "BINANCE"
	EventPayloadKlineExchangeTypeBinanceMargin EventPayloadKlineExchangeType = "BINANCE_MARGIN"
	EventPayloadKlineExchangeTypeB2C2          EventPayloadKlineExchangeType = "B2C2"
	EventPayloadKlineExchangeTypeWintermute    EventPayloadKlineExchangeType = "WINTERMUTE"
	EventPayloadKlineExchangeTypeBlockfills    EventPayloadKlineExchangeType = "BLOCKFILLS"
	EventPayloadKlineExchangeTypeStonex        EventPayloadKlineExchangeType = "STONEX"
)

func (r EventPayloadKlineExchangeType) IsKnown() bool {
	switch r {
	case EventPayloadKlineExchangeTypeBinance, EventPayloadKlineExchangeTypeBinanceMargin, EventPayloadKlineExchangeTypeB2C2, EventPayloadKlineExchangeTypeWintermute, EventPayloadKlineExchangeTypeBlockfills, EventPayloadKlineExchangeTypeStonex:
		return true
	}
	return false
}

type EventPayloadKlineInterval string

const (
	EventPayloadKlineInterval1s  EventPayloadKlineInterval = "1s"
	EventPayloadKlineInterval1m  EventPayloadKlineInterval = "1m"
	EventPayloadKlineInterval5m  EventPayloadKlineInterval = "5m"
	EventPayloadKlineInterval15m EventPayloadKlineInterval = "15m"
	EventPayloadKlineInterval30m EventPayloadKlineInterval = "30m"
	EventPayloadKlineInterval1h  EventPayloadKlineInterval = "1h"
	EventPayloadKlineInterval2h  EventPayloadKlineInterval = "2h"
	EventPayloadKlineInterval1d  EventPayloadKlineInterval = "1d"
	EventPayloadKlineInterval1w  EventPayloadKlineInterval = "1w"
)

func (r EventPayloadKlineInterval) IsKnown() bool {
	switch r {
	case EventPayloadKlineInterval1s, EventPayloadKlineInterval1m, EventPayloadKlineInterval5m, EventPayloadKlineInterval15m, EventPayloadKlineInterval30m, EventPayloadKlineInterval1h, EventPayloadKlineInterval2h, EventPayloadKlineInterval1d, EventPayloadKlineInterval1w:
		return true
	}
	return false
}

// Order type
type EventPayloadOrderType string

const (
	EventPayloadOrderTypeMarket          EventPayloadOrderType = "MARKET"
	EventPayloadOrderTypeLimit           EventPayloadOrderType = "LIMIT"
	EventPayloadOrderTypeStopLoss        EventPayloadOrderType = "STOP_LOSS"
	EventPayloadOrderTypeStopLossLimit   EventPayloadOrderType = "STOP_LOSS_LIMIT"
	EventPayloadOrderTypeTakeProfit      EventPayloadOrderType = "TAKE_PROFIT"
	EventPayloadOrderTypeTakeProfitLimit EventPayloadOrderType = "TAKE_PROFIT_LIMIT"
	EventPayloadOrderTypeQuoted          EventPayloadOrderType = "QUOTED"
)

func (r EventPayloadOrderType) IsKnown() bool {
	switch r {
	case EventPayloadOrderTypeMarket, EventPayloadOrderTypeLimit, EventPayloadOrderTypeStopLoss, EventPayloadOrderTypeStopLossLimit, EventPayloadOrderTypeTakeProfit, EventPayloadOrderTypeTakeProfitLimit, EventPayloadOrderTypeQuoted:
		return true
	}
	return false
}

// Time in force
type EventPayloadTimeInForce string

const (
	EventPayloadTimeInForceDay EventPayloadTimeInForce = "DAY"
	EventPayloadTimeInForceGtc EventPayloadTimeInForce = "GTC"
	EventPayloadTimeInForceGtx EventPayloadTimeInForce = "GTX"
	EventPayloadTimeInForceGtd EventPayloadTimeInForce = "GTD"
	EventPayloadTimeInForceOpg EventPayloadTimeInForce = "OPG"
	EventPayloadTimeInForceCls EventPayloadTimeInForce = "CLS"
	EventPayloadTimeInForceIoc EventPayloadTimeInForce = "IOC"
	EventPayloadTimeInForceFok EventPayloadTimeInForce = "FOK"
	EventPayloadTimeInForceGfa EventPayloadTimeInForce = "GFA"
	EventPayloadTimeInForceGfs EventPayloadTimeInForce = "GFS"
	EventPayloadTimeInForceGtm EventPayloadTimeInForce = "GTM"
	EventPayloadTimeInForceMoo EventPayloadTimeInForce = "MOO"
	EventPayloadTimeInForceMoc EventPayloadTimeInForce = "MOC"
	EventPayloadTimeInForceExt EventPayloadTimeInForce = "EXT"
)

func (r EventPayloadTimeInForce) IsKnown() bool {
	switch r {
	case EventPayloadTimeInForceDay, EventPayloadTimeInForceGtc, EventPayloadTimeInForceGtx, EventPayloadTimeInForceGtd, EventPayloadTimeInForceOpg, EventPayloadTimeInForceCls, EventPayloadTimeInForceIoc, EventPayloadTimeInForceFok, EventPayloadTimeInForceGfa, EventPayloadTimeInForceGfs, EventPayloadTimeInForceGtm, EventPayloadTimeInForceMoo, EventPayloadTimeInForceMoc, EventPayloadTimeInForceExt:
		return true
	}
	return false
}

// Route policy. For PRIORITY, the order request will be routed to the exchange
// account with the highest priority. For QUOTE, the system will execute the
// execution plan based on the quote. Order request with route policy QUOTE will
// only accept two parameters, quoteRequestId and priceSlippageTolerance
type EventPayloadRoutePolicy string

const (
	EventPayloadRoutePolicyPriority EventPayloadRoutePolicy = "PRIORITY"
	EventPayloadRoutePolicyQuote    EventPayloadRoutePolicy = "QUOTE"
)

func (r EventPayloadRoutePolicy) IsKnown() bool {
	switch r {
	case EventPayloadRoutePolicyPriority, EventPayloadRoutePolicyQuote:
		return true
	}
	return false
}

// Order status
type EventPayloadStatus string

const (
	EventPayloadStatusSubmitted       EventPayloadStatus = "SUBMITTED"
	EventPayloadStatusAccepted        EventPayloadStatus = "ACCEPTED"
	EventPayloadStatusOpen            EventPayloadStatus = "OPEN"
	EventPayloadStatusPartiallyFilled EventPayloadStatus = "PARTIALLY_FILLED"
	EventPayloadStatusFilled          EventPayloadStatus = "FILLED"
	EventPayloadStatusCanceled        EventPayloadStatus = "CANCELED"
	EventPayloadStatusPendingCancel   EventPayloadStatus = "PENDING_CANCEL"
	EventPayloadStatusRejected        EventPayloadStatus = "REJECTED"
	EventPayloadStatusExpired         EventPayloadStatus = "EXPIRED"
	EventPayloadStatusRevoked         EventPayloadStatus = "REVOKED"
)

func (r EventPayloadStatus) IsKnown() bool {
	switch r {
	case EventPayloadStatusSubmitted, EventPayloadStatusAccepted, EventPayloadStatusOpen, EventPayloadStatusPartiallyFilled, EventPayloadStatusFilled, EventPayloadStatusCanceled, EventPayloadStatusPendingCancel, EventPayloadStatusRejected, EventPayloadStatusExpired, EventPayloadStatusRevoked:
		return true
	}
	return false
}

type EventPayloadInterval string

const (
	EventPayloadInterval1s  EventPayloadInterval = "1s"
	EventPayloadInterval1m  EventPayloadInterval = "1m"
	EventPayloadInterval5m  EventPayloadInterval = "5m"
	EventPayloadInterval15m EventPayloadInterval = "15m"
	EventPayloadInterval30m EventPayloadInterval = "30m"
	EventPayloadInterval1h  EventPayloadInterval = "1h"
	EventPayloadInterval2h  EventPayloadInterval = "2h"
	EventPayloadInterval1d  EventPayloadInterval = "1d"
	EventPayloadInterval1w  EventPayloadInterval = "1w"
)

func (r EventPayloadInterval) IsKnown() bool {
	switch r {
	case EventPayloadInterval1s, EventPayloadInterval1m, EventPayloadInterval5m, EventPayloadInterval15m, EventPayloadInterval30m, EventPayloadInterval1h, EventPayloadInterval2h, EventPayloadInterval1d, EventPayloadInterval1w:
		return true
	}
	return false
}

type WebhookPubsubResponse = interface{}

type WebhookPubsubParams struct {
	Event EventParam `json:"event,required"`
}

func (r WebhookPubsubParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Event)
}
