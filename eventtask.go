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

// PubSub event handler placeholder for cancel order request acknowledgment event
func (r *EventTaskService) TaskCancelOrderRequestAck(ctx context.Context, body EventTaskTaskCancelOrderRequestAckParams, opts ...option.RequestOption) (res *TaskCancelOrderRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/task/cancelOrderRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for place order request acknowledgment event
func (r *EventTaskService) TaskPlaceOrderRequestAck(ctx context.Context, body EventTaskTaskPlaceOrderRequestAckParams, opts ...option.RequestOption) (res *TaskPlaceOrderRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/task/placeOrderRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for quote request acknowledgment event
func (r *EventTaskService) TaskQuoteRequestAck(ctx context.Context, body EventTaskTaskQuoteRequestAckParams, opts ...option.RequestOption) (res *TaskQuoteRequestAck, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/task/quoteRequestAck"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TaskCancelOrderRequestAck struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType TaskCancelOrderRequestAckEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64              `json:"timestamp,required"`
	Payload   CancelOrderRequest `json:"payload"`
	// The source system or module that generated the event.
	Source string                        `json:"source"`
	JSON   taskCancelOrderRequestAckJSON `json:"-"`
}

// taskCancelOrderRequestAckJSON contains the JSON metadata for the struct
// [TaskCancelOrderRequestAck]
type taskCancelOrderRequestAckJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskCancelOrderRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskCancelOrderRequestAckJSON) RawJSON() string {
	return r.raw
}

// Event Type
type TaskCancelOrderRequestAckEventType string

const (
	TaskCancelOrderRequestAckEventTypeCadenzaTaskQuoteRequestAck       TaskCancelOrderRequestAckEventType = "cadenza.task.quoteRequestAck"
	TaskCancelOrderRequestAckEventTypeCadenzaTaskPlaceOrderRequestAck  TaskCancelOrderRequestAckEventType = "cadenza.task.placeOrderRequestAck"
	TaskCancelOrderRequestAckEventTypeCadenzaTaskCancelOrderRequestAck TaskCancelOrderRequestAckEventType = "cadenza.task.cancelOrderRequestAck"
	TaskCancelOrderRequestAckEventTypeCadenzaDropCopyQuote             TaskCancelOrderRequestAckEventType = "cadenza.dropCopy.quote"
	TaskCancelOrderRequestAckEventTypeCadenzaDropCopyOrder             TaskCancelOrderRequestAckEventType = "cadenza.dropCopy.order"
	TaskCancelOrderRequestAckEventTypeCadenzaDropCopyPortfolio         TaskCancelOrderRequestAckEventType = "cadenza.dropCopy.portfolio"
	TaskCancelOrderRequestAckEventTypeCadenzaMarketDataOrderBook       TaskCancelOrderRequestAckEventType = "cadenza.marketData.orderBook"
	TaskCancelOrderRequestAckEventTypeCadenzaMarketDataKline           TaskCancelOrderRequestAckEventType = "cadenza.marketData.kline"
)

func (r TaskCancelOrderRequestAckEventType) IsKnown() bool {
	switch r {
	case TaskCancelOrderRequestAckEventTypeCadenzaTaskQuoteRequestAck, TaskCancelOrderRequestAckEventTypeCadenzaTaskPlaceOrderRequestAck, TaskCancelOrderRequestAckEventTypeCadenzaTaskCancelOrderRequestAck, TaskCancelOrderRequestAckEventTypeCadenzaDropCopyQuote, TaskCancelOrderRequestAckEventTypeCadenzaDropCopyOrder, TaskCancelOrderRequestAckEventTypeCadenzaDropCopyPortfolio, TaskCancelOrderRequestAckEventTypeCadenzaMarketDataOrderBook, TaskCancelOrderRequestAckEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type TaskCancelOrderRequestAckParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[TaskCancelOrderRequestAckEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                   `json:"timestamp,required"`
	Payload   param.Field[CancelOrderRequestParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r TaskCancelOrderRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskPlaceOrderRequestAck struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType TaskPlaceOrderRequestAckEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64             `json:"timestamp,required"`
	Payload   PlaceOrderRequest `json:"payload"`
	// The source system or module that generated the event.
	Source string                       `json:"source"`
	JSON   taskPlaceOrderRequestAckJSON `json:"-"`
}

// taskPlaceOrderRequestAckJSON contains the JSON metadata for the struct
// [TaskPlaceOrderRequestAck]
type taskPlaceOrderRequestAckJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskPlaceOrderRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskPlaceOrderRequestAckJSON) RawJSON() string {
	return r.raw
}

// Event Type
type TaskPlaceOrderRequestAckEventType string

const (
	TaskPlaceOrderRequestAckEventTypeCadenzaTaskQuoteRequestAck       TaskPlaceOrderRequestAckEventType = "cadenza.task.quoteRequestAck"
	TaskPlaceOrderRequestAckEventTypeCadenzaTaskPlaceOrderRequestAck  TaskPlaceOrderRequestAckEventType = "cadenza.task.placeOrderRequestAck"
	TaskPlaceOrderRequestAckEventTypeCadenzaTaskCancelOrderRequestAck TaskPlaceOrderRequestAckEventType = "cadenza.task.cancelOrderRequestAck"
	TaskPlaceOrderRequestAckEventTypeCadenzaDropCopyQuote             TaskPlaceOrderRequestAckEventType = "cadenza.dropCopy.quote"
	TaskPlaceOrderRequestAckEventTypeCadenzaDropCopyOrder             TaskPlaceOrderRequestAckEventType = "cadenza.dropCopy.order"
	TaskPlaceOrderRequestAckEventTypeCadenzaDropCopyPortfolio         TaskPlaceOrderRequestAckEventType = "cadenza.dropCopy.portfolio"
	TaskPlaceOrderRequestAckEventTypeCadenzaMarketDataOrderBook       TaskPlaceOrderRequestAckEventType = "cadenza.marketData.orderBook"
	TaskPlaceOrderRequestAckEventTypeCadenzaMarketDataKline           TaskPlaceOrderRequestAckEventType = "cadenza.marketData.kline"
)

func (r TaskPlaceOrderRequestAckEventType) IsKnown() bool {
	switch r {
	case TaskPlaceOrderRequestAckEventTypeCadenzaTaskQuoteRequestAck, TaskPlaceOrderRequestAckEventTypeCadenzaTaskPlaceOrderRequestAck, TaskPlaceOrderRequestAckEventTypeCadenzaTaskCancelOrderRequestAck, TaskPlaceOrderRequestAckEventTypeCadenzaDropCopyQuote, TaskPlaceOrderRequestAckEventTypeCadenzaDropCopyOrder, TaskPlaceOrderRequestAckEventTypeCadenzaDropCopyPortfolio, TaskPlaceOrderRequestAckEventTypeCadenzaMarketDataOrderBook, TaskPlaceOrderRequestAckEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type TaskPlaceOrderRequestAckParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[TaskPlaceOrderRequestAckEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]                  `json:"timestamp,required"`
	Payload   param.Field[PlaceOrderRequestParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r TaskPlaceOrderRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TaskQuoteRequestAck struct {
	// A unique identifier for the event.
	EventID string `json:"eventId,required"`
	// Event Type
	EventType TaskQuoteRequestAckEventType `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp int64        `json:"timestamp,required"`
	Payload   QuoteRequest `json:"payload"`
	// The source system or module that generated the event.
	Source string                  `json:"source"`
	JSON   taskQuoteRequestAckJSON `json:"-"`
}

// taskQuoteRequestAckJSON contains the JSON metadata for the struct
// [TaskQuoteRequestAck]
type taskQuoteRequestAckJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
	Payload     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TaskQuoteRequestAck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r taskQuoteRequestAckJSON) RawJSON() string {
	return r.raw
}

// Event Type
type TaskQuoteRequestAckEventType string

const (
	TaskQuoteRequestAckEventTypeCadenzaTaskQuoteRequestAck       TaskQuoteRequestAckEventType = "cadenza.task.quoteRequestAck"
	TaskQuoteRequestAckEventTypeCadenzaTaskPlaceOrderRequestAck  TaskQuoteRequestAckEventType = "cadenza.task.placeOrderRequestAck"
	TaskQuoteRequestAckEventTypeCadenzaTaskCancelOrderRequestAck TaskQuoteRequestAckEventType = "cadenza.task.cancelOrderRequestAck"
	TaskQuoteRequestAckEventTypeCadenzaDropCopyQuote             TaskQuoteRequestAckEventType = "cadenza.dropCopy.quote"
	TaskQuoteRequestAckEventTypeCadenzaDropCopyOrder             TaskQuoteRequestAckEventType = "cadenza.dropCopy.order"
	TaskQuoteRequestAckEventTypeCadenzaDropCopyPortfolio         TaskQuoteRequestAckEventType = "cadenza.dropCopy.portfolio"
	TaskQuoteRequestAckEventTypeCadenzaMarketDataOrderBook       TaskQuoteRequestAckEventType = "cadenza.marketData.orderBook"
	TaskQuoteRequestAckEventTypeCadenzaMarketDataKline           TaskQuoteRequestAckEventType = "cadenza.marketData.kline"
)

func (r TaskQuoteRequestAckEventType) IsKnown() bool {
	switch r {
	case TaskQuoteRequestAckEventTypeCadenzaTaskQuoteRequestAck, TaskQuoteRequestAckEventTypeCadenzaTaskPlaceOrderRequestAck, TaskQuoteRequestAckEventTypeCadenzaTaskCancelOrderRequestAck, TaskQuoteRequestAckEventTypeCadenzaDropCopyQuote, TaskQuoteRequestAckEventTypeCadenzaDropCopyOrder, TaskQuoteRequestAckEventTypeCadenzaDropCopyPortfolio, TaskQuoteRequestAckEventTypeCadenzaMarketDataOrderBook, TaskQuoteRequestAckEventTypeCadenzaMarketDataKline:
		return true
	}
	return false
}

type TaskQuoteRequestAckParam struct {
	// A unique identifier for the event.
	EventID param.Field[string] `json:"eventId,required"`
	// Event Type
	EventType param.Field[TaskQuoteRequestAckEventType] `json:"eventType,required"`
	// Unix timestamp in milliseconds when the event was generated.
	Timestamp param.Field[int64]             `json:"timestamp,required"`
	Payload   param.Field[QuoteRequestParam] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r TaskQuoteRequestAckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventTaskTaskCancelOrderRequestAckParams struct {
	TaskCancelOrderRequestAck TaskCancelOrderRequestAckParam `json:"taskCancelOrderRequestAck,required"`
}

func (r EventTaskTaskCancelOrderRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TaskCancelOrderRequestAck)
}

type EventTaskTaskPlaceOrderRequestAckParams struct {
	TaskPlaceOrderRequestAck TaskPlaceOrderRequestAckParam `json:"taskPlaceOrderRequestAck,required"`
}

func (r EventTaskTaskPlaceOrderRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TaskPlaceOrderRequestAck)
}

type EventTaskTaskQuoteRequestAckParams struct {
	TaskQuoteRequestAck TaskQuoteRequestAckParam `json:"taskQuoteRequestAck,required"`
}

func (r EventTaskTaskQuoteRequestAckParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TaskQuoteRequestAck)
}
