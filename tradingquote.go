// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzalite

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/cadenza-lite-go/internal/apijson"
	"github.com/stainless-sdks/cadenza-lite-go/internal/param"
	"github.com/stainless-sdks/cadenza-lite-go/internal/requestconfig"
	"github.com/stainless-sdks/cadenza-lite-go/option"
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
func (r *TradingQuoteService) RequestForQuote(ctx context.Context, body TradingQuoteRequestForQuoteParams, opts ...option.RequestOption) (res *[]QuoteWithOrderCandidates, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/trading/fetchQuotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type QuoteWithOrderCandidates struct {
	// Base currency
	BaseCurrency string `json:"baseCurrency,required"`
	// Quote currency
	QuoteCurrency string `json:"quoteCurrency,required"`
	// Quote request ID
	QuoteRequestID string `json:"quoteRequestId,required" format:"uuid"`
	// Create time of the quote
	Timestamp int64 `json:"timestamp,required"`
	// Expiration time of the quote
	ValidUntil int64 `json:"validUntil,required"`
	// Ask price
	AskPrice float64 `json:"askPrice"`
	// Ask quantity
	AskQuantity float64 `json:"askQuantity"`
	// Bid price
	BidPrice float64 `json:"bidPrice"`
	// Bid quantity
	BidQuantity     float64                                  `json:"bidQuantity"`
	OrderCandidates []QuoteWithOrderCandidatesOrderCandidate `json:"orderCandidates"`
	JSON            quoteWithOrderCandidatesJSON             `json:"-"`
}

// quoteWithOrderCandidatesJSON contains the JSON metadata for the struct
// [QuoteWithOrderCandidates]
type quoteWithOrderCandidatesJSON struct {
	BaseCurrency    apijson.Field
	QuoteCurrency   apijson.Field
	QuoteRequestID  apijson.Field
	Timestamp       apijson.Field
	ValidUntil      apijson.Field
	AskPrice        apijson.Field
	AskQuantity     apijson.Field
	BidPrice        apijson.Field
	BidQuantity     apijson.Field
	OrderCandidates apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *QuoteWithOrderCandidates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quoteWithOrderCandidatesJSON) RawJSON() string {
	return r.raw
}

type QuoteWithOrderCandidatesOrderCandidate struct {
	// Exchange account ID
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Order side
	OrderSide QuoteWithOrderCandidatesOrderCandidatesOrderSide `json:"orderSide"`
	// Order type
	OrderType QuoteWithOrderCandidatesOrderCandidatesOrderType `json:"orderType"`
	// Quantity
	Quantity float64 `json:"quantity"`
	// Quote Quantity
	QuoteQuantity float64 `json:"quoteQuantity"`
	// Quote request ID
	QuoteRequestID string `json:"quoteRequestId" format:"uuid"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy QuoteWithOrderCandidatesOrderCandidatesRoutePolicy `json:"routePolicy"`
	// Symbol
	Symbol string                                     `json:"symbol"`
	JSON   quoteWithOrderCandidatesOrderCandidateJSON `json:"-"`
}

// quoteWithOrderCandidatesOrderCandidateJSON contains the JSON metadata for the
// struct [QuoteWithOrderCandidatesOrderCandidate]
type quoteWithOrderCandidatesOrderCandidateJSON struct {
	ExchangeAccountID apijson.Field
	OrderSide         apijson.Field
	OrderType         apijson.Field
	Quantity          apijson.Field
	QuoteQuantity     apijson.Field
	QuoteRequestID    apijson.Field
	RoutePolicy       apijson.Field
	Symbol            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *QuoteWithOrderCandidatesOrderCandidate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quoteWithOrderCandidatesOrderCandidateJSON) RawJSON() string {
	return r.raw
}

// Order side
type QuoteWithOrderCandidatesOrderCandidatesOrderSide string

const (
	QuoteWithOrderCandidatesOrderCandidatesOrderSideBuy  QuoteWithOrderCandidatesOrderCandidatesOrderSide = "BUY"
	QuoteWithOrderCandidatesOrderCandidatesOrderSideSell QuoteWithOrderCandidatesOrderCandidatesOrderSide = "SELL"
)

func (r QuoteWithOrderCandidatesOrderCandidatesOrderSide) IsKnown() bool {
	switch r {
	case QuoteWithOrderCandidatesOrderCandidatesOrderSideBuy, QuoteWithOrderCandidatesOrderCandidatesOrderSideSell:
		return true
	}
	return false
}

// Order type
type QuoteWithOrderCandidatesOrderCandidatesOrderType string

const (
	QuoteWithOrderCandidatesOrderCandidatesOrderTypeMarket          QuoteWithOrderCandidatesOrderCandidatesOrderType = "MARKET"
	QuoteWithOrderCandidatesOrderCandidatesOrderTypeLimit           QuoteWithOrderCandidatesOrderCandidatesOrderType = "LIMIT"
	QuoteWithOrderCandidatesOrderCandidatesOrderTypeStopLoss        QuoteWithOrderCandidatesOrderCandidatesOrderType = "STOP_LOSS"
	QuoteWithOrderCandidatesOrderCandidatesOrderTypeStopLossLimit   QuoteWithOrderCandidatesOrderCandidatesOrderType = "STOP_LOSS_LIMIT"
	QuoteWithOrderCandidatesOrderCandidatesOrderTypeTakeProfit      QuoteWithOrderCandidatesOrderCandidatesOrderType = "TAKE_PROFIT"
	QuoteWithOrderCandidatesOrderCandidatesOrderTypeTakeProfitLimit QuoteWithOrderCandidatesOrderCandidatesOrderType = "TAKE_PROFIT_LIMIT"
	QuoteWithOrderCandidatesOrderCandidatesOrderTypeQuoted          QuoteWithOrderCandidatesOrderCandidatesOrderType = "QUOTED"
)

func (r QuoteWithOrderCandidatesOrderCandidatesOrderType) IsKnown() bool {
	switch r {
	case QuoteWithOrderCandidatesOrderCandidatesOrderTypeMarket, QuoteWithOrderCandidatesOrderCandidatesOrderTypeLimit, QuoteWithOrderCandidatesOrderCandidatesOrderTypeStopLoss, QuoteWithOrderCandidatesOrderCandidatesOrderTypeStopLossLimit, QuoteWithOrderCandidatesOrderCandidatesOrderTypeTakeProfit, QuoteWithOrderCandidatesOrderCandidatesOrderTypeTakeProfitLimit, QuoteWithOrderCandidatesOrderCandidatesOrderTypeQuoted:
		return true
	}
	return false
}

// Route policy. For PRIORITY, the order request will be routed to the exchange
// account with the highest priority. For QUOTE, the system will execute the
// execution plan based on the quote. Order request with route policy QUOTE will
// only accept two parameters, quoteRequestId and priceSlippageTolerance
type QuoteWithOrderCandidatesOrderCandidatesRoutePolicy string

const (
	QuoteWithOrderCandidatesOrderCandidatesRoutePolicyPriority QuoteWithOrderCandidatesOrderCandidatesRoutePolicy = "PRIORITY"
	QuoteWithOrderCandidatesOrderCandidatesRoutePolicyQuote    QuoteWithOrderCandidatesOrderCandidatesRoutePolicy = "QUOTE"
)

func (r QuoteWithOrderCandidatesOrderCandidatesRoutePolicy) IsKnown() bool {
	switch r {
	case QuoteWithOrderCandidatesOrderCandidatesRoutePolicyPriority, QuoteWithOrderCandidatesOrderCandidatesRoutePolicyQuote:
		return true
	}
	return false
}

type QuoteWithOrderCandidatesParam struct {
	// Base currency
	BaseCurrency param.Field[string] `json:"baseCurrency,required"`
	// Quote currency
	QuoteCurrency param.Field[string] `json:"quoteCurrency,required"`
	// Quote request ID
	QuoteRequestID param.Field[string] `json:"quoteRequestId,required" format:"uuid"`
	// Create time of the quote
	Timestamp param.Field[int64] `json:"timestamp,required"`
	// Expiration time of the quote
	ValidUntil param.Field[int64] `json:"validUntil,required"`
	// Ask price
	AskPrice param.Field[float64] `json:"askPrice"`
	// Ask quantity
	AskQuantity param.Field[float64] `json:"askQuantity"`
	// Bid price
	BidPrice param.Field[float64] `json:"bidPrice"`
	// Bid quantity
	BidQuantity     param.Field[float64]                                       `json:"bidQuantity"`
	OrderCandidates param.Field[[]QuoteWithOrderCandidatesOrderCandidateParam] `json:"orderCandidates"`
}

func (r QuoteWithOrderCandidatesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r QuoteWithOrderCandidatesParam) implementsEventPayloadUnionParam() {}

type QuoteWithOrderCandidatesOrderCandidateParam struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Order side
	OrderSide param.Field[QuoteWithOrderCandidatesOrderCandidatesOrderSide] `json:"orderSide"`
	// Order type
	OrderType param.Field[QuoteWithOrderCandidatesOrderCandidatesOrderType] `json:"orderType"`
	// Quantity
	Quantity param.Field[float64] `json:"quantity"`
	// Quote Quantity
	QuoteQuantity param.Field[float64] `json:"quoteQuantity"`
	// Quote request ID
	QuoteRequestID param.Field[string] `json:"quoteRequestId" format:"uuid"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy param.Field[QuoteWithOrderCandidatesOrderCandidatesRoutePolicy] `json:"routePolicy"`
	// Symbol
	Symbol param.Field[string] `json:"symbol"`
}

func (r QuoteWithOrderCandidatesOrderCandidateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TradingQuoteRequestForQuoteParams struct {
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

func (r TradingQuoteRequestForQuoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
