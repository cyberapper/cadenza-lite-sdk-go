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
	DropCopyExecutionReportEventTypeCadenzaTaskQuoteRequestAck       DropCopyExecutionReportEventType = "cadenza.task.quoteRequestAck"
	DropCopyExecutionReportEventTypeCadenzaTaskPlaceOrderRequestAck  DropCopyExecutionReportEventType = "cadenza.task.placeOrderRequestAck"
	DropCopyExecutionReportEventTypeCadenzaTaskCancelOrderRequestAck DropCopyExecutionReportEventType = "cadenza.task.cancelOrderRequestAck"
	DropCopyExecutionReportEventTypeCadenzaDropCopyQuote             DropCopyExecutionReportEventType = "cadenza.dropCopy.quote"
	DropCopyExecutionReportEventTypeCadenzaDropCopyOrder             DropCopyExecutionReportEventType = "cadenza.dropCopy.order"
	DropCopyExecutionReportEventTypeCadenzaDropCopyPortfolio         DropCopyExecutionReportEventType = "cadenza.dropCopy.portfolio"
	DropCopyExecutionReportEventTypeCadenzaMarketDataOrderBook       DropCopyExecutionReportEventType = "cadenza.marketData.orderBook"
	DropCopyExecutionReportEventTypeCadenzaMarketDataKline           DropCopyExecutionReportEventType = "cadenza.marketData.kline"
)

