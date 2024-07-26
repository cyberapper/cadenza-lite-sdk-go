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

// EventDropCopyService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventDropCopyService] method instead.
type EventDropCopyService struct {
	Options []option.RequestOption
}

// NewEventDropCopyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventDropCopyService(opts ...option.RequestOption) (r *EventDropCopyService) {
	r = &EventDropCopyService{}
	r.Options = opts
	return
}

// PubSub event handler placeholder for cancel order request acknowledgment event
func (r *EventDropCopyService) DropCopyCancelOrderRequestAck(ctx context.Context, body EventDropCopyDropCopyCancelOrderRequestAckParams, opts ...option.RequestOption) (res *DropCopyCancelOrderRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/cancelOrderRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler for execution report drop copy event
func (r *EventDropCopyService) DropCopyExecutionReport(ctx context.Context, body EventDropCopyDropCopyExecutionReportParams, opts ...option.RequestOption) (res *DropCopyExecutionReport, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/executionReport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for order event
func (r *EventDropCopyService) DropCopyOrder(ctx context.Context, body EventDropCopyDropCopyOrderParams, opts ...option.RequestOption) (res *DropCopyOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for place order request acknowledgment event
func (r *EventDropCopyService) DropCopyPlaceOrderRequestAck(ctx context.Context, body EventDropCopyDropCopyPlaceOrderRequestAckParams, opts ...option.RequestOption) (res *DropCopyPlaceOrderRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/placeOrderRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for portfolio event
func (r *EventDropCopyService) DropCopyPortfolio(ctx context.Context, body EventDropCopyDropCopyPortfolioParams, opts ...option.RequestOption) (res *DropCopyPortfolio, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/portfolio"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for quote event
func (r *EventDropCopyService) DropCopyQuote(ctx context.Context, body EventDropCopyDropCopyQuoteParams, opts ...option.RequestOption) (res *DropCopyQuote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for quote request acknowledgment event
func (r *EventDropCopyService) DropCopyQuoteRequestAck(ctx context.Context, body EventDropCopyDropCopyQuoteRequestAckParams, opts ...option.RequestOption) (res *DropCopyRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/quoteRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DropCopyCancelOrderRequestAck struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType DropCopyCancelOrderRequestAckEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64              `json:"timestamp,required"`
	Payload   CancelOrderRequest `json:"payload"`
	// The source system or module that generated the event.
	Source string                            `json:"source"`
	JSON   dropCopyCancelOrderRequestAckJSON `json:"-"`
}

// dropCopyCancelOrderRequestAckJSON contains the JSON metadata for the struct
// [DropCopyCancelOrderRequestAck]
type dropCopyCancelOrderRequestAckJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyCancelOrderRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyCancelOrderRequestAckJSON) RawJSON() string {
	return r.raw
}

// Event Type
type DropCopyCancelOrderRequestAckEventType string

const (
	DropCopyCancelOrderRequestAckEventTypeCadenzaTaskQuote                     DropCopyCancelOrderRequestAckEventType = "cadenza.task.quote"
	DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyQuoteRequestAck       DropCopyCancelOrderRequestAckEventType = "cadenza.dropCopy.quoteRequestAck"
	DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyPlaceOrderRequestAck  DropCopyCancelOrderRequestAckEventType = "cadenza.dropCopy.placeOrderRequestAck"
	DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyCancelOrderRequestAck DropCopyCancelOrderRequestAckEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyQuote                 DropCopyCancelOrderRequestAckEventType = "cadenza.dropCopy.quote"
	DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyOrder                 DropCopyCancelOrderRequestAckEventType = "cadenza.dropCopy.order"
	DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyExecutionReport       DropCopyCancelOrderRequestAckEventType = "cadenza.dropCopy.executionReport"
	DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyPortfolio             DropCopyCancelOrderRequestAckEventType = "cadenza.dropCopy.portfolio"
	DropCopyCancelOrderRequestAckEventTypeCadenzaMarketDataOrderBook           DropCopyCancelOrderRequestAckEventType = "cadenza.marketData.orderBook"
	DropCopyCancelOrderRequestAckEventTypeCadenzaMarketDataKline               DropCopyCancelOrderRequestAckEventType = "cadenza.marketData.kline"
)

func (r DropCopyCancelOrderRequestAckEventType) IsKnown() bool {
	switch r {
	case DropCopyCancelOrderRequestAckEventTypeCadenzaTaskQuote, DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyQuoteRequestAck, DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyPlaceOrderRequestAck, DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyCancelOrderRequestAck, DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyQuote, DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyOrder, DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyExecutionReport, DropCopyCancelOrderRequestAckEventTypeCadenzaDropCopyPortfolio, DropCopyCancelOrderRequestAckEventTypeCadenzaMarketDataOrderBook, DropCopyCancelOrderRequestAckEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type DropCopyCancelOrderRequestAckParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[DropCopyCancelOrderRequestAckEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                   `json:"timestamp,required"`
	Payload   param.Field[CancelOrderRequestParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r DropCopyCancelOrderRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyExecutionReport struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType DropCopyExecutionReportEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64           `json:"timestamp,required"`
	Payload   ExecutionReport `json:"payload"`
	// The source system or module that generated the event.
	Source string                      `json:"source"`
	JSON   dropCopyExecutionReportJSON `json:"-"`
}

// dropCopyExecutionReportJSON contains the JSON metadata for the struct
// [DropCopyExecutionReport]
type dropCopyExecutionReportJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyExecutionReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyExecutionReportJSON) RawJSON() string {
	return r.raw
}

// Event Type
type DropCopyExecutionReportEventType string

const (
	DropCopyExecutionReportEventTypeCadenzaTaskQuote                     DropCopyExecutionReportEventType = "cadenza.task.quote"
	DropCopyExecutionReportEventTypeCadenzaDropCopyQuoteRequestAck       DropCopyExecutionReportEventType = "cadenza.dropCopy.quoteRequestAck"
	DropCopyExecutionReportEventTypeCadenzaDropCopyPlaceOrderRequestAck  DropCopyExecutionReportEventType = "cadenza.dropCopy.placeOrderRequestAck"
	DropCopyExecutionReportEventTypeCadenzaDropCopyCancelOrderRequestAck DropCopyExecutionReportEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	DropCopyExecutionReportEventTypeCadenzaDropCopyQuote                 DropCopyExecutionReportEventType = "cadenza.dropCopy.quote"
	DropCopyExecutionReportEventTypeCadenzaDropCopyOrder                 DropCopyExecutionReportEventType = "cadenza.dropCopy.order"
	DropCopyExecutionReportEventTypeCadenzaDropCopyExecutionReport       DropCopyExecutionReportEventType = "cadenza.dropCopy.executionReport"
	DropCopyExecutionReportEventTypeCadenzaDropCopyPortfolio             DropCopyExecutionReportEventType = "cadenza.dropCopy.portfolio"
	DropCopyExecutionReportEventTypeCadenzaMarketDataOrderBook           DropCopyExecutionReportEventType = "cadenza.marketData.orderBook"
	DropCopyExecutionReportEventTypeCadenzaMarketDataKline               DropCopyExecutionReportEventType = "cadenza.marketData.kline"
)

func (r DropCopyExecutionReportEventType) IsKnown() bool {
	switch r {
	case DropCopyExecutionReportEventTypeCadenzaTaskQuote, DropCopyExecutionReportEventTypeCadenzaDropCopyQuoteRequestAck, DropCopyExecutionReportEventTypeCadenzaDropCopyPlaceOrderRequestAck, DropCopyExecutionReportEventTypeCadenzaDropCopyCancelOrderRequestAck, DropCopyExecutionReportEventTypeCadenzaDropCopyQuote, DropCopyExecutionReportEventTypeCadenzaDropCopyOrder, DropCopyExecutionReportEventTypeCadenzaDropCopyExecutionReport, DropCopyExecutionReportEventTypeCadenzaDropCopyPortfolio, DropCopyExecutionReportEventTypeCadenzaMarketDataOrderBook, DropCopyExecutionReportEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type DropCopyExecutionReportParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[DropCopyExecutionReportEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                `json:"timestamp,required"`
	Payload   param.Field[ExecutionReportParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r DropCopyExecutionReportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyOrder struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType DropCopyOrderEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64 `json:"timestamp,required"`
	Payload   Order `json:"payload"`
	// The source system or module that generated the event.
	Source string            `json:"source"`
	JSON   dropCopyOrderJSON `json:"-"`
}

// dropCopyOrderJSON contains the JSON metadata for the struct [DropCopyOrder]
type dropCopyOrderJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyOrderJSON) RawJSON() string {
	return r.raw
}

// Event Type
type DropCopyOrderEventType string

const (
	DropCopyOrderEventTypeCadenzaTaskQuote                     DropCopyOrderEventType = "cadenza.task.quote"
	DropCopyOrderEventTypeCadenzaDropCopyQuoteRequestAck       DropCopyOrderEventType = "cadenza.dropCopy.quoteRequestAck"
	DropCopyOrderEventTypeCadenzaDropCopyPlaceOrderRequestAck  DropCopyOrderEventType = "cadenza.dropCopy.placeOrderRequestAck"
	DropCopyOrderEventTypeCadenzaDropCopyCancelOrderRequestAck DropCopyOrderEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	DropCopyOrderEventTypeCadenzaDropCopyQuote                 DropCopyOrderEventType = "cadenza.dropCopy.quote"
	DropCopyOrderEventTypeCadenzaDropCopyOrder                 DropCopyOrderEventType = "cadenza.dropCopy.order"
	DropCopyOrderEventTypeCadenzaDropCopyExecutionReport       DropCopyOrderEventType = "cadenza.dropCopy.executionReport"
	DropCopyOrderEventTypeCadenzaDropCopyPortfolio             DropCopyOrderEventType = "cadenza.dropCopy.portfolio"
	DropCopyOrderEventTypeCadenzaMarketDataOrderBook           DropCopyOrderEventType = "cadenza.marketData.orderBook"
	DropCopyOrderEventTypeCadenzaMarketDataKline               DropCopyOrderEventType = "cadenza.marketData.kline"
)

func (r DropCopyOrderEventType) IsKnown() bool {
	switch r {
	case DropCopyOrderEventTypeCadenzaTaskQuote, DropCopyOrderEventTypeCadenzaDropCopyQuoteRequestAck, DropCopyOrderEventTypeCadenzaDropCopyPlaceOrderRequestAck, DropCopyOrderEventTypeCadenzaDropCopyCancelOrderRequestAck, DropCopyOrderEventTypeCadenzaDropCopyQuote, DropCopyOrderEventTypeCadenzaDropCopyOrder, DropCopyOrderEventTypeCadenzaDropCopyExecutionReport, DropCopyOrderEventTypeCadenzaDropCopyPortfolio, DropCopyOrderEventTypeCadenzaMarketDataOrderBook, DropCopyOrderEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type DropCopyOrderParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[DropCopyOrderEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]      `json:"timestamp,required"`
	Payload   param.Field[OrderParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r DropCopyOrderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyPlaceOrderRequestAck struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType DropCopyPlaceOrderRequestAckEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64             `json:"timestamp,required"`
	Payload   PlaceOrderRequest `json:"payload"`
	// The source system or module that generated the event.
	Source string                           `json:"source"`
	JSON   dropCopyPlaceOrderRequestAckJSON `json:"-"`
}

// dropCopyPlaceOrderRequestAckJSON contains the JSON metadata for the struct
// [DropCopyPlaceOrderRequestAck]
type dropCopyPlaceOrderRequestAckJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyPlaceOrderRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyPlaceOrderRequestAckJSON) RawJSON() string {
	return r.raw
}

// Event Type
type DropCopyPlaceOrderRequestAckEventType string

const (
	DropCopyPlaceOrderRequestAckEventTypeCadenzaTaskQuote                     DropCopyPlaceOrderRequestAckEventType = "cadenza.task.quote"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyQuoteRequestAck       DropCopyPlaceOrderRequestAckEventType = "cadenza.dropCopy.quoteRequestAck"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyPlaceOrderRequestAck  DropCopyPlaceOrderRequestAckEventType = "cadenza.dropCopy.placeOrderRequestAck"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyCancelOrderRequestAck DropCopyPlaceOrderRequestAckEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyQuote                 DropCopyPlaceOrderRequestAckEventType = "cadenza.dropCopy.quote"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyOrder                 DropCopyPlaceOrderRequestAckEventType = "cadenza.dropCopy.order"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyExecutionReport       DropCopyPlaceOrderRequestAckEventType = "cadenza.dropCopy.executionReport"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyPortfolio             DropCopyPlaceOrderRequestAckEventType = "cadenza.dropCopy.portfolio"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaMarketDataOrderBook           DropCopyPlaceOrderRequestAckEventType = "cadenza.marketData.orderBook"
	DropCopyPlaceOrderRequestAckEventTypeCadenzaMarketDataKline               DropCopyPlaceOrderRequestAckEventType = "cadenza.marketData.kline"
)

func (r DropCopyPlaceOrderRequestAckEventType) IsKnown() bool {
	switch r {
	case DropCopyPlaceOrderRequestAckEventTypeCadenzaTaskQuote, DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyQuoteRequestAck, DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyPlaceOrderRequestAck, DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyCancelOrderRequestAck, DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyQuote, DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyOrder, DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyExecutionReport, DropCopyPlaceOrderRequestAckEventTypeCadenzaDropCopyPortfolio, DropCopyPlaceOrderRequestAckEventTypeCadenzaMarketDataOrderBook, DropCopyPlaceOrderRequestAckEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type DropCopyPlaceOrderRequestAckParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[DropCopyPlaceOrderRequestAckEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                  `json:"timestamp,required"`
	Payload   param.Field[PlaceOrderRequestParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r DropCopyPlaceOrderRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyPortfolio struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType DropCopyPortfolioEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64                    `json:"timestamp,required"`
	Payload   ExchangeAccountPortfolio `json:"payload"`
	// The source system or module that generated the event.
	Source string                `json:"source"`
	JSON   dropCopyPortfolioJSON `json:"-"`
}

// dropCopyPortfolioJSON contains the JSON metadata for the struct
// [DropCopyPortfolio]
type dropCopyPortfolioJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyPortfolio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyPortfolioJSON) RawJSON() string {
	return r.raw
}

// Event Type
type DropCopyPortfolioEventType string

const (
	DropCopyPortfolioEventTypeCadenzaTaskQuote                     DropCopyPortfolioEventType = "cadenza.task.quote"
	DropCopyPortfolioEventTypeCadenzaDropCopyQuoteRequestAck       DropCopyPortfolioEventType = "cadenza.dropCopy.quoteRequestAck"
	DropCopyPortfolioEventTypeCadenzaDropCopyPlaceOrderRequestAck  DropCopyPortfolioEventType = "cadenza.dropCopy.placeOrderRequestAck"
	DropCopyPortfolioEventTypeCadenzaDropCopyCancelOrderRequestAck DropCopyPortfolioEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	DropCopyPortfolioEventTypeCadenzaDropCopyQuote                 DropCopyPortfolioEventType = "cadenza.dropCopy.quote"
	DropCopyPortfolioEventTypeCadenzaDropCopyOrder                 DropCopyPortfolioEventType = "cadenza.dropCopy.order"
	DropCopyPortfolioEventTypeCadenzaDropCopyExecutionReport       DropCopyPortfolioEventType = "cadenza.dropCopy.executionReport"
	DropCopyPortfolioEventTypeCadenzaDropCopyPortfolio             DropCopyPortfolioEventType = "cadenza.dropCopy.portfolio"
	DropCopyPortfolioEventTypeCadenzaMarketDataOrderBook           DropCopyPortfolioEventType = "cadenza.marketData.orderBook"
	DropCopyPortfolioEventTypeCadenzaMarketDataKline               DropCopyPortfolioEventType = "cadenza.marketData.kline"
)

func (r DropCopyPortfolioEventType) IsKnown() bool {
	switch r {
	case DropCopyPortfolioEventTypeCadenzaTaskQuote, DropCopyPortfolioEventTypeCadenzaDropCopyQuoteRequestAck, DropCopyPortfolioEventTypeCadenzaDropCopyPlaceOrderRequestAck, DropCopyPortfolioEventTypeCadenzaDropCopyCancelOrderRequestAck, DropCopyPortfolioEventTypeCadenzaDropCopyQuote, DropCopyPortfolioEventTypeCadenzaDropCopyOrder, DropCopyPortfolioEventTypeCadenzaDropCopyExecutionReport, DropCopyPortfolioEventTypeCadenzaDropCopyPortfolio, DropCopyPortfolioEventTypeCadenzaMarketDataOrderBook, DropCopyPortfolioEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type DropCopyPortfolioParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[DropCopyPortfolioEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                         `json:"timestamp,required"`
	Payload   param.Field[ExchangeAccountPortfolioParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r DropCopyPortfolioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyQuote struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType DropCopyQuoteEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64 `json:"timestamp,required"`
	Payload   Quote `json:"payload"`
	// The source system or module that generated the event.
	Source string            `json:"source"`
	JSON   dropCopyQuoteJSON `json:"-"`
}

// dropCopyQuoteJSON contains the JSON metadata for the struct [DropCopyQuote]
type dropCopyQuoteJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyQuote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyQuoteJSON) RawJSON() string {
	return r.raw
}

// Event Type
type DropCopyQuoteEventType string

const (
	DropCopyQuoteEventTypeCadenzaTaskQuote                     DropCopyQuoteEventType = "cadenza.task.quote"
	DropCopyQuoteEventTypeCadenzaDropCopyQuoteRequestAck       DropCopyQuoteEventType = "cadenza.dropCopy.quoteRequestAck"
	DropCopyQuoteEventTypeCadenzaDropCopyPlaceOrderRequestAck  DropCopyQuoteEventType = "cadenza.dropCopy.placeOrderRequestAck"
	DropCopyQuoteEventTypeCadenzaDropCopyCancelOrderRequestAck DropCopyQuoteEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	DropCopyQuoteEventTypeCadenzaDropCopyQuote                 DropCopyQuoteEventType = "cadenza.dropCopy.quote"
	DropCopyQuoteEventTypeCadenzaDropCopyOrder                 DropCopyQuoteEventType = "cadenza.dropCopy.order"
	DropCopyQuoteEventTypeCadenzaDropCopyExecutionReport       DropCopyQuoteEventType = "cadenza.dropCopy.executionReport"
	DropCopyQuoteEventTypeCadenzaDropCopyPortfolio             DropCopyQuoteEventType = "cadenza.dropCopy.portfolio"
	DropCopyQuoteEventTypeCadenzaMarketDataOrderBook           DropCopyQuoteEventType = "cadenza.marketData.orderBook"
	DropCopyQuoteEventTypeCadenzaMarketDataKline               DropCopyQuoteEventType = "cadenza.marketData.kline"
)

func (r DropCopyQuoteEventType) IsKnown() bool {
	switch r {
	case DropCopyQuoteEventTypeCadenzaTaskQuote, DropCopyQuoteEventTypeCadenzaDropCopyQuoteRequestAck, DropCopyQuoteEventTypeCadenzaDropCopyPlaceOrderRequestAck, DropCopyQuoteEventTypeCadenzaDropCopyCancelOrderRequestAck, DropCopyQuoteEventTypeCadenzaDropCopyQuote, DropCopyQuoteEventTypeCadenzaDropCopyOrder, DropCopyQuoteEventTypeCadenzaDropCopyExecutionReport, DropCopyQuoteEventTypeCadenzaDropCopyPortfolio, DropCopyQuoteEventTypeCadenzaMarketDataOrderBook, DropCopyQuoteEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type DropCopyQuoteParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[DropCopyQuoteEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]      `json:"timestamp,required"`
	Payload   param.Field[QuoteParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r DropCopyQuoteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyRequestAck struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType DropCopyRequestAckEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64        `json:"timestamp,required"`
	Payload   QuoteRequest `json:"payload"`
	// The source system or module that generated the event.
	Source string                 `json:"source"`
	JSON   dropCopyRequestAckJSON `json:"-"`
}

// dropCopyRequestAckJSON contains the JSON metadata for the struct
// [DropCopyRequestAck]
type dropCopyRequestAckJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyRequestAckJSON) RawJSON() string {
	return r.raw
}

// Event Type
type DropCopyRequestAckEventType string

const (
	DropCopyRequestAckEventTypeCadenzaTaskQuote                     DropCopyRequestAckEventType = "cadenza.task.quote"
	DropCopyRequestAckEventTypeCadenzaDropCopyQuoteRequestAck       DropCopyRequestAckEventType = "cadenza.dropCopy.quoteRequestAck"
	DropCopyRequestAckEventTypeCadenzaDropCopyPlaceOrderRequestAck  DropCopyRequestAckEventType = "cadenza.dropCopy.placeOrderRequestAck"
	DropCopyRequestAckEventTypeCadenzaDropCopyCancelOrderRequestAck DropCopyRequestAckEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	DropCopyRequestAckEventTypeCadenzaDropCopyQuote                 DropCopyRequestAckEventType = "cadenza.dropCopy.quote"
	DropCopyRequestAckEventTypeCadenzaDropCopyOrder                 DropCopyRequestAckEventType = "cadenza.dropCopy.order"
	DropCopyRequestAckEventTypeCadenzaDropCopyExecutionReport       DropCopyRequestAckEventType = "cadenza.dropCopy.executionReport"
	DropCopyRequestAckEventTypeCadenzaDropCopyPortfolio             DropCopyRequestAckEventType = "cadenza.dropCopy.portfolio"
	DropCopyRequestAckEventTypeCadenzaMarketDataOrderBook           DropCopyRequestAckEventType = "cadenza.marketData.orderBook"
	DropCopyRequestAckEventTypeCadenzaMarketDataKline               DropCopyRequestAckEventType = "cadenza.marketData.kline"
)

func (r DropCopyRequestAckEventType) IsKnown() bool {
	switch r {
	case DropCopyRequestAckEventTypeCadenzaTaskQuote, DropCopyRequestAckEventTypeCadenzaDropCopyQuoteRequestAck, DropCopyRequestAckEventTypeCadenzaDropCopyPlaceOrderRequestAck, DropCopyRequestAckEventTypeCadenzaDropCopyCancelOrderRequestAck, DropCopyRequestAckEventTypeCadenzaDropCopyQuote, DropCopyRequestAckEventTypeCadenzaDropCopyOrder, DropCopyRequestAckEventTypeCadenzaDropCopyExecutionReport, DropCopyRequestAckEventTypeCadenzaDropCopyPortfolio, DropCopyRequestAckEventTypeCadenzaMarketDataOrderBook, DropCopyRequestAckEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type DropCopyRequestAckParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[DropCopyRequestAckEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]             `json:"timestamp,required"`
	Payload   param.Field[QuoteRequestParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r DropCopyRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventDropCopyDropCopyCancelOrderRequestAckParams struct {
	DropCopyCancelOrderRequestAck DropCopyCancelOrderRequestAckParam `json:"dropCopyCancelOrderRequestAck,required"`
}

func (r EventDropCopyDropCopyCancelOrderRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyCancelOrderRequestAck)
}

type EventDropCopyDropCopyExecutionReportParams struct {
	DropCopyExecutionReport DropCopyExecutionReportParam `json:"dropCopyExecutionReport,required"`
}

func (r EventDropCopyDropCopyExecutionReportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyExecutionReport)
}

type EventDropCopyDropCopyOrderParams struct {
	DropCopyOrder DropCopyOrderParam `json:"dropCopyOrder,required"`
}

func (r EventDropCopyDropCopyOrderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyOrder)
}

type EventDropCopyDropCopyPlaceOrderRequestAckParams struct {
	DropCopyPlaceOrderRequestAck DropCopyPlaceOrderRequestAckParam `json:"dropCopyPlaceOrderRequestAck,required"`
}

func (r EventDropCopyDropCopyPlaceOrderRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyPlaceOrderRequestAck)
}

type EventDropCopyDropCopyPortfolioParams struct {
	DropCopyPortfolio DropCopyPortfolioParam `json:"dropCopyPortfolio,required"`
}

func (r EventDropCopyDropCopyPortfolioParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyPortfolio)
}

type EventDropCopyDropCopyQuoteParams struct {
	DropCopyQuote DropCopyQuoteParam `json:"dropCopyQuote,required"`
}

func (r EventDropCopyDropCopyQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyQuote)
}

type EventDropCopyDropCopyQuoteRequestAckParams struct {
	DropCopyRequestAck DropCopyRequestAckParam `json:"dropCopyRequestAck,required"`
}

func (r EventDropCopyDropCopyQuoteRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyRequestAck)
}
