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

// EventMarketDataService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventMarketDataService] method instead.
type EventMarketDataService struct {
	Options []option.RequestOption
}

// NewEventMarketDataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventMarketDataService(opts ...option.RequestOption) (r *EventMarketDataService) {
	r = &EventMarketDataService{}
	r.Options = opts
	return
}

// PubSub event handler placeholder for kline event
func (r *EventMarketDataService) MarketDataKline(ctx context.Context, body EventMarketDataMarketDataKlineParams, opts ...option.RequestOption) (res *MarketDataKline, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/kline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for order book event
func (r *EventMarketDataService) MarketDataOrderBook(ctx context.Context, body EventMarketDataMarketDataOrderBookParams, opts ...option.RequestOption) (res *MarketDataOrderBook, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/orderBook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MarketDataKline struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType MarketDataKlineEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64                  `json:"timestamp,required"`
	Payload   MarketDataKlinePayload `json:"payload"`
	// The source system or module that generated the event.
	Source string              `json:"source"`
	JSON   marketDataKlineJSON `json:"-"`
}

// marketDataKlineJSON contains the JSON metadata for the struct [MarketDataKline]
type marketDataKlineJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketDataKline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataKlineJSON) RawJSON() string {
	return r.raw
}

// Event Type
type MarketDataKlineEventType string

const (
	MarketDataKlineEventTypeCadenzaTaskQuoteRequestAck       MarketDataKlineEventType = "cadenza.task.quoteRequestAck"
	MarketDataKlineEventTypeCadenzaTaskPlaceOrderRequestAck  MarketDataKlineEventType = "cadenza.task.placeOrderRequestAck"
	MarketDataKlineEventTypeCadenzaTaskCancelOrderRequestAck MarketDataKlineEventType = "cadenza.task.cancelOrderRequestAck"
	MarketDataKlineEventTypeCadenzaDropCopyQuote             MarketDataKlineEventType = "cadenza.dropCopy.quote"
	MarketDataKlineEventTypeCadenzaDropCopyOrder             MarketDataKlineEventType = "cadenza.dropCopy.order"
	MarketDataKlineEventTypeCadenzaDropCopyPortfolio         MarketDataKlineEventType = "cadenza.dropCopy.portfolio"
	MarketDataKlineEventTypeCadenzaMarketDataOrderBook       MarketDataKlineEventType = "cadenza.marketData.orderBook"
	MarketDataKlineEventTypeCadenzaMarketDataKline           MarketDataKlineEventType = "cadenza.marketData.kline"
)

func (r MarketDataKlineEventType) IsKnown() bool {
	switch r {
	case MarketDataKlineEventTypeCadenzaTaskQuoteRequestAck, MarketDataKlineEventTypeCadenzaTaskPlaceOrderRequestAck, MarketDataKlineEventTypeCadenzaTaskCancelOrderRequestAck, MarketDataKlineEventTypeCadenzaDropCopyQuote, MarketDataKlineEventTypeCadenzaDropCopyOrder, MarketDataKlineEventTypeCadenzaDropCopyPortfolio, MarketDataKlineEventTypeCadenzaMarketDataOrderBook, MarketDataKlineEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
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
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[MarketDataKlineEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                       `json:"timestamp,required"`
	Payload   param.Field[MarketDataKlinePayloadParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
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
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType MarketDataOrderBookEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64     `json:"timestamp,required"`
	Payload   Orderbook `json:"payload"`
	// The source system or module that generated the event.
	Source string                  `json:"source"`
	JSON   marketDataOrderBookJSON `json:"-"`
}

// marketDataOrderBookJSON contains the JSON metadata for the struct
// [MarketDataOrderBook]
type marketDataOrderBookJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketDataOrderBook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataOrderBookJSON) RawJSON() string {
	return r.raw
}

// Event Type
type MarketDataOrderBookEventType string

const (
	MarketDataOrderBookEventTypeCadenzaTaskQuoteRequestAck       MarketDataOrderBookEventType = "cadenza.task.quoteRequestAck"
	MarketDataOrderBookEventTypeCadenzaTaskPlaceOrderRequestAck  MarketDataOrderBookEventType = "cadenza.task.placeOrderRequestAck"
	MarketDataOrderBookEventTypeCadenzaTaskCancelOrderRequestAck MarketDataOrderBookEventType = "cadenza.task.cancelOrderRequestAck"
	MarketDataOrderBookEventTypeCadenzaDropCopyQuote             MarketDataOrderBookEventType = "cadenza.dropCopy.quote"
	MarketDataOrderBookEventTypeCadenzaDropCopyOrder             MarketDataOrderBookEventType = "cadenza.dropCopy.order"
	MarketDataOrderBookEventTypeCadenzaDropCopyPortfolio         MarketDataOrderBookEventType = "cadenza.dropCopy.portfolio"
	MarketDataOrderBookEventTypeCadenzaMarketDataOrderBook       MarketDataOrderBookEventType = "cadenza.marketData.orderBook"
	MarketDataOrderBookEventTypeCadenzaMarketDataKline           MarketDataOrderBookEventType = "cadenza.marketData.kline"
)

func (r MarketDataOrderBookEventType) IsKnown() bool {
	switch r {
	case MarketDataOrderBookEventTypeCadenzaTaskQuoteRequestAck, MarketDataOrderBookEventTypeCadenzaTaskPlaceOrderRequestAck, MarketDataOrderBookEventTypeCadenzaTaskCancelOrderRequestAck, MarketDataOrderBookEventTypeCadenzaDropCopyQuote, MarketDataOrderBookEventTypeCadenzaDropCopyOrder, MarketDataOrderBookEventTypeCadenzaDropCopyPortfolio, MarketDataOrderBookEventTypeCadenzaMarketDataOrderBook, MarketDataOrderBookEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type MarketDataOrderBookParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[MarketDataOrderBookEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]          `json:"timestamp,required"`
	Payload   param.Field[OrderbookParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r MarketDataOrderBookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventMarketDataMarketDataKlineParams struct {
	MarketDataKline MarketDataKlineParam `json:"marketDataKline,required"`
}

func (r EventMarketDataMarketDataKlineParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MarketDataKline)
}

type EventMarketDataMarketDataOrderBookParams struct {
	MarketDataOrderBook MarketDataOrderBookParam `json:"marketDataOrderBook,required"`
}

func (r EventMarketDataMarketDataOrderBookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MarketDataOrderBook)
}
