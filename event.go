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

// EventService contains methods and other services that help with interacting with
// the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options    []option.RequestOption
	Task       *EventTaskService
	DropCopy   *EventDropCopyService
	MarketData *EventMarketDataService
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	r.Task = NewEventTaskService(opts...)
	r.DropCopy = NewEventDropCopyService(opts...)
	r.MarketData = NewEventMarketDataService(opts...)
	return
}

// PubSub event handler placeholder
func (r *EventService) New(ctx context.Context, body EventNewParams, opts ...option.RequestOption) (res *Event, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/event"
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
	EventEventTypeCadenzaTaskQuoteRequestAck       EventEventType = "cadenza.task.quoteRequestAck"
	EventEventTypeCadenzaTaskPlaceOrderRequestAck  EventEventType = "cadenza.task.placeOrderRequestAck"
	EventEventTypeCadenzaTaskCancelOrderRequestAck EventEventType = "cadenza.task.cancelOrderRequestAck"
	EventEventTypeCadenzaDropCopyQuote             EventEventType = "cadenza.dropCopy.quote"
	EventEventTypeCadenzaDropCopyOrder             EventEventType = "cadenza.dropCopy.order"
	EventEventTypeCadenzaDropCopyExecutionReport   EventEventType = "cadenza.dropCopy.executionReport"
	EventEventTypeCadenzaDropCopyPortfolio         EventEventType = "cadenza.dropCopy.portfolio"
	EventEventTypeCadenzaMarketDataOrderBook       EventEventType = "cadenza.marketData.orderBook"
	EventEventTypeCadenzaMarketDataKline           EventEventType = "cadenza.marketData.kline"
)

func (r EventEventType) IsKnown() bool {
	switch r {
	case EventEventTypeCadenzaTaskQuoteRequestAck, EventEventTypeCadenzaTaskPlaceOrderRequestAck, EventEventTypeCadenzaTaskCancelOrderRequestAck, EventEventTypeCadenzaDropCopyQuote, EventEventTypeCadenzaDropCopyOrder, EventEventTypeCadenzaDropCopyExecutionReport, EventEventTypeCadenzaDropCopyPortfolio, EventEventTypeCadenzaMarketDataOrderBook, EventEventTypeCadenzaMarketDataKline:
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
	// deprecated, alias of createdAt, Create time of the quote
	Timestamp int64 `json:"timestamp"`
	// Create time of the quote
	CreatedAt int64 `json:"createdAt"`
	// deprecated, alias of expiredAtExpiration time of the quote
	ValidUntil int64 `json:"validUntil"`
	// Expiration time of the quote
	ExpiredAt int64 `json:"expiredAt"`
	// Exchange type
	ExchangeType string `json:"exchangeType"`
	// The total cost of this order.
	Cost float64 `json:"cost"`
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
	// This field can have the runtime type of [[]BalanceEntry].
	Balances interface{} `json:"balances,required"`
	// This field can have the runtime type of [[]PositionEntry].
	Positions interface{} `json:"positions,required"`
	// Exchange Account Credit Info
	Credit ExchangeAccountCredit `json:"credit"`
	// This field can have the runtime type of [[][]float64].
	Asks interface{} `json:"asks,required"`
	// This field can have the runtime type of [[][]float64].
	Bids interface{} `json:"bids,required"`
	// Order book level
	Level    int64                `json:"level"`
	Interval EventPayloadInterval `json:"interval"`
	Candles  Candles              `json:"candles"`
	JSON     eventPayloadJSON     `json:"-"`
	union    EventPayloadUnion
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
	CreatedAt              apijson.Field
	ValidUntil             apijson.Field
	ExpiredAt              apijson.Field
	ExchangeType           apijson.Field
	Cost                   apijson.Field
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
	Balances               apijson.Field
	Positions              apijson.Field
	Credit                 apijson.Field
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
// [ExchangeAccountPortfolio], [Orderbook], [Kline].
func (r EventPayload) AsUnion() EventPayloadUnion {
	return r.union
}

// The actual data of the event, which varies based on the event type.
//
// Union satisfied by [QuoteRequest], [PlaceOrderRequest], [CancelOrderRequest],
// [Quote], [Order], [ExecutionReport], [ExchangeAccountPortfolio], [Orderbook] or
// [Kline].
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
			Type:       reflect.TypeOf(Kline{}),
		},
	)
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
	// deprecated, alias of createdAt, Create time of the quote
	Timestamp param.Field[int64] `json:"timestamp"`
	// Create time of the quote
	CreatedAt param.Field[int64] `json:"createdAt"`
	// deprecated, alias of expiredAtExpiration time of the quote
	ValidUntil param.Field[int64] `json:"validUntil"`
	// Expiration time of the quote
	ExpiredAt param.Field[int64] `json:"expiredAt"`
	// Exchange type
	ExchangeType param.Field[string] `json:"exchangeType"`
	// The total cost of this order.
	Cost param.Field[float64] `json:"cost"`
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
	Balances   param.Field[interface{}] `json:"balances,required"`
	Positions  param.Field[interface{}] `json:"positions,required"`
	// Exchange Account Credit Info
	Credit param.Field[ExchangeAccountCreditParam] `json:"credit"`
	Asks   param.Field[interface{}]                `json:"asks,required"`
	Bids   param.Field[interface{}]                `json:"bids,required"`
	// Order book level
	Level    param.Field[int64]                `json:"level"`
	Interval param.Field[EventPayloadInterval] `json:"interval"`
	Candles  param.Field[CandlesParam]         `json:"candles"`
}

func (r EventPayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventPayloadParam) implementsEventPayloadUnionParam() {}

// The actual data of the event, which varies based on the event type.
//
// Satisfied by [QuoteRequestParam], [PlaceOrderRequestParam],
// [CancelOrderRequestParam], [QuoteParam], [OrderParam], [ExecutionReportParam],
// [ExchangeAccountPortfolioParam], [OrderbookParam], [KlineParam],
// [EventPayloadParam].
type EventPayloadUnionParam interface {
	implementsEventPayloadUnionParam()
}

type EventNewParams struct {
	Event EventParam `json:"event,required"`
}

func (r EventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Event)
}
