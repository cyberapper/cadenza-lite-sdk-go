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
	Payload interface{} `json:"payload"`
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
	EventEventTypeCadenzaTaskQuote                     EventEventType = "cadenza.task.quote"
	EventEventTypeCadenzaDropCopyQuoteRequestAck       EventEventType = "cadenza.dropCopy.quoteRequestAck"
	EventEventTypeCadenzaDropCopyPlaceOrderRequestAck  EventEventType = "cadenza.dropCopy.placeOrderRequestAck"
	EventEventTypeCadenzaDropCopyCancelOrderRequestAck EventEventType = "cadenza.dropCopy.cancelOrderRequestAck"
	EventEventTypeCadenzaDropCopyQuote                 EventEventType = "cadenza.dropCopy.quote"
	EventEventTypeCadenzaDropCopyOrder                 EventEventType = "cadenza.dropCopy.order"
	EventEventTypeCadenzaDropCopyExecutionReport       EventEventType = "cadenza.dropCopy.executionReport"
	EventEventTypeCadenzaDropCopyPortfolio             EventEventType = "cadenza.dropCopy.portfolio"
	EventEventTypeCadenzaMarketDataOrderBook           EventEventType = "cadenza.marketData.orderBook"
	EventEventTypeCadenzaMarketDataKline               EventEventType = "cadenza.marketData.kline"
)

func (r EventEventType) IsKnown() bool {
	switch r {
	case EventEventTypeCadenzaTaskQuote, EventEventTypeCadenzaDropCopyQuoteRequestAck, EventEventTypeCadenzaDropCopyPlaceOrderRequestAck, EventEventTypeCadenzaDropCopyCancelOrderRequestAck, EventEventTypeCadenzaDropCopyQuote, EventEventTypeCadenzaDropCopyOrder, EventEventTypeCadenzaDropCopyExecutionReport, EventEventTypeCadenzaDropCopyPortfolio, EventEventTypeCadenzaMarketDataOrderBook, EventEventTypeCadenzaMarketDataKline:
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
	Payload param.Field[interface{}] `json:"payload"`
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventNewParams struct {
	Event EventParam `json:"event,required"`
}

func (r EventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Event)
}
