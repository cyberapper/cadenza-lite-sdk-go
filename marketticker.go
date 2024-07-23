// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzalite

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cadenza-lite-go/internal/apijson"
	"github.com/stainless-sdks/cadenza-lite-go/internal/apiquery"
	"github.com/stainless-sdks/cadenza-lite-go/internal/param"
	"github.com/stainless-sdks/cadenza-lite-go/internal/requestconfig"
	"github.com/stainless-sdks/cadenza-lite-go/option"
)

// MarketTickerService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketTickerService] method instead.
type MarketTickerService struct {
	Options []option.RequestOption
}

// NewMarketTickerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketTickerService(opts ...option.RequestOption) (r *MarketTickerService) {
	r = &MarketTickerService{}
	r.Options = opts
	return
}

// Symbol price
func (r *MarketTickerService) Get(ctx context.Context, query MarketTickerGetParams, opts ...option.RequestOption) (res *[]Ticker, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market/ticker"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Ticker struct {
	// Exchange type
	ExchangeType TickerExchangeType `json:"exchangeType,required"`
	// Symbol
	Symbol string `json:"symbol,required"`
	// Unix timestamp in milliseconds
	Timestamp int64 `json:"timestamp,required"`
	// Ask price
	AskPrice float64 `json:"askPrice"`
	// Ask quantity
	AskQuantity float64 `json:"askQuantity"`
	// Bid price
	BidPrice float64 `json:"bidPrice"`
	// Bid quantity
	BidQuantity float64 `json:"bidQuantity"`
	// Last price
	LastPrice float64 `json:"lastPrice"`
	// Last quantity
	LastQuantity float64    `json:"lastQuantity"`
	JSON         tickerJSON `json:"-"`
}

// tickerJSON contains the JSON metadata for the struct [Ticker]
type tickerJSON struct {
	ExchangeType apijson.Field
	Symbol       apijson.Field
	Timestamp    apijson.Field
	AskPrice     apijson.Field
	AskQuantity  apijson.Field
	BidPrice     apijson.Field
	BidQuantity  apijson.Field
	LastPrice    apijson.Field
	LastQuantity apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Ticker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tickerJSON) RawJSON() string {
	return r.raw
}

// Exchange type
type TickerExchangeType string

const (
	TickerExchangeTypeBinance       TickerExchangeType = "BINANCE"
	TickerExchangeTypeBinanceMargin TickerExchangeType = "BINANCE_MARGIN"
	TickerExchangeTypeB2C2          TickerExchangeType = "B2C2"
	TickerExchangeTypeWintermute    TickerExchangeType = "WINTERMUTE"
	TickerExchangeTypeBlockfills    TickerExchangeType = "BLOCKFILLS"
	TickerExchangeTypeStonex        TickerExchangeType = "STONEX"
)

func (r TickerExchangeType) IsKnown() bool {
	switch r {
	case TickerExchangeTypeBinance, TickerExchangeTypeBinanceMargin, TickerExchangeTypeB2C2, TickerExchangeTypeWintermute, TickerExchangeTypeBlockfills, TickerExchangeTypeStonex:
		return true
	}
	return false
}

type MarketTickerGetParams struct {
	// Symbol
	Symbol param.Field[string] `query:"symbol,required"`
	// Exchange type
	ExchangeType param.Field[MarketTickerGetParamsExchangeType] `query:"exchangeType"`
}

// URLQuery serializes [MarketTickerGetParams]'s query parameters as `url.Values`.
func (r MarketTickerGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exchange type
type MarketTickerGetParamsExchangeType string

const (
	MarketTickerGetParamsExchangeTypeBinance       MarketTickerGetParamsExchangeType = "BINANCE"
	MarketTickerGetParamsExchangeTypeBinanceMargin MarketTickerGetParamsExchangeType = "BINANCE_MARGIN"
	MarketTickerGetParamsExchangeTypeB2C2          MarketTickerGetParamsExchangeType = "B2C2"
	MarketTickerGetParamsExchangeTypeWintermute    MarketTickerGetParamsExchangeType = "WINTERMUTE"
	MarketTickerGetParamsExchangeTypeBlockfills    MarketTickerGetParamsExchangeType = "BLOCKFILLS"
	MarketTickerGetParamsExchangeTypeStonex        MarketTickerGetParamsExchangeType = "STONEX"
)

func (r MarketTickerGetParamsExchangeType) IsKnown() bool {
	switch r {
	case MarketTickerGetParamsExchangeTypeBinance, MarketTickerGetParamsExchangeTypeBinanceMargin, MarketTickerGetParamsExchangeTypeB2C2, MarketTickerGetParamsExchangeTypeWintermute, MarketTickerGetParamsExchangeTypeBlockfills, MarketTickerGetParamsExchangeTypeStonex:
		return true
	}
	return false
}
