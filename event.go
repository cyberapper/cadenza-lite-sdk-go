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

// EventService contains methods and other services that help with interacting with
// the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// PubSub event handler placeholder for order event
func (r *EventService) DropCopyOrder(ctx context.Context, body EventDropCopyOrderParams, opts ...option.RequestOption) (res *DropCopyOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/order"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for portfolio event
func (r *EventService) DropCopyPortfolio(ctx context.Context, body EventDropCopyPortfolioParams, opts ...option.RequestOption) (res *DropCopyPortfolio, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/portfolio"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for quote event
func (r *EventService) DropCopyQuote(ctx context.Context, body EventDropCopyQuoteParams, opts ...option.RequestOption) (res *DropCopyQuote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/dropCopy/quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for kline event
func (r *EventService) MarketDataKline(ctx context.Context, body EventMarketDataKlineParams, opts ...option.RequestOption) (res *MarketDataKline, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/kline"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PubSub event handler placeholder for order book event
func (r *EventService) MarketDataOrderBook(ctx context.Context, body EventMarketDataOrderBookParams, opts ...option.RequestOption) (res *MarketDataOrderBook, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/webhook/pubsub/marketData/orderBook"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DropCopyOrder struct {
	Payload Order             `json:"payload"`
	JSON    dropCopyOrderJSON `json:"-"`
	Event
}

// dropCopyOrderJSON contains the JSON metadata for the struct [DropCopyOrder]
type dropCopyOrderJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyOrderJSON) RawJSON() string {
	return r.raw
}

type DropCopyOrderParam struct {
	Payload param.Field[OrderParam] `json:"payload"`
	EventParam
}

func (r DropCopyOrderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DropCopyPortfolio struct {
	Payload ExchangeAccountPortfolio `json:"payload"`
	JSON    dropCopyPortfolioJSON    `json:"-"`
	Event
}

// dropCopyPortfolioJSON contains the JSON metadata for the struct
// [DropCopyPortfolio]
type dropCopyPortfolioJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DropCopyPortfolio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dropCopyPortfolioJSON) RawJSON() string {
	return r.raw
}

type DropCopyPortfolioParam struct {
	Payload param.Field[ExchangeAccountPortfolioParam] `json:"payload"`
	EventParam
}

func (r DropCopyPortfolioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type MarketDataKline struct {
	Payload MarketDataKlinePayload `json:"payload"`
	JSON    marketDataKlineJSON    `json:"-"`
	Event
}

// marketDataKlineJSON contains the JSON metadata for the struct [MarketDataKline]
type marketDataKlineJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketDataKline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataKlineJSON) RawJSON() string {
	return r.raw
}

type MarketDataKlinePayload struct {
	Candles []Ohlcv `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType MarketDataKlinePayloadExchangeType `json:"exchangeType"`
	Interval     MarketDataKlinePayloadInterval     `json:"interval"`
	Symbol       string                             `json:"symbol"`
	JSON         marketDataKlinePayloadJSON         `json:"-"`
}

// marketDataKlinePayloadJSON contains the JSON metadata for the struct
// [MarketDataKlinePayload]
type marketDataKlinePayloadJSON struct {
	Candles           apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Interval          apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MarketDataKlinePayload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataKlinePayloadJSON) RawJSON() string {
	return r.raw
}

// Exchange type
type MarketDataKlinePayloadExchangeType string

const (
	MarketDataKlinePayloadExchangeTypeBinance       MarketDataKlinePayloadExchangeType = "BINANCE"
	MarketDataKlinePayloadExchangeTypeBinanceMargin MarketDataKlinePayloadExchangeType = "BINANCE_MARGIN"
	MarketDataKlinePayloadExchangeTypeB2C2          MarketDataKlinePayloadExchangeType = "B2C2"
	MarketDataKlinePayloadExchangeTypeWintermute    MarketDataKlinePayloadExchangeType = "WINTERMUTE"
	MarketDataKlinePayloadExchangeTypeBlockfills    MarketDataKlinePayloadExchangeType = "BLOCKFILLS"
	MarketDataKlinePayloadExchangeTypeStonex        MarketDataKlinePayloadExchangeType = "STONEX"
)

func (r MarketDataKlinePayloadExchangeType) IsKnown() bool {
	switch r {
	case MarketDataKlinePayloadExchangeTypeBinance, MarketDataKlinePayloadExchangeTypeBinanceMargin, MarketDataKlinePayloadExchangeTypeB2C2, MarketDataKlinePayloadExchangeTypeWintermute, MarketDataKlinePayloadExchangeTypeBlockfills, MarketDataKlinePayloadExchangeTypeStonex:
		return true
	}
	return false
}

type MarketDataKlinePayloadInterval string

const (
	MarketDataKlinePayloadInterval1s  MarketDataKlinePayloadInterval = "1s"
	MarketDataKlinePayloadInterval1m  MarketDataKlinePayloadInterval = "1m"
	MarketDataKlinePayloadInterval5m  MarketDataKlinePayloadInterval = "5m"
	MarketDataKlinePayloadInterval15m MarketDataKlinePayloadInterval = "15m"
	MarketDataKlinePayloadInterval30m MarketDataKlinePayloadInterval = "30m"
	MarketDataKlinePayloadInterval1h  MarketDataKlinePayloadInterval = "1h"
	MarketDataKlinePayloadInterval2h  MarketDataKlinePayloadInterval = "2h"
	MarketDataKlinePayloadInterval1d  MarketDataKlinePayloadInterval = "1d"
	MarketDataKlinePayloadInterval1w  MarketDataKlinePayloadInterval = "1w"
)

func (r MarketDataKlinePayloadInterval) IsKnown() bool {
	switch r {
	case MarketDataKlinePayloadInterval1s, MarketDataKlinePayloadInterval1m, MarketDataKlinePayloadInterval5m, MarketDataKlinePayloadInterval15m, MarketDataKlinePayloadInterval30m, MarketDataKlinePayloadInterval1h, MarketDataKlinePayloadInterval2h, MarketDataKlinePayloadInterval1d, MarketDataKlinePayloadInterval1w:
		return true
	}
	return false
}

type MarketDataKlineParam struct {
	Payload param.Field[MarketDataKlinePayloadParam] `json:"payload"`
	EventParam
}

func (r MarketDataKlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarketDataKlinePayloadParam struct {
	Candles param.Field[[]OhlcvParam] `json:"candles"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[MarketDataKlinePayloadExchangeType] `json:"exchangeType"`
	Interval     param.Field[MarketDataKlinePayloadInterval]     `json:"interval"`
	Symbol       param.Field[string]                             `json:"symbol"`
}

func (r MarketDataKlinePayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MarketDataOrderBook struct {
	Payload Orderbook               `json:"payload"`
	JSON    marketDataOrderBookJSON `json:"-"`
	Event
}

// marketDataOrderBookJSON contains the JSON metadata for the struct
// [MarketDataOrderBook]
type marketDataOrderBookJSON struct {
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MarketDataOrderBook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r marketDataOrderBookJSON) RawJSON() string {
	return r.raw
}

type MarketDataOrderBookParam struct {
	Payload param.Field[OrderbookParam] `json:"payload"`
	EventParam
}

func (r MarketDataOrderBookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventDropCopyOrderParams struct {
	DropCopyOrder DropCopyOrderParam `json:"dropCopyOrder,required"`
}

func (r EventDropCopyOrderParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyOrder)
}

type EventDropCopyPortfolioParams struct {
	DropCopyPortfolio DropCopyPortfolioParam `json:"dropCopyPortfolio,required"`
}

func (r EventDropCopyPortfolioParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyPortfolio)
}

type EventDropCopyQuoteParams struct {
	DropCopyQuote DropCopyQuoteParam `json:"dropCopyQuote,required"`
}

func (r EventDropCopyQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DropCopyQuote)
}

type EventMarketDataKlineParams struct {
	MarketDataKline MarketDataKlineParam `json:"marketDataKline,required"`
}

func (r EventMarketDataKlineParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MarketDataKline)
}

type EventMarketDataOrderBookParams struct {
	MarketDataOrderBook MarketDataOrderBookParam `json:"marketDataOrderBook,required"`
}

func (r EventMarketDataOrderBookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MarketDataOrderBook)
}
