// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cyberapper/cadenza-lite-sdk-go/internal/apijson"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/param"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/requestconfig"
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
	"github.com/tidwall/gjson"
)

// WebhookService contains methods and other services that help with interacting
// with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options        []option.RequestOption
	CloudScheduler *WebhookCloudSchedulerService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	r.CloudScheduler = NewWebhookCloudSchedulerService(opts...)
	return
}

// PubSub Event Handler
func (r *WebhookService) Pubsub(ctx context.Context, body WebhookPubsubParams, opts ...option.RequestOption) (res *WebhookPubsubResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Event struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType EventEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64 `json:"timestamp,required"`
	// The actual data of the event, which varies based on the event type.
	Payload EventPayload `json:"payload"`
	// The source system or module that generated the event.
	Source string    `json:"source"`
	JSON   eventJSON `json:"-"`
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventJSON) RawJSON() string {
	return r.raw
}

// Event Type
type EventEventType string

const (
	EventEventTypeCadenzaTaskPlaceOrderRequestAck  EventEventType = "cadenza.task.placeOrderRequestAck"
	EventEventTypeCadenzaTaskCancelOrderRequestAck EventEventType = "cadenza.task.cancelOrderRequestAck"
	EventEventTypeCadenzaDropCopyQuote             EventEventType = "cadenza.dropCopy.quote"
	EventEventTypeCadenzaDropCopyOrder             EventEventType = "cadenza.dropCopy.order"
	EventEventTypeCadenzaDropCopyPortfolio         EventEventType = "cadenza.dropCopy.portfolio"
	EventEventTypeCadenzaMarketDataOrderBook       EventEventType = "cadenza.marketData.orderBook"
	EventEventTypeCadenzaMarketDataKline           EventEventType = "cadenza.marketData.kline"
)

