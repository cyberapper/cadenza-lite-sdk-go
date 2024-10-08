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

// TradingQuoteService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTradingQuoteService] method instead.
type TradingQuoteService struct {
	Options []option.RequestOption
}

// NewTradingQuoteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTradingQuoteService(opts ...option.RequestOption) (r *TradingQuoteService) {
	r = &TradingQuoteService{}
	r.Options = opts
	return
}

// Quote will give the best quote from all available exchange accounts
func (r *TradingQuoteService) Post(ctx context.Context, body TradingQuotePostParams, opts ...option.RequestOption) (res *[]Quote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/trading/fetchQuotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Quote will give the best quote from all available exchange accounts
func (r *TradingQuoteService) RequestForQuote(ctx context.Context, body TradingQuoteRequestForQuoteParams, opts ...option.RequestOption) (res *[]Quote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/trading/fetchQuotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Quote struct {
	// Base currency
	BaseCurrency string `json:"baseCurrency,required"`
	// Quote currency
	QuoteCurrency string `json:"quoteCurrency,required"`
	// Quote request ID
	QuoteRequestID string `json:"quoteRequestId,required" format:"uuid"`
	// deprecated, alias of createdAt, Create time of the quote
	Timestamp int64 `json:"timestamp,required"`
	// deprecated, alias of expiredAtExpiration time of the quote
	ValidUntil int64 `json:"validUntil,required"`
	// Ask price
	AskPrice float64 `json:"askPrice"`
	// Ask quantity
	AskQuantity float64 `json:"askQuantity"`
	// Bid price
	BidPrice float64 `json:"bidPrice"`
	// Bid quantity
	BidQuantity float64 `json:"bidQuantity"`
	// Create time of the quote
	CreatedAt int64 `json:"createdAt"`
	// Exchange Account ID
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType QuoteExchangeType `json:"exchangeType"`
	// Expiration time of the quote
	ExpiredAt int64     `json:"expiredAt"`
	JSON      quoteJSON `json:"-"`
}

// quoteJSON contains the JSON metadata for the struct [Quote]
type quoteJSON struct {
	BaseCurrency      apijson.Field
	QuoteCurrency     apijson.Field
	QuoteRequestID    apijson.Field
	Timestamp         apijson.Field
	ValidUntil        apijson.Field
	AskPrice          apijson.Field
	AskQuantity       apijson.Field
	BidPrice          apijson.Field
	BidQuantity       apijson.Field
	CreatedAt         apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	ExpiredAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Quote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quoteJSON) RawJSON() string {
	return r.raw
}

// Exchange type
type QuoteExchangeType string

const (
	QuoteExchangeTypeBinance       QuoteExchangeType = "BINANCE"
	QuoteExchangeTypeBinanceMargin QuoteExchangeType = "BINANCE_MARGIN"
	QuoteExchangeTypeB2C2          QuoteExchangeType = "B2C2"
	QuoteExchangeTypeWintermute    QuoteExchangeType = "WINTERMUTE"
	QuoteExchangeTypeBlockfills    QuoteExchangeType = "BLOCKFILLS"
	QuoteExchangeTypeStonex        QuoteExchangeType = "STONEX"
)

func (r QuoteExchangeType) IsKnown() bool {
	switch r {
	case QuoteExchangeTypeBinance, QuoteExchangeTypeBinanceMargin, QuoteExchangeTypeB2C2, QuoteExchangeTypeWintermute, QuoteExchangeTypeBlockfills, QuoteExchangeTypeStonex:
		return true
	}
	return false
}

type QuoteParam struct {
	// Base currency
	BaseCurrency param.Field[string] `json:"baseCurrency,required"`
	// Quote currency
	QuoteCurrency param.Field[string] `json:"quoteCurrency,required"`
	// Quote request ID
	QuoteRequestID param.Field[string] `json:"quoteRequestId,required" format:"uuid"`
	// deprecated, alias of createdAt, Create time of the quote
	Timestamp param.Field[int64] `json:"timestamp,required"`
	// deprecated, alias of expiredAtExpiration time of the quote
	ValidUntil param.Field[int64] `json:"validUntil,required"`
	// Ask price
	AskPrice param.Field[float64] `json:"askPrice"`
	// Ask quantity
	AskQuantity param.Field[float64] `json:"askQuantity"`
	// Bid price
	BidPrice param.Field[float64] `json:"bidPrice"`
	// Bid quantity
	BidQuantity param.Field[float64] `json:"bidQuantity"`
	// Create time of the quote
	CreatedAt param.Field[int64] `json:"createdAt"`
	// Exchange Account ID
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[QuoteExchangeType] `json:"exchangeType"`
	// Expiration time of the quote
	ExpiredAt param.Field[int64] `json:"expiredAt"`
}

func (r QuoteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QuoteRequest struct {
	// Base currency is the currency you want to buy or sell
	BaseCurrency string `json:"baseCurrency,required"`
	// Order side, BUY or SELL
	OrderSide string `json:"orderSide,required"`
	// Quote currency is the currency you want to pay or receive, and the price of the
	// base currency is quoted in the quote currency
	QuoteCurrency string `json:"quoteCurrency,required"`
	// The identifier for the exchange account
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Amount of the base currency
	Quantity float64 `json:"quantity"`
	// Amount of the quote currency
	QuoteQuantity float64          `json:"quoteQuantity"`
	JSON          quoteRequestJSON `json:"-"`
}

// quoteRequestJSON contains the JSON metadata for the struct [QuoteRequest]
type quoteRequestJSON struct {
	BaseCurrency      apijson.Field
	OrderSide         apijson.Field
	QuoteCurrency     apijson.Field
	ExchangeAccountID apijson.Field
	Quantity          apijson.Field
	QuoteQuantity     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QuoteRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quoteRequestJSON) RawJSON() string {
	return r.raw
}

type QuoteRequestParam struct {
	// Base currency is the currency you want to buy or sell
	BaseCurrency param.Field[string] `json:"baseCurrency,required"`
	// Order side, BUY or SELL
	OrderSide param.Field[string] `json:"orderSide,required"`
	// Quote currency is the currency you want to pay or receive, and the price of the
	// base currency is quoted in the quote currency
	QuoteCurrency param.Field[string] `json:"quoteCurrency,required"`
	// The identifier for the exchange account
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Amount of the base currency
	Quantity param.Field[float64] `json:"quantity"`
	// Amount of the quote currency
	QuoteQuantity param.Field[float64] `json:"quoteQuantity"`
}

func (r QuoteRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TradingQuotePostParams struct {
	QuoteRequest QuoteRequestParam `json:"quoteRequest,required"`
}

func (r TradingQuotePostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QuoteRequest)
}

type TradingQuoteRequestForQuoteParams struct {
	QuoteRequest QuoteRequestParam `json:"quoteRequest,required"`
}

func (r TradingQuoteRequestForQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QuoteRequest)
}
