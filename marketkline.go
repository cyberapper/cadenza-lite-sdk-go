// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cyberapper/cadenza-lite-sdk-go/internal/apijson"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/apiquery"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/param"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/requestconfig"
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
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
func (r *MarketKlineService) Get(ctx context.Context, query MarketKlineGetParams, opts ...option.RequestOption) (res *Kline, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market/kline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Candles []CandlesItem

type CandlesItem struct {
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
	V    float64         `json:"v"`
	JSON candlesItemJSON `json:"-"`
}

// candlesItemJSON contains the JSON metadata for the struct [CandlesItem]
type candlesItemJSON struct {
	C           apijson.Field
	H           apijson.Field
	L           apijson.Field
	O           apijson.Field
	T           apijson.Field
	V           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CandlesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r candlesItemJSON) RawJSON() string {
	return r.raw
}

type CandlesParam []CandlesItemParam

type CandlesItemParam struct {
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

func (r CandlesItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Kline struct {
	Candles Candles `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType KlineExchangeType `json:"exchangeType"`
	Interval     KlineInterval     `json:"interval"`
	Symbol       string            `json:"symbol"`
	JSON         klineJSON         `json:"-"`
}

// klineJSON contains the JSON metadata for the struct [Kline]
type klineJSON struct {
	Candles           apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Interval          apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Kline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r klineJSON) RawJSON() string {
	return r.raw
}

func (r Kline) implementsEventPayload() {}

// Exchange type
type KlineExchangeType string

const (
	KlineExchangeTypeBinance       KlineExchangeType = "BINANCE"
	KlineExchangeTypeBinanceMargin KlineExchangeType = "BINANCE_MARGIN"
	KlineExchangeTypeB2C2          KlineExchangeType = "B2C2"
	KlineExchangeTypeWintermute    KlineExchangeType = "WINTERMUTE"
	KlineExchangeTypeBlockfills    KlineExchangeType = "BLOCKFILLS"
	KlineExchangeTypeStonex        KlineExchangeType = "STONEX"
)

func (r KlineExchangeType) IsKnown() bool {
	switch r {
	case KlineExchangeTypeBinance, KlineExchangeTypeBinanceMargin, KlineExchangeTypeB2C2, KlineExchangeTypeWintermute, KlineExchangeTypeBlockfills, KlineExchangeTypeStonex:
		return true
	}
	return false
}

type KlineInterval string

const (
	KlineInterval1s  KlineInterval = "1s"
	KlineInterval1m  KlineInterval = "1m"
	KlineInterval5m  KlineInterval = "5m"
	KlineInterval15m KlineInterval = "15m"
	KlineInterval30m KlineInterval = "30m"
	KlineInterval1h  KlineInterval = "1h"
	KlineInterval2h  KlineInterval = "2h"
	KlineInterval1d  KlineInterval = "1d"
	KlineInterval1w  KlineInterval = "1w"
)

func (r KlineInterval) IsKnown() bool {
	switch r {
	case KlineInterval1s, KlineInterval1m, KlineInterval5m, KlineInterval15m, KlineInterval30m, KlineInterval1h, KlineInterval2h, KlineInterval1d, KlineInterval1w:
		return true
	}
	return false
}

type KlineParam struct {
	Candles param.Field[CandlesParam] `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[KlineExchangeType] `json:"exchangeType"`
	Interval     param.Field[KlineInterval]     `json:"interval"`
	Symbol       param.Field[string]            `json:"symbol"`
}

func (r KlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KlineParam) implementsEventPayloadUnionParam() {}

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
