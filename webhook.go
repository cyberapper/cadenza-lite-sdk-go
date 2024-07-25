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

// WebhookService contains methods and other services that help with interacting
// with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options        []option.RequestOption
	CloudScheduler *WebhookCloudSchedulerService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	r.CloudScheduler = NewWebhookCloudSchedulerService(opts...)
	return
}

// PubSub Event Handler
func (r *WebhookService) Pubsub(ctx context.Context, body WebhookPubsubParams, opts ...option.RequestOption) (res *WebhookPubsubResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub"
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
	// The source system or module that generated the event.
	Source string    `json:"source"`
	JSON   eventJSON `json:"-"`
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
	EventID     apijson.Field
	EventType   apijson.Field
	Timestamp   apijson.Field
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
	EventEventTypeCadenzaDropCopyPortfolio         EventEventType = "cadenza.dropCopy.portfolio"
	EventEventTypeCadenzaMarketDataOrderBook       EventEventType = "cadenza.marketData.orderBook"
	EventEventTypeCadenzaMarketDataKline           EventEventType = "cadenza.marketData.kline"
)

func (r EventEventType) IsKnown() bool {
	switch r {
	case EventEventTypeCadenzaTaskQuoteRequestAck, EventEventTypeCadenzaTaskPlaceOrderRequestAck, EventEventTypeCadenzaTaskCancelOrderRequestAck, EventEventTypeCadenzaDropCopyQuote, EventEventTypeCadenzaDropCopyOrder, EventEventTypeCadenzaDropCopyPortfolio, EventEventTypeCadenzaMarketDataOrderBook, EventEventTypeCadenzaMarketDataKline:
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
	// The source system or module that generated the event.
	Source param.Field[string] `json:"source"`
}

func (r EventParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookPubsubResponse struct {
	Data string                    `json:"data"`
	JSON webhookPubsubResponseJSON `json:"-"`
}

// webhookPubsubResponseJSON contains the JSON metadata for the struct
// [WebhookPubsubResponse]
type webhookPubsubResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookPubsubResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookPubsubResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookPubsubParams struct {
	Message param.Field[WebhookPubsubParamsMessage] `json:"message,required"`
	// The subscription name.
	Subscription param.Field[string] `json:"subscription,required"`
}

func (r WebhookPubsubParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookPubsubParamsMessage struct {
	// The unique identifier for the message.
	ID param.Field[string] `json:"id,required"`
	// The base64-encoded data.
	Data param.Field[string] `json:"data" format:"byte"`
}

func (r WebhookPubsubParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
