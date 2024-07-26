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
	Timestamp int64 `json:"timestamp,required"`
	Payload   Kline `json:"payload"`
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
	MarketDataKlineEventTypeCadenzaTaskQuote                     MarketDataKlineEventType = "cadenza.task.quote"
	MarketDataKlineEventTypeCadenzaDropCopyQuoteRequestAck       MarketDataKlineEventType = "cadenza.dropCopy.quoteRequestAck"
	MarketDataKlineEventTypeCadenzaDropCopyPlaceOrderRequestAck  MarketDataKlineEventType = "cadenza.dropCopy.placeOrderRequestAck"
	MarketDataKlineEventTypeCadenzaDropCopyCancelOrderRequestAck MarketDataKlineEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	MarketDataKlineEventTypeCadenzaDropCopyQuote                 MarketDataKlineEventType = "cadenza.dropCopy.quote"
	MarketDataKlineEventTypeCadenzaDropCopyOrder                 MarketDataKlineEventType = "cadenza.dropCopy.order"
	MarketDataKlineEventTypeCadenzaDropCopyExecutionReport       MarketDataKlineEventType = "cadenza.dropCopy.executionReport"
	MarketDataKlineEventTypeCadenzaDropCopyPortfolio             MarketDataKlineEventType = "cadenza.dropCopy.portfolio"
	MarketDataKlineEventTypeCadenzaMarketDataOrderBook           MarketDataKlineEventType = "cadenza.marketData.orderBook"
	MarketDataKlineEventTypeCadenzaMarketDataKline               MarketDataKlineEventType = "cadenza.marketData.kline"
)

func (r MarketDataKlineEventType) IsKnown() bool {
	switch r {
	case MarketDataKlineEventTypeCadenzaTaskQuote, MarketDataKlineEventTypeCadenzaDropCopyQuoteRequestAck, MarketDataKlineEventTypeCadenzaDropCopyPlaceOrderRequestAck, MarketDataKlineEventTypeCadenzaDropCopyCancelOrderRequestAck, MarketDataKlineEventTypeCadenzaDropCopyQuote, MarketDataKlineEventTypeCadenzaDropCopyOrder, MarketDataKlineEventTypeCadenzaDropCopyExecutionReport, MarketDataKlineEventTypeCadenzaDropCopyPortfolio, MarketDataKlineEventTypeCadenzaMarketDataOrderBook, MarketDataKlineEventTypeCadenzaMarketDataKline:
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
	Timestamp param.Field[int64]      `json:"timestamp,required"`
	Payload   param.Field[KlineParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r MarketDataKlineParam) MarshalJSON() (data []byte, err error) {
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
	MarketDataOrderBookEventTypeCadenzaTaskQuote                     MarketDataOrderBookEventType = "cadenza.task.quote"
	MarketDataOrderBookEventTypeCadenzaDropCopyQuoteRequestAck       MarketDataOrderBookEventType = "cadenza.dropCopy.quoteRequestAck"
	MarketDataOrderBookEventTypeCadenzaDropCopyPlaceOrderRequestAck  MarketDataOrderBookEventType = "cadenza.dropCopy.placeOrderRequestAck"
	MarketDataOrderBookEventTypeCadenzaDropCopyCancelOrderRequestAck MarketDataOrderBookEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	MarketDataOrderBookEventTypeCadenzaDropCopyQuote                 MarketDataOrderBookEventType = "cadenza.dropCopy.quote"
	MarketDataOrderBookEventTypeCadenzaDropCopyOrder                 MarketDataOrderBookEventType = "cadenza.dropCopy.order"
	MarketDataOrderBookEventTypeCadenzaDropCopyExecutionReport       MarketDataOrderBookEventType = "cadenza.dropCopy.executionReport"
	MarketDataOrderBookEventTypeCadenzaDropCopyPortfolio             MarketDataOrderBookEventType = "cadenza.dropCopy.portfolio"
	MarketDataOrderBookEventTypeCadenzaMarketDataOrderBook           MarketDataOrderBookEventType = "cadenza.marketData.orderBook"
	MarketDataOrderBookEventTypeCadenzaMarketDataKline               MarketDataOrderBookEventType = "cadenza.marketData.kline"
)

func (r MarketDataOrderBookEventType) IsKnown() bool {
	switch r {
	case MarketDataOrderBookEventTypeCadenzaTaskQuote, MarketDataOrderBookEventTypeCadenzaDropCopyQuoteRequestAck, MarketDataOrderBookEventTypeCadenzaDropCopyPlaceOrderRequestAck, MarketDataOrderBookEventTypeCadenzaDropCopyCancelOrderRequestAck, MarketDataOrderBookEventTypeCadenzaDropCopyQuote, MarketDataOrderBookEventTypeCadenzaDropCopyOrder, MarketDataOrderBookEventTypeCadenzaDropCopyExecutionReport, MarketDataOrderBookEventTypeCadenzaDropCopyPortfolio, MarketDataOrderBookEventTypeCadenzaMarketDataOrderBook, MarketDataOrderBookEventTypeCadenzaMarketDataKline:
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
