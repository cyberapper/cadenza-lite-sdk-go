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

// EventTaskService contains methods and other services that help with interacting
// with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventTaskService] method instead.
type EventTaskService struct {
	Options []option.RequestOption
}

// NewEventTaskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEventTaskService(opts ...option.RequestOption) (r *EventTaskService) {
	r = &EventTaskService{}
	r.Options = opts
	return
}

// PubSub event handler placeholder for quote request acknowledgment event
func (r *EventTaskService) TaskQuote(ctx context.Context, body EventTaskTaskQuoteParams, opts ...option.RequestOption) (res *TaskQuote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/task/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TaskQuote struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType TaskQuoteEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64        `json:"timestamp,required"`
	Payload   QuoteRequest `json:"payload"`
	// The source system or module that generated the event.
	Source string        `json:"source"`
	JSON   taskQuoteJSON `json:"-"`
}

// taskQuoteJSON contains the JSON metadata for the struct [TaskQuote]
type taskQuoteJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskQuote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskQuoteJSON) RawJSON() string {
	return r.raw
}

// Event Type
type TaskQuoteEventType string

const (
	TaskQuoteEventTypeCadenzaTaskQuote                     TaskQuoteEventType = "cadenza.task.quote"
	TaskQuoteEventTypeCadenzaDropCopyQuoteRequestAck       TaskQuoteEventType = "cadenza.dropCopy.quoteRequestAck"
	TaskQuoteEventTypeCadenzaDropCopyPlaceOrderRequestAck  TaskQuoteEventType = "cadenza.dropCopy.placeOrderRequestAck"
	TaskQuoteEventTypeCadenzaDropCopyCancelOrderRequestAck TaskQuoteEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	TaskQuoteEventTypeCadenzaDropCopyQuote                 TaskQuoteEventType = "cadenza.dropCopy.quote"
	TaskQuoteEventTypeCadenzaDropCopyOrder                 TaskQuoteEventType = "cadenza.dropCopy.order"
	TaskQuoteEventTypeCadenzaDropCopyExecutionReport       TaskQuoteEventType = "cadenza.dropCopy.executionReport"
	TaskQuoteEventTypeCadenzaDropCopyPortfolio             TaskQuoteEventType = "cadenza.dropCopy.portfolio"
	TaskQuoteEventTypeCadenzaMarketDataOrderBook           TaskQuoteEventType = "cadenza.marketData.orderBook"
	TaskQuoteEventTypeCadenzaMarketDataKline               TaskQuoteEventType = "cadenza.marketData.kline"
)

func (r TaskQuoteEventType) IsKnown() bool {
	switch r {
	case TaskQuoteEventTypeCadenzaTaskQuote, TaskQuoteEventTypeCadenzaDropCopyQuoteRequestAck, TaskQuoteEventTypeCadenzaDropCopyPlaceOrderRequestAck, TaskQuoteEventTypeCadenzaDropCopyCancelOrderRequestAck, TaskQuoteEventTypeCadenzaDropCopyQuote, TaskQuoteEventTypeCadenzaDropCopyOrder, TaskQuoteEventTypeCadenzaDropCopyExecutionReport, TaskQuoteEventTypeCadenzaDropCopyPortfolio, TaskQuoteEventTypeCadenzaMarketDataOrderBook, TaskQuoteEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type TaskQuoteParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[TaskQuoteEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]             `json:"timestamp,required"`
	Payload   param.Field[QuoteRequestParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r TaskQuoteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventTaskTaskQuoteParams struct {
	TaskQuote TaskQuoteParam `json:"taskQuote,required"`
}

func (r EventTaskTaskQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TaskQuote)
}
