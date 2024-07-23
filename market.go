// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzalite

import (
	"github.com/stainless-sdks/cadenza-lite-go/option"
)

// MarketService contains methods and other services that help with interacting
// with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketService] method instead.
type MarketService struct {
	Options    []option.RequestOption
	Instrument *MarketInstrumentService
	Ticker     *MarketTickerService
	Orderbook  *MarketOrderbookService
	Kline      *MarketKlineService
}

// NewMarketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMarketService(opts ...option.RequestOption) (r *MarketService) {
	r = &MarketService{}
	r.Options = opts
	r.Instrument = NewMarketInstrumentService(opts...)
	r.Ticker = NewMarketTickerService(opts...)
	r.Orderbook = NewMarketOrderbookService(opts...)
	r.Kline = NewMarketKlineService(opts...)
	return
}
