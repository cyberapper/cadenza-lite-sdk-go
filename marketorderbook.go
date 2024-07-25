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

// MarketOrderbookService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketOrderbookService] method instead.
type MarketOrderbookService struct {
	Options []option.RequestOption
}

// NewMarketOrderbookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMarketOrderbookService(opts ...option.RequestOption) (r *MarketOrderbookService) {
	r = &MarketOrderbookService{}
	r.Options = opts
	return
}

// Get order book
func (r *MarketOrderbookService) Get(ctx context.Context, query MarketOrderbookGetParams, opts ...option.RequestOption) (res *[]Orderbook, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market/orderbook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
	// UUID string
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	ExchangeType      string `json:"exchangeType"`
	// Order book level
	Level  int64         `json:"level"`
	Symbol string        `json:"symbol"`
	JSON   orderbookJSON `json:"-"`
}

// orderbookJSON contains the JSON metadata for the struct [Orderbook]
type orderbookJSON struct {
	Asks              apijson.Field
	Bids              apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Level             apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Orderbook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderbookJSON) RawJSON() string {
	return r.raw
}

func (r Orderbook) implementsEventPayload() {}

func (r Orderbook) implementsOrder() {}

func (r Orderbook) implementsExchangeAccountPortfolio() {}

func (r Orderbook) implementsQuote() {}

func (r Orderbook) implementsOrderbook() {}

type OrderbookParam struct {
	Asks param.Field[[][]float64] `json:"asks"`
	Bids param.Field[[][]float64] `json:"bids"`
	// UUID string
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	ExchangeType      param.Field[string] `json:"exchangeType"`
	// Order book level
	Level  param.Field[int64]  `json:"level"`
	Symbol param.Field[string] `json:"symbol"`
}

func (r OrderbookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrderbookParam) implementsOrderUnionParam() {}

func (r OrderbookParam) implementsExchangeAccountPortfolioUnionParam() {}

func (r OrderbookParam) implementsQuoteUnionParam() {}

func (r OrderbookParam) implementsEventMarketDataKlineParamsPayloadUnion() {}

func (r OrderbookParam) implementsOrderbookUnionParam() {}

type MarketOrderbookGetParams struct {
	// Exchange type
	ExchangeType param.Field[MarketOrderbookGetParamsExchangeType] `query:"exchangeType,required"`
	// Symbol
	Symbol param.Field[string] `query:"symbol,required"`
	// Limit the number of returned results.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [MarketOrderbookGetParams]'s query parameters as
// `url.Values`.
func (r MarketOrderbookGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exchange type
type MarketOrderbookGetParamsExchangeType string

const (
	MarketOrderbookGetParamsExchangeTypeBinance       MarketOrderbookGetParamsExchangeType = "BINANCE"
	MarketOrderbookGetParamsExchangeTypeBinanceMargin MarketOrderbookGetParamsExchangeType = "BINANCE_MARGIN"
	MarketOrderbookGetParamsExchangeTypeB2C2          MarketOrderbookGetParamsExchangeType = "B2C2"
	MarketOrderbookGetParamsExchangeTypeWintermute    MarketOrderbookGetParamsExchangeType = "WINTERMUTE"
	MarketOrderbookGetParamsExchangeTypeBlockfills    MarketOrderbookGetParamsExchangeType = "BLOCKFILLS"
	MarketOrderbookGetParamsExchangeTypeStonex        MarketOrderbookGetParamsExchangeType = "STONEX"
)

func (r MarketOrderbookGetParamsExchangeType) IsKnown() bool {
	switch r {
	case MarketOrderbookGetParamsExchangeTypeBinance, MarketOrderbookGetParamsExchangeTypeBinanceMargin, MarketOrderbookGetParamsExchangeTypeB2C2, MarketOrderbookGetParamsExchangeTypeWintermute, MarketOrderbookGetParamsExchangeTypeBlockfills, MarketOrderbookGetParamsExchangeTypeStonex:
		return true
	}
	return false
}
