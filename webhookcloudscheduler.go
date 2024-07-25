// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk

import (
	"context"
	"net/http"

	"github.com/cyberapper/cadenza-lite-sdk-go/internal/requestconfig"
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
)

// WebhookCloudSchedulerService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookCloudSchedulerService] method instead.
type WebhookCloudSchedulerService struct {
	Options []option.RequestOption
}

// NewWebhookCloudSchedulerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebhookCloudSchedulerService(opts ...option.RequestOption) (r *WebhookCloudSchedulerService) {
	r = &WebhookCloudSchedulerService{}
	r.Options = opts
	return
}

// Cloud scheduler update portfolio routine task
func (r *WebhookCloudSchedulerService) UpdatePortfolioRoutine(ctx context.Context, opts ...option.RequestOption) (res *WebhookCloudSchedulerUpdatePortfolioRoutineResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/cloudScheduler/updatePortfolioRoutine"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type WebhookCloudSchedulerUpdatePortfolioRoutineResponse = interface{}
