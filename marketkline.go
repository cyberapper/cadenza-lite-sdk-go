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

// MarketKlineService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketKlineService] method instead.
type MarketKlineService struct {
	Options []option.RequestOption
}

// NewMarketKlineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketKlineService(opts ...option.RequestOption) (r *MarketKlineService) {
	r = &MarketKlineService{}
	r.Options = opts
	return
}

// Get historical kline data
func (r *MarketKlineService) Get(ctx context.Context, query MarketKlineGetParams, opts ...option.RequestOption) (res *MarketKlineGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market/kline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Ohlcv struct {
	// Close price
	C float64 `json:"c"`
	// High price
	H float64 `json:"h"`
	// Low price
	L float64 `json:"l"`
	// Open price
	O float64 `json:"o"`
	// Start time (in unix milliseconds)
	T int64 `json:"t"`
	// Volume
	V    float64   `json:"v"`
	JSON ohlcvJSON `json:"-"`
}

// ohlcvJSON contains the JSON metadata for the struct [Ohlcv]
type ohlcvJSON struct {
	C           apijson.Field
	H           apijson.Field
	L           apijson.Field
	O           apijson.Field
	T           apijson.Field
	V           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ohlcv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ohlcvJSON) RawJSON() string {
	return r.raw
}

type OhlcvParam struct {
	// Close price
	C param.Field[float64] `json:"c"`
	// High price
	H param.Field[float64] `json:"h"`
	// Low price
	L param.Field[float64] `json:"l"`
	// Open price
	O param.Field[float64] `json:"o"`
	// Start time (in unix milliseconds)
	T param.Field[int64] `json:"t"`
	// Volume
	V param.Field[float64] `json:"v"`
}

func (r OhlcvParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarketKlineGetResponse struct {
	Candles []Ohlcv `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType MarketKlineGetResponseExchangeType `json:"exchangeType"`
	Interval     MarketKlineGetResponseInterval     `json:"interval"`
	Symbol       string                             `json:"symbol"`
	JSON         marketKlineGetResponseJSON         `json:"-"`
}

// marketKlineGetResponseJSON contains the JSON metadata for the struct
// [MarketKlineGetResponse]
type marketKlineGetResponseJSON struct {
	Candles           apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Interval          apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MarketKlineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketKlineGetResponseJSON) RawJSON() string {
	return r.raw
}

// Exchange type
type MarketKlineGetResponseExchangeType string

const (
	MarketKlineGetResponseExchangeTypeBinance       MarketKlineGetResponseExchangeType = "BINANCE"
	MarketKlineGetResponseExchangeTypeBinanceMargin MarketKlineGetResponseExchangeType = "BINANCE_MARGIN"
	MarketKlineGetResponseExchangeTypeB2C2          MarketKlineGetResponseExchangeType = "B2C2"
	MarketKlineGetResponseExchangeTypeWintermute    MarketKlineGetResponseExchangeType = "WINTERMUTE"
	MarketKlineGetResponseExchangeTypeBlockfills    MarketKlineGetResponseExchangeType = "BLOCKFILLS"
	MarketKlineGetResponseExchangeTypeStonex        MarketKlineGetResponseExchangeType = "STONEX"
)

func (r MarketKlineGetResponseExchangeType) IsKnown() bool {
	switch r {
	case MarketKlineGetResponseExchangeTypeBinance, MarketKlineGetResponseExchangeTypeBinanceMargin, MarketKlineGetResponseExchangeTypeB2C2, MarketKlineGetResponseExchangeTypeWintermute, MarketKlineGetResponseExchangeTypeBlockfills, MarketKlineGetResponseExchangeTypeStonex:
		return true
	}
	return false
}

type MarketKlineGetResponseInterval string

const (
	MarketKlineGetResponseInterval1s  MarketKlineGetResponseInterval = "1s"
	MarketKlineGetResponseInterval1m  MarketKlineGetResponseInterval = "1m"
	MarketKlineGetResponseInterval5m  MarketKlineGetResponseInterval = "5m"
	MarketKlineGetResponseInterval15m MarketKlineGetResponseInterval = "15m"
	MarketKlineGetResponseInterval30m MarketKlineGetResponseInterval = "30m"
	MarketKlineGetResponseInterval1h  MarketKlineGetResponseInterval = "1h"
	MarketKlineGetResponseInterval2h  MarketKlineGetResponseInterval = "2h"
	MarketKlineGetResponseInterval1d  MarketKlineGetResponseInterval = "1d"
	MarketKlineGetResponseInterval1w  MarketKlineGetResponseInterval = "1w"
)

func (r MarketKlineGetResponseInterval) IsKnown() bool {
	switch r {
	case MarketKlineGetResponseInterval1s, MarketKlineGetResponseInterval1m, MarketKlineGetResponseInterval5m, MarketKlineGetResponseInterval15m, MarketKlineGetResponseInterval30m, MarketKlineGetResponseInterval1h, MarketKlineGetResponseInterval2h, MarketKlineGetResponseInterval1d, MarketKlineGetResponseInterval1w:
		return true
	}
	return false
}

type MarketKlineGetParams struct {
	// Exchange type
	ExchangeType param.Field[MarketKlineGetParamsExchangeType] `query:"exchangeType,required"`
	// Kline interval
	Interval param.Field[MarketKlineGetParamsInterval] `query:"interval,required"`
	// Symbol
	Symbol param.Field[string] `query:"symbol,required"`
	// End time (in unix milliseconds)
	EndTime param.Field[int64] `query:"endTime"`
	// Limit the number of returned results.
	Limit param.Field[int64] `query:"limit"`
	// Start time (in unix milliseconds)
	StartTime param.Field[int64] `query:"startTime"`
}

// URLQuery serializes [MarketKlineGetParams]'s query parameters as `url.Values`.
func (r MarketKlineGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exchange type
type MarketKlineGetParamsExchangeType string

const (
	MarketKlineGetParamsExchangeTypeBinance       MarketKlineGetParamsExchangeType = "BINANCE"
	MarketKlineGetParamsExchangeTypeBinanceMargin MarketKlineGetParamsExchangeType = "BINANCE_MARGIN"
	MarketKlineGetParamsExchangeTypeB2C2          MarketKlineGetParamsExchangeType = "B2C2"
	MarketKlineGetParamsExchangeTypeWintermute    MarketKlineGetParamsExchangeType = "WINTERMUTE"
	MarketKlineGetParamsExchangeTypeBlockfills    MarketKlineGetParamsExchangeType = "BLOCKFILLS"
	MarketKlineGetParamsExchangeTypeStonex        MarketKlineGetParamsExchangeType = "STONEX"
)

func (r MarketKlineGetParamsExchangeType) IsKnown() bool {
	switch r {
	case MarketKlineGetParamsExchangeTypeBinance, MarketKlineGetParamsExchangeTypeBinanceMargin, MarketKlineGetParamsExchangeTypeB2C2, MarketKlineGetParamsExchangeTypeWintermute, MarketKlineGetParamsExchangeTypeBlockfills, MarketKlineGetParamsExchangeTypeStonex:
		return true
	}
	return false
}

// Kline interval
type MarketKlineGetParamsInterval string

const (
	MarketKlineGetParamsInterval1s  MarketKlineGetParamsInterval = "1s"
	MarketKlineGetParamsInterval1m  MarketKlineGetParamsInterval = "1m"
	MarketKlineGetParamsInterval5m  MarketKlineGetParamsInterval = "5m"
	MarketKlineGetParamsInterval15m MarketKlineGetParamsInterval = "15m"
	MarketKlineGetParamsInterval30m MarketKlineGetParamsInterval = "30m"
	MarketKlineGetParamsInterval1h  MarketKlineGetParamsInterval = "1h"
	MarketKlineGetParamsInterval2h  MarketKlineGetParamsInterval = "2h"
	MarketKlineGetParamsInterval1d  MarketKlineGetParamsInterval = "1d"
	MarketKlineGetParamsInterval1w  MarketKlineGetParamsInterval = "1w"
)

func (r MarketKlineGetParamsInterval) IsKnown() bool {
	switch r {
	case MarketKlineGetParamsInterval1s, MarketKlineGetParamsInterval1m, MarketKlineGetParamsInterval5m, MarketKlineGetParamsInterval15m, MarketKlineGetParamsInterval30m, MarketKlineGetParamsInterval1h, MarketKlineGetParamsInterval2h, MarketKlineGetParamsInterval1d, MarketKlineGetParamsInterval1w:
		return true
	}
	return false
}