func (r DropCopyExecutionReportEventType) IsKnown() bool {
	switch r {
	case DropCopyExecutionReportEventTypeCadenzaTaskQuoteRequestAck, DropCopyExecutionReportEventTypeCadenzaTaskPlaceOrderRequestAck, DropCopyExecutionReportEventTypeCadenzaTaskCancelOrderRequestAck, DropCopyExecutionReportEventTypeCadenzaDropCopyQuote, DropCopyExecutionReportEventTypeCadenzaDropCopyOrder, DropCopyExecutionReportEventTypeCadenzaDropCopyPortfolio, DropCopyExecutionReportEventTypeCadenzaMarketDataOrderBook, DropCopyExecutionReportEventTypeCadenzaMarketDataKline:
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
	DropCopyOrderEventTypeCadenzaTaskQuoteRequestAck       DropCopyOrderEventType = "cadenza.task.quoteRequestAck"
	DropCopyOrderEventTypeCadenzaTaskPlaceOrderRequestAck  DropCopyOrderEventType = "cadenza.task.placeOrderRequestAck"
	DropCopyOrderEventTypeCadenzaTaskCancelOrderRequestAck DropCopyOrderEventType = "cadenza.task.cancelOrderRequestAck"
	DropCopyOrderEventTypeCadenzaDropCopyQuote             DropCopyOrderEventType = "cadenza.dropCopy.quote"
	DropCopyOrderEventTypeCadenzaDropCopyOrder             DropCopyOrderEventType = "cadenza.dropCopy.order"
	DropCopyOrderEventTypeCadenzaDropCopyPortfolio         DropCopyOrderEventType = "cadenza.dropCopy.portfolio"
	DropCopyOrderEventTypeCadenzaMarketDataOrderBook       DropCopyOrderEventType = "cadenza.marketData.orderBook"
	DropCopyOrderEventTypeCadenzaMarketDataKline           DropCopyOrderEventType = "cadenza.marketData.kline"
)

func (r DropCopyOrderEventType) IsKnown() bool {
	switch r {
	case DropCopyOrderEventTypeCadenzaTaskQuoteRequestAck, DropCopyOrderEventTypeCadenzaTaskPlaceOrderRequestAck, DropCopyOrderEventTypeCadenzaTaskCancelOrderRequestAck, DropCopyOrderEventTypeCadenzaDropCopyQuote, DropCopyOrderEventTypeCadenzaDropCopyOrder, DropCopyOrderEventTypeCadenzaDropCopyPortfolio, DropCopyOrderEventTypeCadenzaMarketDataOrderBook, DropCopyOrderEventTypeCadenzaMarketDataKline:
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
	DropCopyPortfolioEventTypeCadenzaTaskQuoteRequestAck       DropCopyPortfolioEventType = "cadenza.task.quoteRequestAck"
	DropCopyPortfolioEventTypeCadenzaTaskPlaceOrderRequestAck  DropCopyPortfolioEventType = "cadenza.task.placeOrderRequestAck"
	DropCopyPortfolioEventTypeCadenzaTaskCancelOrderRequestAck DropCopyPortfolioEventType = "cadenza.task.cancelOrderRequestAck"
	DropCopyPortfolioEventTypeCadenzaDropCopyQuote             DropCopyPortfolioEventType = "cadenza.dropCopy.quote"
	DropCopyPortfolioEventTypeCadenzaDropCopyOrder             DropCopyPortfolioEventType = "cadenza.dropCopy.order"
	DropCopyPortfolioEventTypeCadenzaDropCopyPortfolio         DropCopyPortfolioEventType = "cadenza.dropCopy.portfolio"
	DropCopyPortfolioEventTypeCadenzaMarketDataOrderBook       DropCopyPortfolioEventType = "cadenza.marketData.orderBook"
	DropCopyPortfolioEventTypeCadenzaMarketDataKline           DropCopyPortfolioEventType = "cadenza.marketData.kline"
)

func (r DropCopyPortfolioEventType) IsKnown() bool {
	switch r {
	case DropCopyPortfolioEventTypeCadenzaTaskQuoteRequestAck, DropCopyPortfolioEventTypeCadenzaTaskPlaceOrderRequestAck, DropCopyPortfolioEventTypeCadenzaTaskCancelOrderRequestAck, DropCopyPortfolioEventTypeCadenzaDropCopyQuote, DropCopyPortfolioEventTypeCadenzaDropCopyOrder, DropCopyPortfolioEventTypeCadenzaDropCopyPortfolio, DropCopyPortfolioEventTypeCadenzaMarketDataOrderBook, DropCopyPortfolioEventTypeCadenzaMarketDataKline:
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
	DropCopyQuoteEventTypeCadenzaTaskQuoteRequestAck       DropCopyQuoteEventType = "cadenza.task.quoteRequestAck"
	DropCopyQuoteEventTypeCadenzaTaskPlaceOrderRequestAck  DropCopyQuoteEventType = "cadenza.task.placeOrderRequestAck"
	DropCopyQuoteEventTypeCadenzaTaskCancelOrderRequestAck DropCopyQuoteEventType = "cadenza.task.cancelOrderRequestAck"
	DropCopyQuoteEventTypeCadenzaDropCopyQuote             DropCopyQuoteEventType = "cadenza.dropCopy.quote"
	DropCopyQuoteEventTypeCadenzaDropCopyOrder             DropCopyQuoteEventType = "cadenza.dropCopy.order"
	DropCopyQuoteEventTypeCadenzaDropCopyPortfolio         DropCopyQuoteEventType = "cadenza.dropCopy.portfolio"
	DropCopyQuoteEventTypeCadenzaMarketDataOrderBook       DropCopyQuoteEventType = "cadenza.marketData.orderBook"
	DropCopyQuoteEventTypeCadenzaMarketDataKline           DropCopyQuoteEventType = "cadenza.marketData.kline"
)

func (r DropCopyQuoteEventType) IsKnown() bool {
	switch r {
	case DropCopyQuoteEventTypeCadenzaTaskQuoteRequestAck, DropCopyQuoteEventTypeCadenzaTaskPlaceOrderRequestAck, DropCopyQuoteEventTypeCadenzaTaskCancelOrderRequestAck, DropCopyQuoteEventTypeCadenzaDropCopyQuote, DropCopyQuoteEventTypeCadenzaDropCopyOrder, DropCopyQuoteEventTypeCadenzaDropCopyPortfolio, DropCopyQuoteEventTypeCadenzaMarketDataOrderBook, DropCopyQuoteEventTypeCadenzaMarketDataKline:
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
