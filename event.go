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

// PubSub event handler for execution report drop copy event
func (r *EventService) DropCopyExecutionReport(ctx context.Context, body EventDropCopyExecutionReportParams, opts ...option.RequestOption) (res *DropCopyExecutionReport, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/executionReport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for order event
func (r *EventService) DropCopyOrder(ctx context.Context, body EventDropCopyOrderParams, opts ...option.RequestOption) (res *DropCopyOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for portfolio event
func (r *EventService) DropCopyPortfolio(ctx context.Context, body EventDropCopyPortfolioParams, opts ...option.RequestOption) (res *DropCopyPortfolio, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/portfolio"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for quote event
func (r *EventService) DropCopyQuote(ctx context.Context, body EventDropCopyQuoteParams, opts ...option.RequestOption) (res *DropCopyQuote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for kline event
func (r *EventService) MarketDataKline(ctx context.Context, body EventMarketDataKlineParams, opts ...option.RequestOption) (res *MarketDataKline, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/kline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for order book event
func (r *EventService) MarketDataOrderBook(ctx context.Context, body EventMarketDataOrderBookParams, opts ...option.RequestOption) (res *MarketDataOrderBook, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/orderBook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder
func (r *EventService) New(ctx context.Context, body EventNewParams, opts ...option.RequestOption) (res *GenericEvent, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for cancel order request acknowledgment event
func (r *EventService) TaskCancelOrderRequestAck(ctx context.Context, body EventTaskCancelOrderRequestAckParams, opts ...option.RequestOption) (res *TaskCancelOrderRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/task/cancelOrderRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for place order request acknowledgment event
func (r *EventService) TaskPlaceOrderRequestAck(ctx context.Context, body EventTaskPlaceOrderRequestAckParams, opts ...option.RequestOption) (res *TaskPlaceOrderRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/task/placeOrderRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for quote request acknowledgment event
func (r *EventService) TaskQuoteRequestAck(ctx context.Context, body EventTaskQuoteRequestAckParams, opts ...option.RequestOption) (res *TaskQuoteRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/task/quoteRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DropCopyExecutionReport struct {
	Payload ExecutionReport             `json:"payload"`
	JSON    dropCopyExecutionReportJSON `json:"-"`
	Event
}

// dropCopyExecutionReportJSON contains the JSON metadata for the struct
// [DropCopyExecutionReport]
type dropCopyExecutionReportJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyExecutionReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyExecutionReportJSON) RawJSON() string {
	return r.raw
}

type DropCopyExecutionReportParam struct {
	Payload param.Field[ExecutionReportParam] `json:"payload"`
	EventParam
}

func (r DropCopyExecutionReportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyOrder struct {
	Payload Order             `json:"payload"`
	JSON    dropCopyOrderJSON `json:"-"`
	Event
}

// dropCopyOrderJSON contains the JSON metadata for the struct [DropCopyOrder]
type dropCopyOrderJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyOrderJSON) RawJSON() string {
	return r.raw
}

type DropCopyOrderParam struct {
	Payload param.Field[OrderParam] `json:"payload"`
	EventParam
}

func (r DropCopyOrderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyPortfolio struct {
	Payload ExchangeAccountPortfolio `json:"payload"`
	JSON    dropCopyPortfolioJSON    `json:"-"`
	Event
}

// dropCopyPortfolioJSON contains the JSON metadata for the struct
// [DropCopyPortfolio]
type dropCopyPortfolioJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyPortfolio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyPortfolioJSON) RawJSON() string {
	return r.raw
}

type DropCopyPortfolioParam struct {
	Payload param.Field[ExchangeAccountPortfolioParam] `json:"payload"`
	EventParam
}

func (r DropCopyPortfolioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyQuote struct {
	Payload Quote             `json:"payload"`
	JSON    dropCopyQuoteJSON `json:"-"`
	Event
}

// dropCopyQuoteJSON contains the JSON metadata for the struct [DropCopyQuote]
type dropCopyQuoteJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyQuote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyQuoteJSON) RawJSON() string {
	return r.raw
}

type DropCopyQuoteParam struct {
	Payload param.Field[QuoteParam] `json:"payload"`
	EventParam
}

func (r DropCopyQuoteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GenericEvent struct {
	// The actual data of the event, which varies based on the event type.
	Payload GenericEventPayload `json:"payload"`
	JSON    genericEventJSON    `json:"-"`
	Event
}

// genericEventJSON contains the JSON metadata for the struct [GenericEvent]
type genericEventJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GenericEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r genericEventJSON) RawJSON() string {
	return r.raw
}

// The actual data of the event, which varies based on the event type.
type GenericEventPayload struct {
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
	OrderType GenericEventPayloadOrderType `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID string `json:"positionId" format:"uuid"`
	// Price
	Price float64 `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance float64 `json:"priceSlippageTolerance"`
	// Symbol
	Symbol string `json:"symbol"`
	// Time in force
	TimeInForce GenericEventPayloadTimeInForce `json:"timeInForce"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy GenericEventPayloadRoutePolicy `json:"routePolicy"`
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
	Status GenericEventPayloadStatus `json:"status"`
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
	Level    int64                       `json:"level"`
	Interval GenericEventPayloadInterval `json:"interval"`
	// This field can have the runtime type of [[]Ohlcv].
	Candles interface{}             `json:"candles,required"`
	JSON    genericEventPayloadJSON `json:"-"`
	union   GenericEventPayloadUnion
}

// genericEventPayloadJSON contains the JSON metadata for the struct
// [GenericEventPayload]
type genericEventPayloadJSON struct {
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

func (r genericEventPayloadJSON) RawJSON() string {
	return r.raw
}

func (r *GenericEventPayload) UnmarshalJSON(data []byte) (err error) {
	*r = GenericEventPayload{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [GenericEventPayloadUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [QuoteRequest], [PlaceOrderRequest],
// [CancelOrderRequest], [Quote], [Order], [ExecutionReport],
// [ExchangeAccountPortfolio], [Orderbook], [GenericEventPayloadKline].
func (r GenericEventPayload) AsUnion() GenericEventPayloadUnion {
	return r.union
}

// The actual data of the event, which varies based on the event type.
//
// Union satisfied by [QuoteRequest], [PlaceOrderRequest], [CancelOrderRequest],
// [Quote], [Order], [ExecutionReport], [ExchangeAccountPortfolio], [Orderbook] or
// [GenericEventPayloadKline].
type GenericEventPayloadUnion interface {
	implementsGenericEventPayload()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GenericEventPayloadUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(GenericEventPayloadKline{}),
		},
	)
}

type GenericEventPayloadKline struct {
	Candles []Ohlcv `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType GenericEventPayloadKlineExchangeType `json:"exchangeType"`
	Interval     GenericEventPayloadKlineInterval     `json:"interval"`
	Symbol       string                               `json:"symbol"`
	JSON         genericEventPayloadKlineJSON         `json:"-"`
}

// genericEventPayloadKlineJSON contains the JSON metadata for the struct
// [GenericEventPayloadKline]
type genericEventPayloadKlineJSON struct {
	Candles           apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Interval          apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GenericEventPayloadKline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r genericEventPayloadKlineJSON) RawJSON() string {
	return r.raw
}

func (r GenericEventPayloadKline) implementsGenericEventPayload() {}

// Exchange type
type GenericEventPayloadKlineExchangeType string

const (
	GenericEventPayloadKlineExchangeTypeBinance       GenericEventPayloadKlineExchangeType = "BINANCE"
	GenericEventPayloadKlineExchangeTypeBinanceMargin GenericEventPayloadKlineExchangeType = "BINANCE_MARGIN"
	GenericEventPayloadKlineExchangeTypeB2C2          GenericEventPayloadKlineExchangeType = "B2C2"
	GenericEventPayloadKlineExchangeTypeWintermute    GenericEventPayloadKlineExchangeType = "WINTERMUTE"
	GenericEventPayloadKlineExchangeTypeBlockfills    GenericEventPayloadKlineExchangeType = "BLOCKFILLS"
	GenericEventPayloadKlineExchangeTypeStonex        GenericEventPayloadKlineExchangeType = "STONEX"
)

func (r GenericEventPayloadKlineExchangeType) IsKnown() bool {
	switch r {
	case GenericEventPayloadKlineExchangeTypeBinance, GenericEventPayloadKlineExchangeTypeBinanceMargin, GenericEventPayloadKlineExchangeTypeB2C2, GenericEventPayloadKlineExchangeTypeWintermute, GenericEventPayloadKlineExchangeTypeBlockfills, GenericEventPayloadKlineExchangeTypeStonex:
		return true
	}
	return false
}

type GenericEventPayloadKlineInterval string

const (
	GenericEventPayloadKlineInterval1s  GenericEventPayloadKlineInterval = "1s"
	GenericEventPayloadKlineInterval1m  GenericEventPayloadKlineInterval = "1m"
	GenericEventPayloadKlineInterval5m  GenericEventPayloadKlineInterval = "5m"
	GenericEventPayloadKlineInterval15m GenericEventPayloadKlineInterval = "15m"
	GenericEventPayloadKlineInterval30m GenericEventPayloadKlineInterval = "30m"
	GenericEventPayloadKlineInterval1h  GenericEventPayloadKlineInterval = "1h"
	GenericEventPayloadKlineInterval2h  GenericEventPayloadKlineInterval = "2h"
	GenericEventPayloadKlineInterval1d  GenericEventPayloadKlineInterval = "1d"
	GenericEventPayloadKlineInterval1w  GenericEventPayloadKlineInterval = "1w"
)

func (r GenericEventPayloadKlineInterval) IsKnown() bool {
	switch r {
	case GenericEventPayloadKlineInterval1s, GenericEventPayloadKlineInterval1m, GenericEventPayloadKlineInterval5m, GenericEventPayloadKlineInterval15m, GenericEventPayloadKlineInterval30m, GenericEventPayloadKlineInterval1h, GenericEventPayloadKlineInterval2h, GenericEventPayloadKlineInterval1d, GenericEventPayloadKlineInterval1w:
		return true
	}
	return false
}

// Order type
type GenericEventPayloadOrderType string

const (
	GenericEventPayloadOrderTypeMarket          GenericEventPayloadOrderType = "MARKET"
	GenericEventPayloadOrderTypeLimit           GenericEventPayloadOrderType = "LIMIT"
	GenericEventPayloadOrderTypeStopLoss        GenericEventPayloadOrderType = "STOP_LOSS"
	GenericEventPayloadOrderTypeStopLossLimit   GenericEventPayloadOrderType = "STOP_LOSS_LIMIT"
	GenericEventPayloadOrderTypeTakeProfit      GenericEventPayloadOrderType = "TAKE_PROFIT"
	GenericEventPayloadOrderTypeTakeProfitLimit GenericEventPayloadOrderType = "TAKE_PROFIT_LIMIT"
	GenericEventPayloadOrderTypeQuoted          GenericEventPayloadOrderType = "QUOTED"
)

func (r GenericEventPayloadOrderType) IsKnown() bool {
	switch r {
	case GenericEventPayloadOrderTypeMarket, GenericEventPayloadOrderTypeLimit, GenericEventPayloadOrderTypeStopLoss, GenericEventPayloadOrderTypeStopLossLimit, GenericEventPayloadOrderTypeTakeProfit, GenericEventPayloadOrderTypeTakeProfitLimit, GenericEventPayloadOrderTypeQuoted:
		return true
	}
	return false
}

// Time in force
type GenericEventPayloadTimeInForce string

const (
	GenericEventPayloadTimeInForceDay GenericEventPayloadTimeInForce = "DAY"
	GenericEventPayloadTimeInForceGtc GenericEventPayloadTimeInForce = "GTC"
	GenericEventPayloadTimeInForceGtx GenericEventPayloadTimeInForce = "GTX"
	GenericEventPayloadTimeInForceGtd GenericEventPayloadTimeInForce = "GTD"
	GenericEventPayloadTimeInForceOpg GenericEventPayloadTimeInForce = "OPG"
	GenericEventPayloadTimeInForceCls GenericEventPayloadTimeInForce = "CLS"
	GenericEventPayloadTimeInForceIoc GenericEventPayloadTimeInForce = "IOC"
	GenericEventPayloadTimeInForceFok GenericEventPayloadTimeInForce = "FOK"
	GenericEventPayloadTimeInForceGfa GenericEventPayloadTimeInForce = "GFA"
	GenericEventPayloadTimeInForceGfs GenericEventPayloadTimeInForce = "GFS"
	GenericEventPayloadTimeInForceGtm GenericEventPayloadTimeInForce = "GTM"
	GenericEventPayloadTimeInForceMoo GenericEventPayloadTimeInForce = "MOO"
	GenericEventPayloadTimeInForceMoc GenericEventPayloadTimeInForce = "MOC"
	GenericEventPayloadTimeInForceExt GenericEventPayloadTimeInForce = "EXT"
)

func (r GenericEventPayloadTimeInForce) IsKnown() bool {
	switch r {
	case GenericEventPayloadTimeInForceDay, GenericEventPayloadTimeInForceGtc, GenericEventPayloadTimeInForceGtx, GenericEventPayloadTimeInForceGtd, GenericEventPayloadTimeInForceOpg, GenericEventPayloadTimeInForceCls, GenericEventPayloadTimeInForceIoc, GenericEventPayloadTimeInForceFok, GenericEventPayloadTimeInForceGfa, GenericEventPayloadTimeInForceGfs, GenericEventPayloadTimeInForceGtm, GenericEventPayloadTimeInForceMoo, GenericEventPayloadTimeInForceMoc, GenericEventPayloadTimeInForceExt:
		return true
	}
	return false
}

// Route policy. For PRIORITY, the order request will be routed to the exchange
// account with the highest priority. For QUOTE, the system will execute the
// execution plan based on the quote. Order request with route policy QUOTE will
// only accept two parameters, quoteRequestId and priceSlippageTolerance
type GenericEventPayloadRoutePolicy string

const (
	GenericEventPayloadRoutePolicyPriority GenericEventPayloadRoutePolicy = "PRIORITY"
	GenericEventPayloadRoutePolicyQuote    GenericEventPayloadRoutePolicy = "QUOTE"
)

func (r GenericEventPayloadRoutePolicy) IsKnown() bool {
	switch r {
	case GenericEventPayloadRoutePolicyPriority, GenericEventPayloadRoutePolicyQuote:
		return true
	}
	return false
}

// Order status
type GenericEventPayloadStatus string

const (
	GenericEventPayloadStatusSubmitted       GenericEventPayloadStatus = "SUBMITTED"
	GenericEventPayloadStatusAccepted        GenericEventPayloadStatus = "ACCEPTED"
	GenericEventPayloadStatusOpen            GenericEventPayloadStatus = "OPEN"
	GenericEventPayloadStatusPartiallyFilled GenericEventPayloadStatus = "PARTIALLY_FILLED"
	GenericEventPayloadStatusFilled          GenericEventPayloadStatus = "FILLED"
	GenericEventPayloadStatusCanceled        GenericEventPayloadStatus = "CANCELED"
	GenericEventPayloadStatusPendingCancel   GenericEventPayloadStatus = "PENDING_CANCEL"
	GenericEventPayloadStatusRejected        GenericEventPayloadStatus = "REJECTED"
	GenericEventPayloadStatusExpired         GenericEventPayloadStatus = "EXPIRED"
	GenericEventPayloadStatusRevoked         GenericEventPayloadStatus = "REVOKED"
)

func (r GenericEventPayloadStatus) IsKnown() bool {
	switch r {
	case GenericEventPayloadStatusSubmitted, GenericEventPayloadStatusAccepted, GenericEventPayloadStatusOpen, GenericEventPayloadStatusPartiallyFilled, GenericEventPayloadStatusFilled, GenericEventPayloadStatusCanceled, GenericEventPayloadStatusPendingCancel, GenericEventPayloadStatusRejected, GenericEventPayloadStatusExpired, GenericEventPayloadStatusRevoked:
		return true
	}
	return false
}

type GenericEventPayloadInterval string

const (
	GenericEventPayloadInterval1s  GenericEventPayloadInterval = "1s"
	GenericEventPayloadInterval1m  GenericEventPayloadInterval = "1m"
	GenericEventPayloadInterval5m  GenericEventPayloadInterval = "5m"
	GenericEventPayloadInterval15m GenericEventPayloadInterval = "15m"
	GenericEventPayloadInterval30m GenericEventPayloadInterval = "30m"
	GenericEventPayloadInterval1h  GenericEventPayloadInterval = "1h"
	GenericEventPayloadInterval2h  GenericEventPayloadInterval = "2h"
	GenericEventPayloadInterval1d  GenericEventPayloadInterval = "1d"
	GenericEventPayloadInterval1w  GenericEventPayloadInterval = "1w"
)

func (r GenericEventPayloadInterval) IsKnown() bool {
	switch r {
	case GenericEventPayloadInterval1s, GenericEventPayloadInterval1m, GenericEventPayloadInterval5m, GenericEventPayloadInterval15m, GenericEventPayloadInterval30m, GenericEventPayloadInterval1h, GenericEventPayloadInterval2h, GenericEventPayloadInterval1d, GenericEventPayloadInterval1w:
		return true
	}
	return false
}

type GenericEventParam struct {
	// The actual data of the event, which varies based on the event type.
	Payload param.Field[GenericEventPayloadUnionParam] `json:"payload"`
	EventParam
}

func (r GenericEventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The actual data of the event, which varies based on the event type.
type GenericEventPayloadParam struct {
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
	OrderType param.Field[GenericEventPayloadOrderType] `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID param.Field[string] `json:"positionId" format:"uuid"`
	// Price
	Price param.Field[float64] `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance param.Field[float64] `json:"priceSlippageTolerance"`
	// Symbol
	Symbol param.Field[string] `json:"symbol"`
	// Time in force
	TimeInForce param.Field[GenericEventPayloadTimeInForce] `json:"timeInForce"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy param.Field[GenericEventPayloadRoutePolicy] `json:"routePolicy"`
	Priority    param.Field[interface{}]                    `json:"priority,required"`
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
	Status param.Field[GenericEventPayloadStatus] `json:"status"`
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
	Level    param.Field[int64]                       `json:"level"`
	Interval param.Field[GenericEventPayloadInterval] `json:"interval"`
	Candles  param.Field[interface{}]                 `json:"candles,required"`
}

func (r GenericEventPayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GenericEventPayloadParam) implementsGenericEventPayloadUnionParam() {}

// The actual data of the event, which varies based on the event type.
//
// Satisfied by [QuoteRequestParam], [PlaceOrderRequestParam],
// [CancelOrderRequestParam], [QuoteParam], [OrderParam], [ExecutionReportParam],
// [ExchangeAccountPortfolioParam], [OrderbookParam],
// [GenericEventPayloadKlineParam], [GenericEventPayloadParam].
type GenericEventPayloadUnionParam interface {
	implementsGenericEventPayloadUnionParam()
}

type GenericEventPayloadKlineParam struct {
	Candles param.Field[[]OhlcvParam] `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[GenericEventPayloadKlineExchangeType] `json:"exchangeType"`
	Interval     param.Field[GenericEventPayloadKlineInterval]     `json:"interval"`
	Symbol       param.Field[string]                               `json:"symbol"`
}

func (r GenericEventPayloadKlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GenericEventPayloadKlineParam) implementsGenericEventPayloadUnionParam() {}

type MarketDataKline struct {
	Payload MarketDataKlinePayload `json:"payload"`
	JSON    marketDataKlineJSON    `json:"-"`
	Event
}

// marketDataKlineJSON contains the JSON metadata for the struct [MarketDataKline]
type marketDataKlineJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketDataKline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataKlineJSON) RawJSON() string {
	return r.raw
}

type MarketDataKlinePayload struct {
	Candles []Ohlcv `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType MarketDataKlinePayloadExchangeType `json:"exchangeType"`
	Interval     MarketDataKlinePayloadInterval     `json:"interval"`
	Symbol       string                             `json:"symbol"`
	JSON         marketDataKlinePayloadJSON         `json:"-"`
}

// marketDataKlinePayloadJSON contains the JSON metadata for the struct
// [MarketDataKlinePayload]
type marketDataKlinePayloadJSON struct {
	Candles           apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Interval          apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MarketDataKlinePayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataKlinePayloadJSON) RawJSON() string {
	return r.raw
}

// Exchange type
type MarketDataKlinePayloadExchangeType string

const (
	MarketDataKlinePayloadExchangeTypeBinance       MarketDataKlinePayloadExchangeType = "BINANCE"
	MarketDataKlinePayloadExchangeTypeBinanceMargin MarketDataKlinePayloadExchangeType = "BINANCE_MARGIN"
	MarketDataKlinePayloadExchangeTypeB2C2          MarketDataKlinePayloadExchangeType = "B2C2"
	MarketDataKlinePayloadExchangeTypeWintermute    MarketDataKlinePayloadExchangeType = "WINTERMUTE"
	MarketDataKlinePayloadExchangeTypeBlockfills    MarketDataKlinePayloadExchangeType = "BLOCKFILLS"
	MarketDataKlinePayloadExchangeTypeStonex        MarketDataKlinePayloadExchangeType = "STONEX"
)

func (r MarketDataKlinePayloadExchangeType) IsKnown() bool {
	switch r {
	case MarketDataKlinePayloadExchangeTypeBinance, MarketDataKlinePayloadExchangeTypeBinanceMargin, MarketDataKlinePayloadExchangeTypeB2C2, MarketDataKlinePayloadExchangeTypeWintermute, MarketDataKlinePayloadExchangeTypeBlockfills, MarketDataKlinePayloadExchangeTypeStonex:
		return true
	}
	return false
}

type MarketDataKlinePayloadInterval string

const (
	MarketDataKlinePayloadInterval1s  MarketDataKlinePayloadInterval = "1s"
	MarketDataKlinePayloadInterval1m  MarketDataKlinePayloadInterval = "1m"
	MarketDataKlinePayloadInterval5m  MarketDataKlinePayloadInterval = "5m"
	MarketDataKlinePayloadInterval15m MarketDataKlinePayloadInterval = "15m"
	MarketDataKlinePayloadInterval30m MarketDataKlinePayloadInterval = "30m"
	MarketDataKlinePayloadInterval1h  MarketDataKlinePayloadInterval = "1h"
	MarketDataKlinePayloadInterval2h  MarketDataKlinePayloadInterval = "2h"
	MarketDataKlinePayloadInterval1d  MarketDataKlinePayloadInterval = "1d"
	MarketDataKlinePayloadInterval1w  MarketDataKlinePayloadInterval = "1w"
)

func (r MarketDataKlinePayloadInterval) IsKnown() bool {
	switch r {
	case MarketDataKlinePayloadInterval1s, MarketDataKlinePayloadInterval1m, MarketDataKlinePayloadInterval5m, MarketDataKlinePayloadInterval15m, MarketDataKlinePayloadInterval30m, MarketDataKlinePayloadInterval1h, MarketDataKlinePayloadInterval2h, MarketDataKlinePayloadInterval1d, MarketDataKlinePayloadInterval1w:
		return true
	}
	return false
}

type MarketDataKlineParam struct {
	Payload param.Field[MarketDataKlinePayloadParam] `json:"payload"`
	EventParam
}

func (r MarketDataKlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarketDataKlinePayloadParam struct {
	Candles param.Field[[]OhlcvParam] `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[MarketDataKlinePayloadExchangeType] `json:"exchangeType"`
	Interval     param.Field[MarketDataKlinePayloadInterval]     `json:"interval"`
	Symbol       param.Field[string]                             `json:"symbol"`
}

func (r MarketDataKlinePayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarketDataOrderBook struct {
	Payload Orderbook               `json:"payload"`
	JSON    marketDataOrderBookJSON `json:"-"`
	Event
}

// marketDataOrderBookJSON contains the JSON metadata for the struct
// [MarketDataOrderBook]
type marketDataOrderBookJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketDataOrderBook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataOrderBookJSON) RawJSON() string {
	return r.raw
}

type MarketDataOrderBookParam struct {
	Payload param.Field[OrderbookParam] `json:"payload"`
	EventParam
}

func (r MarketDataOrderBookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskCancelOrderRequestAck struct {
	Payload CancelOrderRequest            `json:"payload"`
	JSON    taskCancelOrderRequestAckJSON `json:"-"`
	Event
}

// taskCancelOrderRequestAckJSON contains the JSON metadata for the struct
// [TaskCancelOrderRequestAck]
type taskCancelOrderRequestAckJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskCancelOrderRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskCancelOrderRequestAckJSON) RawJSON() string {
	return r.raw
}

type TaskCancelOrderRequestAckParam struct {
	Payload param.Field[CancelOrderRequestParam] `json:"payload"`
	EventParam
}

func (r TaskCancelOrderRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskPlaceOrderRequestAck struct {
	Payload PlaceOrderRequest            `json:"payload"`
	JSON    taskPlaceOrderRequestAckJSON `json:"-"`
	Event
}

// taskPlaceOrderRequestAckJSON contains the JSON metadata for the struct
// [TaskPlaceOrderRequestAck]
type taskPlaceOrderRequestAckJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskPlaceOrderRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskPlaceOrderRequestAckJSON) RawJSON() string {
	return r.raw
}

type TaskPlaceOrderRequestAckParam struct {
	Payload param.Field[PlaceOrderRequestParam] `json:"payload"`
	EventParam
}

func (r TaskPlaceOrderRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskQuoteRequestAck struct {
	Payload QuoteRequest            `json:"payload"`
	JSON    taskQuoteRequestAckJSON `json:"-"`
	Event
}

// taskQuoteRequestAckJSON contains the JSON metadata for the struct
// [TaskQuoteRequestAck]
type taskQuoteRequestAckJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskQuoteRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskQuoteRequestAckJSON) RawJSON() string {
	return r.raw
}

type TaskQuoteRequestAckParam struct {
	Payload param.Field[QuoteRequestParam] `json:"payload"`
	EventParam
}

func (r TaskQuoteRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventDropCopyExecutionReportParams struct {
	DropCopyExecutionReport DropCopyExecutionReportParam `json:"dropCopyExecutionReport,required"`
}

func (r EventDropCopyExecutionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyExecutionReport)
}

type EventDropCopyOrderParams struct {
	DropCopyOrder DropCopyOrderParam `json:"dropCopyOrder,required"`
}

func (r EventDropCopyOrderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyOrder)
}

type EventDropCopyPortfolioParams struct {
	DropCopyPortfolio DropCopyPortfolioParam `json:"dropCopyPortfolio,required"`
}

func (r EventDropCopyPortfolioParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyPortfolio)
}

type EventDropCopyQuoteParams struct {
	DropCopyQuote DropCopyQuoteParam `json:"dropCopyQuote,required"`
}

func (r EventDropCopyQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyQuote)
}

type EventMarketDataKlineParams struct {
	MarketDataKline MarketDataKlineParam `json:"marketDataKline,required"`
}

func (r EventMarketDataKlineParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MarketDataKline)
}

type EventMarketDataOrderBookParams struct {
	MarketDataOrderBook MarketDataOrderBookParam `json:"marketDataOrderBook,required"`
}

func (r EventMarketDataOrderBookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MarketDataOrderBook)
}

type EventNewParams struct {
	GenericEvent GenericEventParam `json:"genericEvent,required"`
}

func (r EventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.GenericEvent)
}

type EventTaskCancelOrderRequestAckParams struct {
	TaskCancelOrderRequestAck TaskCancelOrderRequestAckParam `json:"taskCancelOrderRequestAck,required"`
}

func (r EventTaskCancelOrderRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TaskCancelOrderRequestAck)
}

type EventTaskPlaceOrderRequestAckParams struct {
	TaskPlaceOrderRequestAck TaskPlaceOrderRequestAckParam `json:"taskPlaceOrderRequestAck,required"`
}

func (r EventTaskPlaceOrderRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TaskPlaceOrderRequestAck)
}

type EventTaskQuoteRequestAckParams struct {
	TaskQuoteRequestAck TaskQuoteRequestAckParam `json:"taskQuoteRequestAck,required"`
}

func (r EventTaskQuoteRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TaskQuoteRequestAck)
}