func (r EventEventType) IsKnown() bool {
	switch r {
	case EventEventTypeCadenzaTaskPlaceOrderRequestAck, EventEventTypeCadenzaTaskCancelOrderRequestAck, EventEventTypeCadenzaDropCopyQuote, EventEventTypeCadenzaDropCopyOrder, EventEventTypeCadenzaDropCopyPortfolio, EventEventTypeCadenzaMarketDataOrderBook, EventEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

// The actual data of the event, which varies based on the event type.
type EventPayload struct {
	// Base currency is the currency you want to buy or sell
	BaseCurrency string `json:"baseCurrency"`
	// Quote currency is the currency you want to pay or receive, and the price of the
	// base currency is quoted in the quote currency
	QuoteCurrency string `json:"quoteCurrency"`
	// Order side, BUY or SELL
	OrderSide string `json:"orderSide"`
	// Amount of the base currency
	Quantity float64 `json:"quantity"`
	// Amount of the quote currency
	QuoteQuantity float64 `json:"quoteQuantity"`
	// The identifier for the exchange account
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Quote request ID
	QuoteRequestID string `json:"quoteRequestId" format:"uuid"`
	// Levarage
	Leverage int64 `json:"leverage"`
	// Order type
	OrderType EventPayloadOrderType `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID string `json:"positionId" format:"uuid"`
	// Price
	Price float64 `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance float64 `json:"priceSlippageTolerance"`
	// Symbol
	Symbol string `json:"symbol"`
	// Time in force
	TimeInForce EventPayloadTimeInForce `json:"timeInForce"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy EventPayloadRoutePolicy `json:"routePolicy"`
	// This field can have the runtime type of [[]string].
	Priority interface{} `json:"priority,required"`
	// Quote ID used by exchange for RFQ, e.g. WINTERMUTE need this field to execute
	// QUOTED order
	QuoteID string `json:"quoteId"`
	// Tenant ID
	TenantID string `json:"tenantId"`
	// Order ID
	OrderID string `json:"orderId"`
	// Ask price
	AskPrice float64 `json:"askPrice"`
	// Ask quantity
	AskQuantity float64 `json:"askQuantity"`
	// Bid price
	BidPrice float64 `json:"bidPrice"`
	// Bid quantity
	BidQuantity float64 `json:"bidQuantity"`
	// Create time of the quote
	Timestamp int64 `json:"timestamp"`
	// Expiration time of the quote
	ValidUntil int64 `json:"validUntil"`
	// Exchange type
	ExchangeType string `json:"exchangeType"`
	// The total cost of this order.
	Cost float64 `json:"cost"`
	// Created timestamp
	CreatedAt int64 `json:"createdAt"`
	// The quantity of this order that has been filled.
	Filled float64 `json:"filled"`
	// Order status
	Status EventPayloadStatus `json:"status"`
	// Last updated timestamp
	UpdatedAt int64 `json:"updatedAt"`
	// User ID
	UserID string `json:"userId" format:"uuid"`
	// Fee
	Fee float64 `json:"fee"`
	// Fee currency
	FeeCurrency string `json:"feeCurrency"`
	// Order request ID, Client Order ID
	ClOrdID string `json:"clOrdId" format:"uuid"`
	Order   Order  `json:"order"`
	// This field can have the runtime type of [[]ExecutionReportFee].
	Fees interface{} `json:"fees,required"`
	// This field can have the runtime type of [[]Order].
	Executions interface{} `json:"executions,required"`
	// This field can have the runtime type of [ExchangeAccountPortfolioPayload].
	Payload interface{} `json:"payload,required"`
	// This field can have the runtime type of [[][]float64].
	Asks interface{} `json:"asks,required"`
	// This field can have the runtime type of [[][]float64].
	Bids interface{} `json:"bids,required"`
	// Order book level
	Level    int64                `json:"level"`
	Interval EventPayloadInterval `json:"interval"`
	// This field can have the runtime type of [[]Ohlcv].
	Candles interface{}      `json:"candles,required"`
	JSON    eventPayloadJSON `json:"-"`
	union   EventPayloadUnion
}

// eventPayloadJSON contains the JSON metadata for the struct [EventPayload]
type eventPayloadJSON struct {
	BaseCurrency           apijson.Field
	QuoteCurrency          apijson.Field
	OrderSide              apijson.Field
	Quantity               apijson.Field
	QuoteQuantity          apijson.Field
	ExchangeAccountID      apijson.Field
	QuoteRequestID         apijson.Field
	Leverage               apijson.Field
	OrderType              apijson.Field
	PositionID             apijson.Field
	Price                  apijson.Field
	PriceSlippageTolerance apijson.Field
	Symbol                 apijson.Field
	TimeInForce            apijson.Field
	RoutePolicy            apijson.Field
	Priority               apijson.Field
	QuoteID                apijson.Field
	TenantID               apijson.Field
	OrderID                apijson.Field
	AskPrice               apijson.Field
	AskQuantity            apijson.Field
	BidPrice               apijson.Field
	BidQuantity            apijson.Field
	Timestamp              apijson.Field
	ValidUntil             apijson.Field
	ExchangeType           apijson.Field
	Cost                   apijson.Field
	CreatedAt              apijson.Field
	Filled                 apijson.Field
	Status                 apijson.Field
	UpdatedAt              apijson.Field
	UserID                 apijson.Field
	Fee                    apijson.Field
	FeeCurrency            apijson.Field
	ClOrdID                apijson.Field
	Order                  apijson.Field
	Fees                   apijson.Field
	Executions             apijson.Field
	Payload                apijson.Field
	Asks                   apijson.Field
	Bids                   apijson.Field
	Level                  apijson.Field
	Interval               apijson.Field
	Candles                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r eventPayloadJSON) RawJSON() string {
	return r.raw
}

func (r *EventPayload) UnmarshalJSON(data []byte) (err error) {
	*r = EventPayload{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EventPayloadUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [QuoteRequest], [PlaceOrderRequest],
// [CancelOrderRequest], [Quote], [Order], [ExecutionReport],
// [ExchangeAccountPortfolio], [Orderbook], [EventPayloadKline].
func (r EventPayload) AsUnion() EventPayloadUnion {
	return r.union
}

// The actual data of the event, which varies based on the event type.
//
// Union satisfied by [QuoteRequest], [PlaceOrderRequest], [CancelOrderRequest],
// [Quote], [Order], [ExecutionReport], [ExchangeAccountPortfolio], [Orderbook] or
// [EventPayloadKline].
type EventPayloadUnion interface {
	implementsEventPayload()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EventPayloadUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(QuoteRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PlaceOrderRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CancelOrderRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Quote{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Order{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExecutionReport{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExchangeAccountPortfolio{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Orderbook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventPayloadKline{}),
		},
	)
}

type EventPayloadKline struct {
	Candles []Ohlcv `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType EventPayloadKlineExchangeType `json:"exchangeType"`
	Interval     EventPayloadKlineInterval     `json:"interval"`
	Symbol       string                        `json:"symbol"`
	JSON         eventPayloadKlineJSON         `json:"-"`
}

// eventPayloadKlineJSON contains the JSON metadata for the struct
// [EventPayloadKline]
type eventPayloadKlineJSON struct {
	Candles           apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Interval          apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EventPayloadKline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventPayloadKlineJSON) RawJSON() string {
	return r.raw
}

func (r EventPayloadKline) implementsEventPayload() {}

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

type EventParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[EventEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64] `json:"timestamp,required"`
	// The actual data of the event, which varies based on the event type.
	Payload param.Field[EventPayloadUnionParam] `json:"payload"`
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
	ValidUntil param.Field[int64] `json:"validUntil"`
	// Exchange type
	ExchangeType param.Field[string] `json:"exchangeType"`
	// The total cost of this order.
	Cost param.Field[float64] `json:"cost"`
	// Created timestamp
	CreatedAt param.Field[int64] `json:"createdAt"`
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
// Satisfied by [QuoteRequestParam], [PlaceOrderRequestParam],
// [CancelOrderRequestParam], [QuoteParam], [OrderParam], [ExecutionReportParam],
// [ExchangeAccountPortfolioParam], [OrderbookParam], [EventPayloadKlineParam],
// [EventPayloadParam].
type EventPayloadUnionParam interface {
	implementsEventPayloadUnionParam()
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

type WebhookPubsubResponse struct {
	Data string                    `json:"data"`
	JSON webhookPubsubResponseJSON `json:"-"`
}

// webhookPubsubResponseJSON contains the JSON metadata for the struct
// [WebhookPubsubResponse]
type webhookPubsubResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookPubsubResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookPubsubResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookPubsubParams struct {
	Message param.Field[WebhookPubsubParamsMessage] `json:"message,required"`
	// The subscription name.
	Subscription param.Field[string] `json:"subscription,required"`
}

func (r WebhookPubsubParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookPubsubParamsMessage struct {
	// The unique identifier for the message.
	ID param.Field[string] `json:"id,required"`
	// The base64-encoded data.
	Data param.Field[string] `json:"data" format:"byte"`
}

func (r WebhookPubsubParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
