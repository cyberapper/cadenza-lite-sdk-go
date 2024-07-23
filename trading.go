// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk

import (
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
)

// TradingService contains methods and other services that help with interacting
// with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTradingService] method instead.
type TradingService struct {
	Options         []option.RequestOption
	Order           *TradingOrderService
	Quote           *TradingQuoteService
	ExecutionReport *TradingExecutionReportService
}

// NewTradingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTradingService(opts ...option.RequestOption) (r *TradingService) {
	r = &TradingService{}
	r.Options = opts
	r.Order = NewTradingOrderService(opts...)
	r.Quote = NewTradingQuoteService(opts...)
	r.ExecutionReport = NewTradingExecutionReportService(opts...)
	return
}
