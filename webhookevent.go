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

// WebhookEventService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEventService] method instead.
type WebhookEventService struct {
	Options []option.RequestOption
}

// NewWebhookEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEventService(opts ...option.RequestOption) (r *WebhookEventService) {
	r = &WebhookEventService{}
	r.Options = opts
	return
}

// PubSub event handler placeholder for quote event
func (r *WebhookEventService) DropCopyQuote(ctx context.Context, body WebhookEventDropCopyQuoteParams, opts ...option.RequestOption) (res *DropCopyQuote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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

type WebhookEventDropCopyQuoteParams struct {
	DropCopyQuote DropCopyQuoteParam `json:"dropCopyQuote,required"`
}

func (r WebhookEventDropCopyQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyQuote)
}
