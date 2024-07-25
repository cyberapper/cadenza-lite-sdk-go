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

// EventService contains methods and other services that help with interacting with
// the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// PubSub event handler placeholder for order event
func (r *EventService) DropCopyOrder(ctx context.Context, body EventDropCopyOrderParams, opts ...option.RequestOption) (res *EventDropCopyOrderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for portfolio event
func (r *EventService) DropCopyPortfolio(ctx context.Context, body EventDropCopyPortfolioParams, opts ...option.RequestOption) (res *EventDropCopyPortfolioResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/portfolio"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for quote event
func (r *EventService) DropCopyQuote(ctx context.Context, body EventDropCopyQuoteParams, opts ...option.RequestOption) (res *EventDropCopyQuoteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for kline event
func (r *EventService) MarketDataKline(ctx context.Context, body EventMarketDataKlineParams, opts ...option.RequestOption) (res *EventMarketDataKlineResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/kline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for order book event
func (r *EventService) MarketDataOrderBook(ctx context.Context, body EventMarketDataOrderBookParams, opts ...option.RequestOption) (res *EventMarketDataOrderBookResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/orderBook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EventDropCopyOrderResponse struct {
	EventType EventDropCopyOrderResponseEventType `json:"eventType"`
	Payload   Order                               `json:"payload"`
	JSON      eventDropCopyOrderResponseJSON      `json:"-"`
	Event
}

// eventDropCopyOrderResponseJSON contains the JSON metadata for the struct
// [EventDropCopyOrderResponse]
type eventDropCopyOrderResponseJSON struct {
	EventType   apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventDropCopyOrderResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventDropCopyOrderResponseJSON) RawJSON() string {
	return r.raw
}

type EventDropCopyOrderResponseEventType string

const (
	EventDropCopyOrderResponseEventTypeCadenzaDropCopyOrder EventDropCopyOrderResponseEventType = "cadenza.dropCopy.order"
)

func (r EventDropCopyOrderResponseEventType) IsKnown() bool {
	switch r {
	case EventDropCopyOrderResponseEventTypeCadenzaDropCopyOrder:
		return true
	}
	return false
}

type EventDropCopyPortfolioResponse struct {
	EventType EventDropCopyPortfolioResponseEventType `json:"eventType"`
	Payload   ExchangeAccountPortfolio                `json:"payload"`
	JSON      eventDropCopyPortfolioResponseJSON      `json:"-"`
	Event
}

// eventDropCopyPortfolioResponseJSON contains the JSON metadata for the struct
// [EventDropCopyPortfolioResponse]
type eventDropCopyPortfolioResponseJSON struct {
	EventType   apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventDropCopyPortfolioResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventDropCopyPortfolioResponseJSON) RawJSON() string {
	return r.raw
}

type EventDropCopyPortfolioResponseEventType string

const (
	EventDropCopyPortfolioResponseEventTypeCadenzaDropCopyPortfolio EventDropCopyPortfolioResponseEventType = "cadenza.dropCopy.portfolio"
)

func (r EventDropCopyPortfolioResponseEventType) IsKnown() bool {
	switch r {
	case EventDropCopyPortfolioResponseEventTypeCadenzaDropCopyPortfolio:
		return true
	}
	return false
}

type EventDropCopyQuoteResponse struct {
	EventType EventDropCopyQuoteResponseEventType `json:"eventType"`
	Payload   Quote                               `json:"payload"`
	JSON      eventDropCopyQuoteResponseJSON      `json:"-"`
	Event
}

// eventDropCopyQuoteResponseJSON contains the JSON metadata for the struct
// [EventDropCopyQuoteResponse]
type eventDropCopyQuoteResponseJSON struct {
	EventType   apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventDropCopyQuoteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventDropCopyQuoteResponseJSON) RawJSON() string {
	return r.raw
}

type EventDropCopyQuoteResponseEventType string

const (
	EventDropCopyQuoteResponseEventTypeCadenzaDropCopyQuote EventDropCopyQuoteResponseEventType = "cadenza.dropCopy.quote"
)

func (r EventDropCopyQuoteResponseEventType) IsKnown() bool {
	switch r {
	case EventDropCopyQuoteResponseEventTypeCadenzaDropCopyQuote:
		return true
	}
	return false
}

type EventMarketDataKlineResponse struct {
	EventType EventMarketDataKlineResponseEventType `json:"eventType"`
	Payload   EventMarketDataKlineResponsePayload   `json:"payload"`
	JSON      eventMarketDataKlineResponseJSON      `json:"-"`
	Event
}

// eventMarketDataKlineResponseJSON contains the JSON metadata for the struct
// [EventMarketDataKlineResponse]
type eventMarketDataKlineResponseJSON struct {
	EventType   apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventMarketDataKlineResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventMarketDataKlineResponseJSON) RawJSON() string {
	return r.raw
}

type EventMarketDataKlineResponseEventType string

const (
	EventMarketDataKlineResponseEventTypeCadenzaMarketDataKline EventMarketDataKlineResponseEventType = "cadenza.marketData.kline"
)

func (r EventMarketDataKlineResponseEventType) IsKnown() bool {
	switch r {
	case EventMarketDataKlineResponseEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type EventMarketDataKlineResponsePayload struct {
	Candles []Ohlcv `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType EventMarketDataKlineResponsePayloadExchangeType `json:"exchangeType"`
	Interval     EventMarketDataKlineResponsePayloadInterval     `json:"interval"`
	Symbol       string                                          `json:"symbol"`
	JSON         eventMarketDataKlineResponsePayloadJSON         `json:"-"`
}

// eventMarketDataKlineResponsePayloadJSON contains the JSON metadata for the
// struct [EventMarketDataKlineResponsePayload]
type eventMarketDataKlineResponsePayloadJSON struct {
	Candles           apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Interval          apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EventMarketDataKlineResponsePayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventMarketDataKlineResponsePayloadJSON) RawJSON() string {
	return r.raw
}

// Exchange type
type EventMarketDataKlineResponsePayloadExchangeType string

const (
	EventMarketDataKlineResponsePayloadExchangeTypeBinance       EventMarketDataKlineResponsePayloadExchangeType = "BINANCE"
	EventMarketDataKlineResponsePayloadExchangeTypeBinanceMargin EventMarketDataKlineResponsePayloadExchangeType = "BINANCE_MARGIN"
	EventMarketDataKlineResponsePayloadExchangeTypeB2C2          EventMarketDataKlineResponsePayloadExchangeType = "B2C2"
	EventMarketDataKlineResponsePayloadExchangeTypeWintermute    EventMarketDataKlineResponsePayloadExchangeType = "WINTERMUTE"
	EventMarketDataKlineResponsePayloadExchangeTypeBlockfills    EventMarketDataKlineResponsePayloadExchangeType = "BLOCKFILLS"
	EventMarketDataKlineResponsePayloadExchangeTypeStonex        EventMarketDataKlineResponsePayloadExchangeType = "STONEX"
)

func (r EventMarketDataKlineResponsePayloadExchangeType) IsKnown() bool {
	switch r {
	case EventMarketDataKlineResponsePayloadExchangeTypeBinance, EventMarketDataKlineResponsePayloadExchangeTypeBinanceMargin, EventMarketDataKlineResponsePayloadExchangeTypeB2C2, EventMarketDataKlineResponsePayloadExchangeTypeWintermute, EventMarketDataKlineResponsePayloadExchangeTypeBlockfills, EventMarketDataKlineResponsePayloadExchangeTypeStonex:
		return true
	}
	return false
}

type EventMarketDataKlineResponsePayloadInterval string

const (
	EventMarketDataKlineResponsePayloadInterval1s  EventMarketDataKlineResponsePayloadInterval = "1s"
	EventMarketDataKlineResponsePayloadInterval1m  EventMarketDataKlineResponsePayloadInterval = "1m"
	EventMarketDataKlineResponsePayloadInterval5m  EventMarketDataKlineResponsePayloadInterval = "5m"
	EventMarketDataKlineResponsePayloadInterval15m EventMarketDataKlineResponsePayloadInterval = "15m"
	EventMarketDataKlineResponsePayloadInterval30m EventMarketDataKlineResponsePayloadInterval = "30m"
	EventMarketDataKlineResponsePayloadInterval1h  EventMarketDataKlineResponsePayloadInterval = "1h"
	EventMarketDataKlineResponsePayloadInterval2h  EventMarketDataKlineResponsePayloadInterval = "2h"
	EventMarketDataKlineResponsePayloadInterval1d  EventMarketDataKlineResponsePayloadInterval = "1d"
	EventMarketDataKlineResponsePayloadInterval1w  EventMarketDataKlineResponsePayloadInterval = "1w"
)

func (r EventMarketDataKlineResponsePayloadInterval) IsKnown() bool {
	switch r {
	case EventMarketDataKlineResponsePayloadInterval1s, EventMarketDataKlineResponsePayloadInterval1m, EventMarketDataKlineResponsePayloadInterval5m, EventMarketDataKlineResponsePayloadInterval15m, EventMarketDataKlineResponsePayloadInterval30m, EventMarketDataKlineResponsePayloadInterval1h, EventMarketDataKlineResponsePayloadInterval2h, EventMarketDataKlineResponsePayloadInterval1d, EventMarketDataKlineResponsePayloadInterval1w:
		return true
	}
	return false
}

type EventMarketDataOrderBookResponse struct {
	EventType EventMarketDataOrderBookResponseEventType `json:"eventType"`
	Payload   Orderbook                                 `json:"payload"`
	JSON      eventMarketDataOrderBookResponseJSON      `json:"-"`
	Event
}

// eventMarketDataOrderBookResponseJSON contains the JSON metadata for the struct
// [EventMarketDataOrderBookResponse]
type eventMarketDataOrderBookResponseJSON struct {
	EventType   apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventMarketDataOrderBookResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventMarketDataOrderBookResponseJSON) RawJSON() string {
	return r.raw
}

type EventMarketDataOrderBookResponseEventType string

const (
	EventMarketDataOrderBookResponseEventTypeCadenzaMarketDataOrderBook EventMarketDataOrderBookResponseEventType = "cadenza.marketData.orderBook"
)

func (r EventMarketDataOrderBookResponseEventType) IsKnown() bool {
	switch r {
	case EventMarketDataOrderBookResponseEventTypeCadenzaMarketDataOrderBook:
		return true
	}
	return false
}

type EventDropCopyOrderParams struct {
	// A unique identifier for the event.
	EventID   param.Field[string]                            `json:"eventId,required"`
	EventType param.Field[EventDropCopyOrderParamsEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]           `json:"timestamp,required"`
	Payload   param.Field[OrderUnionParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventDropCopyOrderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventDropCopyOrderParamsEventType string

const (
	EventDropCopyOrderParamsEventTypeCadenzaDropCopyOrder             EventDropCopyOrderParamsEventType = "cadenza.dropCopy.order"
	EventDropCopyOrderParamsEventTypeCadenzaTaskPlaceOrderRequestAck  EventDropCopyOrderParamsEventType = "cadenza.task.placeOrderRequestAck"
	EventDropCopyOrderParamsEventTypeCadenzaTaskCancelOrderRequestAck EventDropCopyOrderParamsEventType = "cadenza.task.cancelOrderRequestAck"
	EventDropCopyOrderParamsEventTypeCadenzaDropCopyQuote             EventDropCopyOrderParamsEventType = "cadenza.dropCopy.quote"
	EventDropCopyOrderParamsEventTypeCadenzaDropCopyPortfolio         EventDropCopyOrderParamsEventType = "cadenza.dropCopy.portfolio"
	EventDropCopyOrderParamsEventTypeCadenzaMarketDataOrderBook       EventDropCopyOrderParamsEventType = "cadenza.marketData.orderBook"
	EventDropCopyOrderParamsEventTypeCadenzaMarketDataKline           EventDropCopyOrderParamsEventType = "cadenza.marketData.kline"
)

func (r EventDropCopyOrderParamsEventType) IsKnown() bool {
	switch r {
	case EventDropCopyOrderParamsEventTypeCadenzaDropCopyOrder, EventDropCopyOrderParamsEventTypeCadenzaTaskPlaceOrderRequestAck, EventDropCopyOrderParamsEventTypeCadenzaTaskCancelOrderRequestAck, EventDropCopyOrderParamsEventTypeCadenzaDropCopyQuote, EventDropCopyOrderParamsEventTypeCadenzaDropCopyPortfolio, EventDropCopyOrderParamsEventTypeCadenzaMarketDataOrderBook, EventDropCopyOrderParamsEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type EventDropCopyPortfolioParams struct {
	// A unique identifier for the event.
	EventID   param.Field[string]                                `json:"eventId,required"`
	EventType param.Field[EventDropCopyPortfolioParamsEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                              `json:"timestamp,required"`
	Payload   param.Field[ExchangeAccountPortfolioUnionParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventDropCopyPortfolioParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventDropCopyPortfolioParamsEventType string

const (
	EventDropCopyPortfolioParamsEventTypeCadenzaDropCopyPortfolio         EventDropCopyPortfolioParamsEventType = "cadenza.dropCopy.portfolio"
	EventDropCopyPortfolioParamsEventTypeCadenzaTaskPlaceOrderRequestAck  EventDropCopyPortfolioParamsEventType = "cadenza.task.placeOrderRequestAck"
	EventDropCopyPortfolioParamsEventTypeCadenzaTaskCancelOrderRequestAck EventDropCopyPortfolioParamsEventType = "cadenza.task.cancelOrderRequestAck"
	EventDropCopyPortfolioParamsEventTypeCadenzaDropCopyQuote             EventDropCopyPortfolioParamsEventType = "cadenza.dropCopy.quote"
	EventDropCopyPortfolioParamsEventTypeCadenzaDropCopyOrder             EventDropCopyPortfolioParamsEventType = "cadenza.dropCopy.order"
	EventDropCopyPortfolioParamsEventTypeCadenzaMarketDataOrderBook       EventDropCopyPortfolioParamsEventType = "cadenza.marketData.orderBook"
	EventDropCopyPortfolioParamsEventTypeCadenzaMarketDataKline           EventDropCopyPortfolioParamsEventType = "cadenza.marketData.kline"
)

func (r EventDropCopyPortfolioParamsEventType) IsKnown() bool {
	switch r {
	case EventDropCopyPortfolioParamsEventTypeCadenzaDropCopyPortfolio, EventDropCopyPortfolioParamsEventTypeCadenzaTaskPlaceOrderRequestAck, EventDropCopyPortfolioParamsEventTypeCadenzaTaskCancelOrderRequestAck, EventDropCopyPortfolioParamsEventTypeCadenzaDropCopyQuote, EventDropCopyPortfolioParamsEventTypeCadenzaDropCopyOrder, EventDropCopyPortfolioParamsEventTypeCadenzaMarketDataOrderBook, EventDropCopyPortfolioParamsEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type EventDropCopyQuoteParams struct {
	// A unique identifier for the event.
	EventID   param.Field[string]                            `json:"eventId,required"`
	EventType param.Field[EventDropCopyQuoteParamsEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]           `json:"timestamp,required"`
	Payload   param.Field[QuoteUnionParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventDropCopyQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventDropCopyQuoteParamsEventType string

const (
	EventDropCopyQuoteParamsEventTypeCadenzaDropCopyQuote             EventDropCopyQuoteParamsEventType = "cadenza.dropCopy.quote"
	EventDropCopyQuoteParamsEventTypeCadenzaTaskPlaceOrderRequestAck  EventDropCopyQuoteParamsEventType = "cadenza.task.placeOrderRequestAck"
	EventDropCopyQuoteParamsEventTypeCadenzaTaskCancelOrderRequestAck EventDropCopyQuoteParamsEventType = "cadenza.task.cancelOrderRequestAck"
	EventDropCopyQuoteParamsEventTypeCadenzaDropCopyOrder             EventDropCopyQuoteParamsEventType = "cadenza.dropCopy.order"
	EventDropCopyQuoteParamsEventTypeCadenzaDropCopyPortfolio         EventDropCopyQuoteParamsEventType = "cadenza.dropCopy.portfolio"
	EventDropCopyQuoteParamsEventTypeCadenzaMarketDataOrderBook       EventDropCopyQuoteParamsEventType = "cadenza.marketData.orderBook"
	EventDropCopyQuoteParamsEventTypeCadenzaMarketDataKline           EventDropCopyQuoteParamsEventType = "cadenza.marketData.kline"
)

func (r EventDropCopyQuoteParamsEventType) IsKnown() bool {
	switch r {
	case EventDropCopyQuoteParamsEventTypeCadenzaDropCopyQuote, EventDropCopyQuoteParamsEventTypeCadenzaTaskPlaceOrderRequestAck, EventDropCopyQuoteParamsEventTypeCadenzaTaskCancelOrderRequestAck, EventDropCopyQuoteParamsEventTypeCadenzaDropCopyOrder, EventDropCopyQuoteParamsEventTypeCadenzaDropCopyPortfolio, EventDropCopyQuoteParamsEventTypeCadenzaMarketDataOrderBook, EventDropCopyQuoteParamsEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type EventMarketDataKlineParams struct {
	// A unique identifier for the event.
	EventID   param.Field[string]                              `json:"eventId,required"`
	EventType param.Field[EventMarketDataKlineParamsEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                                  `json:"timestamp,required"`
	Payload   param.Field[EventMarketDataKlineParamsPayloadUnion] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventMarketDataKlineParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventMarketDataKlineParamsEventType string

const (
	EventMarketDataKlineParamsEventTypeCadenzaMarketDataKline           EventMarketDataKlineParamsEventType = "cadenza.marketData.kline"
	EventMarketDataKlineParamsEventTypeCadenzaTaskPlaceOrderRequestAck  EventMarketDataKlineParamsEventType = "cadenza.task.placeOrderRequestAck"
	EventMarketDataKlineParamsEventTypeCadenzaTaskCancelOrderRequestAck EventMarketDataKlineParamsEventType = "cadenza.task.cancelOrderRequestAck"
	EventMarketDataKlineParamsEventTypeCadenzaDropCopyQuote             EventMarketDataKlineParamsEventType = "cadenza.dropCopy.quote"
	EventMarketDataKlineParamsEventTypeCadenzaDropCopyOrder             EventMarketDataKlineParamsEventType = "cadenza.dropCopy.order"
	EventMarketDataKlineParamsEventTypeCadenzaDropCopyPortfolio         EventMarketDataKlineParamsEventType = "cadenza.dropCopy.portfolio"
	EventMarketDataKlineParamsEventTypeCadenzaMarketDataOrderBook       EventMarketDataKlineParamsEventType = "cadenza.marketData.orderBook"
)

func (r EventMarketDataKlineParamsEventType) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsEventTypeCadenzaMarketDataKline, EventMarketDataKlineParamsEventTypeCadenzaTaskPlaceOrderRequestAck, EventMarketDataKlineParamsEventTypeCadenzaTaskCancelOrderRequestAck, EventMarketDataKlineParamsEventTypeCadenzaDropCopyQuote, EventMarketDataKlineParamsEventTypeCadenzaDropCopyOrder, EventMarketDataKlineParamsEventTypeCadenzaDropCopyPortfolio, EventMarketDataKlineParamsEventTypeCadenzaMarketDataOrderBook:
		return true
	}
	return false
}

type EventMarketDataKlineParamsPayload struct {
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
	OrderType param.Field[EventMarketDataKlineParamsPayloadOrderType] `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID param.Field[string] `json:"positionId" format:"uuid"`
	// Price
	Price param.Field[float64] `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance param.Field[float64] `json:"priceSlippageTolerance"`
	// Symbol
	Symbol param.Field[string] `json:"symbol"`
	// Time in force
	TimeInForce param.Field[EventMarketDataKlineParamsPayloadTimeInForce] `json:"timeInForce"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy param.Field[EventMarketDataKlineParamsPayloadRoutePolicy] `json:"routePolicy"`
	Priority    param.Field[interface{}]                                  `json:"priority,required"`
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
	Status param.Field[EventMarketDataKlineParamsPayloadStatus] `json:"status"`
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
	Level    param.Field[int64]                                     `json:"level"`
	Interval param.Field[EventMarketDataKlineParamsPayloadInterval] `json:"interval"`
	Candles  param.Field[interface{}]                               `json:"candles,required"`
}

func (r EventMarketDataKlineParamsPayload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventMarketDataKlineParamsPayload) implementsEventMarketDataKlineParamsPayloadUnion() {}

// Satisfied by [QuoteRequestParam], [PlaceOrderRequestParam],
// [CancelOrderRequestParam], [QuoteParam], [OrderParam], [ExecutionReportParam],
// [ExchangeAccountPortfolioParam], [OrderbookParam],
// [EventMarketDataKlineParamsPayloadKline],
// [EventMarketDataKlineParamsPayloadKline], [EventMarketDataKlineParamsPayload].
type EventMarketDataKlineParamsPayloadUnion interface {
	implementsEventMarketDataKlineParamsPayloadUnion()
}

type EventMarketDataKlineParamsPayloadKline struct {
	Candles param.Field[[]OhlcvParam] `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[EventMarketDataKlineParamsPayloadKlineExchangeType] `json:"exchangeType"`
	Interval     param.Field[EventMarketDataKlineParamsPayloadKlineInterval]     `json:"interval"`
	Symbol       param.Field[string]                                             `json:"symbol"`
}

func (r EventMarketDataKlineParamsPayloadKline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EventMarketDataKlineParamsPayloadKline) implementsEventMarketDataKlineParamsPayloadUnion() {}

// Exchange type
type EventMarketDataKlineParamsPayloadKlineExchangeType string

const (
	EventMarketDataKlineParamsPayloadKlineExchangeTypeBinance       EventMarketDataKlineParamsPayloadKlineExchangeType = "BINANCE"
	EventMarketDataKlineParamsPayloadKlineExchangeTypeBinanceMargin EventMarketDataKlineParamsPayloadKlineExchangeType = "BINANCE_MARGIN"
	EventMarketDataKlineParamsPayloadKlineExchangeTypeB2C2          EventMarketDataKlineParamsPayloadKlineExchangeType = "B2C2"
	EventMarketDataKlineParamsPayloadKlineExchangeTypeWintermute    EventMarketDataKlineParamsPayloadKlineExchangeType = "WINTERMUTE"
	EventMarketDataKlineParamsPayloadKlineExchangeTypeBlockfills    EventMarketDataKlineParamsPayloadKlineExchangeType = "BLOCKFILLS"
	EventMarketDataKlineParamsPayloadKlineExchangeTypeStonex        EventMarketDataKlineParamsPayloadKlineExchangeType = "STONEX"
)

func (r EventMarketDataKlineParamsPayloadKlineExchangeType) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsPayloadKlineExchangeTypeBinance, EventMarketDataKlineParamsPayloadKlineExchangeTypeBinanceMargin, EventMarketDataKlineParamsPayloadKlineExchangeTypeB2C2, EventMarketDataKlineParamsPayloadKlineExchangeTypeWintermute, EventMarketDataKlineParamsPayloadKlineExchangeTypeBlockfills, EventMarketDataKlineParamsPayloadKlineExchangeTypeStonex:
		return true
	}
	return false
}

type EventMarketDataKlineParamsPayloadKlineInterval string

const (
	EventMarketDataKlineParamsPayloadKlineInterval1s  EventMarketDataKlineParamsPayloadKlineInterval = "1s"
	EventMarketDataKlineParamsPayloadKlineInterval1m  EventMarketDataKlineParamsPayloadKlineInterval = "1m"
	EventMarketDataKlineParamsPayloadKlineInterval5m  EventMarketDataKlineParamsPayloadKlineInterval = "5m"
	EventMarketDataKlineParamsPayloadKlineInterval15m EventMarketDataKlineParamsPayloadKlineInterval = "15m"
	EventMarketDataKlineParamsPayloadKlineInterval30m EventMarketDataKlineParamsPayloadKlineInterval = "30m"
	EventMarketDataKlineParamsPayloadKlineInterval1h  EventMarketDataKlineParamsPayloadKlineInterval = "1h"
	EventMarketDataKlineParamsPayloadKlineInterval2h  EventMarketDataKlineParamsPayloadKlineInterval = "2h"
	EventMarketDataKlineParamsPayloadKlineInterval1d  EventMarketDataKlineParamsPayloadKlineInterval = "1d"
	EventMarketDataKlineParamsPayloadKlineInterval1w  EventMarketDataKlineParamsPayloadKlineInterval = "1w"
)

func (r EventMarketDataKlineParamsPayloadKlineInterval) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsPayloadKlineInterval1s, EventMarketDataKlineParamsPayloadKlineInterval1m, EventMarketDataKlineParamsPayloadKlineInterval5m, EventMarketDataKlineParamsPayloadKlineInterval15m, EventMarketDataKlineParamsPayloadKlineInterval30m, EventMarketDataKlineParamsPayloadKlineInterval1h, EventMarketDataKlineParamsPayloadKlineInterval2h, EventMarketDataKlineParamsPayloadKlineInterval1d, EventMarketDataKlineParamsPayloadKlineInterval1w:
		return true
	}
	return false
}

// Order type
type EventMarketDataKlineParamsPayloadOrderType string

const (
	EventMarketDataKlineParamsPayloadOrderTypeMarket          EventMarketDataKlineParamsPayloadOrderType = "MARKET"
	EventMarketDataKlineParamsPayloadOrderTypeLimit           EventMarketDataKlineParamsPayloadOrderType = "LIMIT"
	EventMarketDataKlineParamsPayloadOrderTypeStopLoss        EventMarketDataKlineParamsPayloadOrderType = "STOP_LOSS"
	EventMarketDataKlineParamsPayloadOrderTypeStopLossLimit   EventMarketDataKlineParamsPayloadOrderType = "STOP_LOSS_LIMIT"
	EventMarketDataKlineParamsPayloadOrderTypeTakeProfit      EventMarketDataKlineParamsPayloadOrderType = "TAKE_PROFIT"
	EventMarketDataKlineParamsPayloadOrderTypeTakeProfitLimit EventMarketDataKlineParamsPayloadOrderType = "TAKE_PROFIT_LIMIT"
	EventMarketDataKlineParamsPayloadOrderTypeQuoted          EventMarketDataKlineParamsPayloadOrderType = "QUOTED"
)

func (r EventMarketDataKlineParamsPayloadOrderType) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsPayloadOrderTypeMarket, EventMarketDataKlineParamsPayloadOrderTypeLimit, EventMarketDataKlineParamsPayloadOrderTypeStopLoss, EventMarketDataKlineParamsPayloadOrderTypeStopLossLimit, EventMarketDataKlineParamsPayloadOrderTypeTakeProfit, EventMarketDataKlineParamsPayloadOrderTypeTakeProfitLimit, EventMarketDataKlineParamsPayloadOrderTypeQuoted:
		return true
	}
	return false
}

// Time in force
type EventMarketDataKlineParamsPayloadTimeInForce string

const (
	EventMarketDataKlineParamsPayloadTimeInForceDay EventMarketDataKlineParamsPayloadTimeInForce = "DAY"
	EventMarketDataKlineParamsPayloadTimeInForceGtc EventMarketDataKlineParamsPayloadTimeInForce = "GTC"
	EventMarketDataKlineParamsPayloadTimeInForceGtx EventMarketDataKlineParamsPayloadTimeInForce = "GTX"
	EventMarketDataKlineParamsPayloadTimeInForceGtd EventMarketDataKlineParamsPayloadTimeInForce = "GTD"
	EventMarketDataKlineParamsPayloadTimeInForceOpg EventMarketDataKlineParamsPayloadTimeInForce = "OPG"
	EventMarketDataKlineParamsPayloadTimeInForceCls EventMarketDataKlineParamsPayloadTimeInForce = "CLS"
	EventMarketDataKlineParamsPayloadTimeInForceIoc EventMarketDataKlineParamsPayloadTimeInForce = "IOC"
	EventMarketDataKlineParamsPayloadTimeInForceFok EventMarketDataKlineParamsPayloadTimeInForce = "FOK"
	EventMarketDataKlineParamsPayloadTimeInForceGfa EventMarketDataKlineParamsPayloadTimeInForce = "GFA"
	EventMarketDataKlineParamsPayloadTimeInForceGfs EventMarketDataKlineParamsPayloadTimeInForce = "GFS"
	EventMarketDataKlineParamsPayloadTimeInForceGtm EventMarketDataKlineParamsPayloadTimeInForce = "GTM"
	EventMarketDataKlineParamsPayloadTimeInForceMoo EventMarketDataKlineParamsPayloadTimeInForce = "MOO"
	EventMarketDataKlineParamsPayloadTimeInForceMoc EventMarketDataKlineParamsPayloadTimeInForce = "MOC"
	EventMarketDataKlineParamsPayloadTimeInForceExt EventMarketDataKlineParamsPayloadTimeInForce = "EXT"
)

func (r EventMarketDataKlineParamsPayloadTimeInForce) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsPayloadTimeInForceDay, EventMarketDataKlineParamsPayloadTimeInForceGtc, EventMarketDataKlineParamsPayloadTimeInForceGtx, EventMarketDataKlineParamsPayloadTimeInForceGtd, EventMarketDataKlineParamsPayloadTimeInForceOpg, EventMarketDataKlineParamsPayloadTimeInForceCls, EventMarketDataKlineParamsPayloadTimeInForceIoc, EventMarketDataKlineParamsPayloadTimeInForceFok, EventMarketDataKlineParamsPayloadTimeInForceGfa, EventMarketDataKlineParamsPayloadTimeInForceGfs, EventMarketDataKlineParamsPayloadTimeInForceGtm, EventMarketDataKlineParamsPayloadTimeInForceMoo, EventMarketDataKlineParamsPayloadTimeInForceMoc, EventMarketDataKlineParamsPayloadTimeInForceExt:
		return true
	}
	return false
}

// Route policy. For PRIORITY, the order request will be routed to the exchange
// account with the highest priority. For QUOTE, the system will execute the
// execution plan based on the quote. Order request with route policy QUOTE will
// only accept two parameters, quoteRequestId and priceSlippageTolerance
type EventMarketDataKlineParamsPayloadRoutePolicy string

const (
	EventMarketDataKlineParamsPayloadRoutePolicyPriority EventMarketDataKlineParamsPayloadRoutePolicy = "PRIORITY"
	EventMarketDataKlineParamsPayloadRoutePolicyQuote    EventMarketDataKlineParamsPayloadRoutePolicy = "QUOTE"
)

func (r EventMarketDataKlineParamsPayloadRoutePolicy) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsPayloadRoutePolicyPriority, EventMarketDataKlineParamsPayloadRoutePolicyQuote:
		return true
	}
	return false
}

// Order status
type EventMarketDataKlineParamsPayloadStatus string

const (
	EventMarketDataKlineParamsPayloadStatusSubmitted       EventMarketDataKlineParamsPayloadStatus = "SUBMITTED"
	EventMarketDataKlineParamsPayloadStatusAccepted        EventMarketDataKlineParamsPayloadStatus = "ACCEPTED"
	EventMarketDataKlineParamsPayloadStatusOpen            EventMarketDataKlineParamsPayloadStatus = "OPEN"
	EventMarketDataKlineParamsPayloadStatusPartiallyFilled EventMarketDataKlineParamsPayloadStatus = "PARTIALLY_FILLED"
	EventMarketDataKlineParamsPayloadStatusFilled          EventMarketDataKlineParamsPayloadStatus = "FILLED"
	EventMarketDataKlineParamsPayloadStatusCanceled        EventMarketDataKlineParamsPayloadStatus = "CANCELED"
	EventMarketDataKlineParamsPayloadStatusPendingCancel   EventMarketDataKlineParamsPayloadStatus = "PENDING_CANCEL"
	EventMarketDataKlineParamsPayloadStatusRejected        EventMarketDataKlineParamsPayloadStatus = "REJECTED"
	EventMarketDataKlineParamsPayloadStatusExpired         EventMarketDataKlineParamsPayloadStatus = "EXPIRED"
	EventMarketDataKlineParamsPayloadStatusRevoked         EventMarketDataKlineParamsPayloadStatus = "REVOKED"
)

func (r EventMarketDataKlineParamsPayloadStatus) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsPayloadStatusSubmitted, EventMarketDataKlineParamsPayloadStatusAccepted, EventMarketDataKlineParamsPayloadStatusOpen, EventMarketDataKlineParamsPayloadStatusPartiallyFilled, EventMarketDataKlineParamsPayloadStatusFilled, EventMarketDataKlineParamsPayloadStatusCanceled, EventMarketDataKlineParamsPayloadStatusPendingCancel, EventMarketDataKlineParamsPayloadStatusRejected, EventMarketDataKlineParamsPayloadStatusExpired, EventMarketDataKlineParamsPayloadStatusRevoked:
		return true
	}
	return false
}

type EventMarketDataKlineParamsPayloadInterval string

const (
	EventMarketDataKlineParamsPayloadInterval1s  EventMarketDataKlineParamsPayloadInterval = "1s"
	EventMarketDataKlineParamsPayloadInterval1m  EventMarketDataKlineParamsPayloadInterval = "1m"
	EventMarketDataKlineParamsPayloadInterval5m  EventMarketDataKlineParamsPayloadInterval = "5m"
	EventMarketDataKlineParamsPayloadInterval15m EventMarketDataKlineParamsPayloadInterval = "15m"
	EventMarketDataKlineParamsPayloadInterval30m EventMarketDataKlineParamsPayloadInterval = "30m"
	EventMarketDataKlineParamsPayloadInterval1h  EventMarketDataKlineParamsPayloadInterval = "1h"
	EventMarketDataKlineParamsPayloadInterval2h  EventMarketDataKlineParamsPayloadInterval = "2h"
	EventMarketDataKlineParamsPayloadInterval1d  EventMarketDataKlineParamsPayloadInterval = "1d"
	EventMarketDataKlineParamsPayloadInterval1w  EventMarketDataKlineParamsPayloadInterval = "1w"
)

func (r EventMarketDataKlineParamsPayloadInterval) IsKnown() bool {
	switch r {
	case EventMarketDataKlineParamsPayloadInterval1s, EventMarketDataKlineParamsPayloadInterval1m, EventMarketDataKlineParamsPayloadInterval5m, EventMarketDataKlineParamsPayloadInterval15m, EventMarketDataKlineParamsPayloadInterval30m, EventMarketDataKlineParamsPayloadInterval1h, EventMarketDataKlineParamsPayloadInterval2h, EventMarketDataKlineParamsPayloadInterval1d, EventMarketDataKlineParamsPayloadInterval1w:
		return true
	}
	return false
}

type EventMarketDataOrderBookParams struct {
	// A unique identifier for the event.
	EventID   param.Field[string]                                  `json:"eventId,required"`
	EventType param.Field[EventMarketDataOrderBookParamsEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]               `json:"timestamp,required"`
	Payload   param.Field[OrderbookUnionParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventMarketDataOrderBookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventMarketDataOrderBookParamsEventType string

const (
	EventMarketDataOrderBookParamsEventTypeCadenzaMarketDataOrderBook       EventMarketDataOrderBookParamsEventType = "cadenza.marketData.orderBook"
	EventMarketDataOrderBookParamsEventTypeCadenzaTaskPlaceOrderRequestAck  EventMarketDataOrderBookParamsEventType = "cadenza.task.placeOrderRequestAck"
	EventMarketDataOrderBookParamsEventTypeCadenzaTaskCancelOrderRequestAck EventMarketDataOrderBookParamsEventType = "cadenza.task.cancelOrderRequestAck"
	EventMarketDataOrderBookParamsEventTypeCadenzaDropCopyQuote             EventMarketDataOrderBookParamsEventType = "cadenza.dropCopy.quote"
	EventMarketDataOrderBookParamsEventTypeCadenzaDropCopyOrder             EventMarketDataOrderBookParamsEventType = "cadenza.dropCopy.order"
	EventMarketDataOrderBookParamsEventTypeCadenzaDropCopyPortfolio         EventMarketDataOrderBookParamsEventType = "cadenza.dropCopy.portfolio"
	EventMarketDataOrderBookParamsEventTypeCadenzaMarketDataKline           EventMarketDataOrderBookParamsEventType = "cadenza.marketData.kline"
)

func (r EventMarketDataOrderBookParamsEventType) IsKnown() bool {
	switch r {
	case EventMarketDataOrderBookParamsEventTypeCadenzaMarketDataOrderBook, EventMarketDataOrderBookParamsEventTypeCadenzaTaskPlaceOrderRequestAck, EventMarketDataOrderBookParamsEventTypeCadenzaTaskCancelOrderRequestAck, EventMarketDataOrderBookParamsEventTypeCadenzaDropCopyQuote, EventMarketDataOrderBookParamsEventTypeCadenzaDropCopyOrder, EventMarketDataOrderBookParamsEventTypeCadenzaDropCopyPortfolio, EventMarketDataOrderBookParamsEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}
