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

// MarketInstrumentService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketInstrumentService] method instead.
type MarketInstrumentService struct {
	Options []option.RequestOption
}

// NewMarketInstrumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMarketInstrumentService(opts ...option.RequestOption) (r *MarketInstrumentService) {
	r = &MarketInstrumentService{}
	r.Options = opts
	return
}

// List available exchange symbols
func (r *MarketInstrumentService) List(ctx context.Context, query MarketInstrumentListParams, opts ...option.RequestOption) (res *[]Instrument, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market/listSymbolInfo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Trading Instrument
type Instrument struct {
	// Exchange type
	ExchangeType InstrumentExchangeType `json:"exchangeType,required"`
	// Symbol name
	Symbol string `json:"symbol,required"`
	// Base currency
	BaseSymbol string `json:"baseSymbol"`
	// Symbol description
	Description string `json:"description"`
	// Margin rate
	MarginRate float64 `json:"marginRate"`
	// Max quantity
	MaxQuantity float64 `json:"maxQuantity"`
	// Min quantity
	MinQuantity float64 `json:"minQuantity"`
	// Min tick, Price Tick
	MinTick float64 `json:"minTick"`
	// Supported order types
	OrderTypes []InstrumentOrderType `json:"orderTypes"`
	// Pip size
	PipSize float64 `json:"pipSize"`
	// Pip value
	PipValue float64 `json:"pipValue"`
	// Price precision
	PricePrecision int64 `json:"pricePrecision"`
	// Quantity precision
	QuantityPrecision int64 `json:"quantityPrecision"`
	// Quantity step, round lot
	QuantityStep float64 `json:"quantityStep"`
	// Quoted currency
	QuoteSymbol string `json:"quoteSymbol"`
	// Security type
	SecurityType InstrumentSecurityType `json:"securityType"`
	// Supported time in force
	TimeInForce []InstrumentTimeInForce `json:"timeInForce"`
	JSON        instrumentJSON          `json:"-"`
}

// instrumentJSON contains the JSON metadata for the struct [Instrument]
type instrumentJSON struct {
	ExchangeType      apijson.Field
	Symbol            apijson.Field
	BaseSymbol        apijson.Field
	Description       apijson.Field
	MarginRate        apijson.Field
	MaxQuantity       apijson.Field
	MinQuantity       apijson.Field
	MinTick           apijson.Field
	OrderTypes        apijson.Field
	PipSize           apijson.Field
	PipValue          apijson.Field
	PricePrecision    apijson.Field
	QuantityPrecision apijson.Field
	QuantityStep      apijson.Field
	QuoteSymbol       apijson.Field
	SecurityType      apijson.Field
	TimeInForce       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Instrument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instrumentJSON) RawJSON() string {
	return r.raw
}

// Exchange type
type InstrumentExchangeType string

const (
	InstrumentExchangeTypeBinance       InstrumentExchangeType = "BINANCE"
	InstrumentExchangeTypeBinanceMargin InstrumentExchangeType = "BINANCE_MARGIN"
	InstrumentExchangeTypeB2C2          InstrumentExchangeType = "B2C2"
	InstrumentExchangeTypeWintermute    InstrumentExchangeType = "WINTERMUTE"
	InstrumentExchangeTypeBlockfills    InstrumentExchangeType = "BLOCKFILLS"
	InstrumentExchangeTypeStonex        InstrumentExchangeType = "STONEX"
)

func (r InstrumentExchangeType) IsKnown() bool {
	switch r {
	case InstrumentExchangeTypeBinance, InstrumentExchangeTypeBinanceMargin, InstrumentExchangeTypeB2C2, InstrumentExchangeTypeWintermute, InstrumentExchangeTypeBlockfills, InstrumentExchangeTypeStonex:
		return true
	}
	return false
}

// Order type
type InstrumentOrderType string

const (
	InstrumentOrderTypeMarket          InstrumentOrderType = "MARKET"
	InstrumentOrderTypeLimit           InstrumentOrderType = "LIMIT"
	InstrumentOrderTypeStopLoss        InstrumentOrderType = "STOP_LOSS"
	InstrumentOrderTypeStopLossLimit   InstrumentOrderType = "STOP_LOSS_LIMIT"
	InstrumentOrderTypeTakeProfit      InstrumentOrderType = "TAKE_PROFIT"
	InstrumentOrderTypeTakeProfitLimit InstrumentOrderType = "TAKE_PROFIT_LIMIT"
	InstrumentOrderTypeQuoted          InstrumentOrderType = "QUOTED"
)

func (r InstrumentOrderType) IsKnown() bool {
	switch r {
	case InstrumentOrderTypeMarket, InstrumentOrderTypeLimit, InstrumentOrderTypeStopLoss, InstrumentOrderTypeStopLossLimit, InstrumentOrderTypeTakeProfit, InstrumentOrderTypeTakeProfitLimit, InstrumentOrderTypeQuoted:
		return true
	}
	return false
}

// Security type
type InstrumentSecurityType string

const (
	InstrumentSecurityTypeSpot       InstrumentSecurityType = "SPOT"
	InstrumentSecurityTypeCash       InstrumentSecurityType = "CASH"
	InstrumentSecurityTypeStock      InstrumentSecurityType = "STOCK"
	InstrumentSecurityTypeCrypto     InstrumentSecurityType = "CRYPTO"
	InstrumentSecurityTypeDerivative InstrumentSecurityType = "DERIVATIVE"
	InstrumentSecurityTypeOption     InstrumentSecurityType = "OPTION"
	InstrumentSecurityTypeFuture     InstrumentSecurityType = "FUTURE"
	InstrumentSecurityTypeForex      InstrumentSecurityType = "FOREX"
	InstrumentSecurityTypeCommodity  InstrumentSecurityType = "COMMODITY"
)

func (r InstrumentSecurityType) IsKnown() bool {
	switch r {
	case InstrumentSecurityTypeSpot, InstrumentSecurityTypeCash, InstrumentSecurityTypeStock, InstrumentSecurityTypeCrypto, InstrumentSecurityTypeDerivative, InstrumentSecurityTypeOption, InstrumentSecurityTypeFuture, InstrumentSecurityTypeForex, InstrumentSecurityTypeCommodity:
		return true
	}
	return false
}

// Time in force
type InstrumentTimeInForce string

const (
	InstrumentTimeInForceDay InstrumentTimeInForce = "DAY"
	InstrumentTimeInForceGtc InstrumentTimeInForce = "GTC"
	InstrumentTimeInForceGtx InstrumentTimeInForce = "GTX"
	InstrumentTimeInForceGtd InstrumentTimeInForce = "GTD"
	InstrumentTimeInForceOpg InstrumentTimeInForce = "OPG"
	InstrumentTimeInForceCls InstrumentTimeInForce = "CLS"
	InstrumentTimeInForceIoc InstrumentTimeInForce = "IOC"
	InstrumentTimeInForceFok InstrumentTimeInForce = "FOK"
	InstrumentTimeInForceGfa InstrumentTimeInForce = "GFA"
	InstrumentTimeInForceGfs InstrumentTimeInForce = "GFS"
	InstrumentTimeInForceGtm InstrumentTimeInForce = "GTM"
	InstrumentTimeInForceMoo InstrumentTimeInForce = "MOO"
	InstrumentTimeInForceMoc InstrumentTimeInForce = "MOC"
	InstrumentTimeInForceExt InstrumentTimeInForce = "EXT"
)

func (r InstrumentTimeInForce) IsKnown() bool {
	switch r {
	case InstrumentTimeInForceDay, InstrumentTimeInForceGtc, InstrumentTimeInForceGtx, InstrumentTimeInForceGtd, InstrumentTimeInForceOpg, InstrumentTimeInForceCls, InstrumentTimeInForceIoc, InstrumentTimeInForceFok, InstrumentTimeInForceGfa, InstrumentTimeInForceGfs, InstrumentTimeInForceGtm, InstrumentTimeInForceMoo, InstrumentTimeInForceMoc, InstrumentTimeInForceExt:
		return true
	}
	return false
}

type MarketInstrumentListParams struct {
	// Whether to return detailed information
	Detail param.Field[bool] `query:"detail"`
	// Exchange type
	ExchangeType param.Field[MarketInstrumentListParamsExchangeType] `query:"exchangeType"`
	// Symbol
	Symbol param.Field[string] `query:"symbol"`
}

// URLQuery serializes [MarketInstrumentListParams]'s query parameters as
// `url.Values`.
func (r MarketInstrumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Exchange type
type MarketInstrumentListParamsExchangeType string

const (
	MarketInstrumentListParamsExchangeTypeBinance       MarketInstrumentListParamsExchangeType = "BINANCE"
	MarketInstrumentListParamsExchangeTypeBinanceMargin MarketInstrumentListParamsExchangeType = "BINANCE_MARGIN"
	MarketInstrumentListParamsExchangeTypeB2C2          MarketInstrumentListParamsExchangeType = "B2C2"
	MarketInstrumentListParamsExchangeTypeWintermute    MarketInstrumentListParamsExchangeType = "WINTERMUTE"
	MarketInstrumentListParamsExchangeTypeBlockfills    MarketInstrumentListParamsExchangeType = "BLOCKFILLS"
	MarketInstrumentListParamsExchangeTypeStonex        MarketInstrumentListParamsExchangeType = "STONEX"
)

func (r MarketInstrumentListParamsExchangeType) IsKnown() bool {
	switch r {
	case MarketInstrumentListParamsExchangeTypeBinance, MarketInstrumentListParamsExchangeTypeBinanceMargin, MarketInstrumentListParamsExchangeTypeB2C2, MarketInstrumentListParamsExchangeTypeWintermute, MarketInstrumentListParamsExchangeTypeBlockfills, MarketInstrumentListParamsExchangeTypeStonex:
		return true
	}
	return false
}
