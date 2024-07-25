// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cyberapper/cadenza-lite-sdk-go/internal/apijson"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/apiquery"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/pagination"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/param"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/requestconfig"
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
)

// TradingOrderService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTradingOrderService] method instead.
type TradingOrderService struct {
	Options []option.RequestOption
}

// NewTradingOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTradingOrderService(opts ...option.RequestOption) (r *TradingOrderService) {
	r = &TradingOrderService{}
	r.Options = opts
	return
}

// Place order
func (r *TradingOrderService) New(ctx context.Context, params TradingOrderNewParams, opts ...option.RequestOption) (res *[]Order, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/trading/placeOrder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List orders
func (r *TradingOrderService) List(ctx context.Context, query TradingOrderListParams, opts ...option.RequestOption) (res *pagination.Offset[Order], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/trading/listOrders"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List orders
func (r *TradingOrderService) ListAutoPaging(ctx context.Context, query TradingOrderListParams, opts ...option.RequestOption) *pagination.OffsetAutoPager[Order] {
	return pagination.NewOffsetAutoPager(r.List(ctx, query, opts...))
}

// Cancel order. If the order is already filled, it will return an error.
func (r *TradingOrderService) Cancel(ctx context.Context, body TradingOrderCancelParams, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/trading/cancelOrder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CancelOrderRequest struct {
	// Order ID
	OrderID string                 `json:"orderId,required"`
	JSON    cancelOrderRequestJSON `json:"-"`
}

// cancelOrderRequestJSON contains the JSON metadata for the struct
// [CancelOrderRequest]
type cancelOrderRequestJSON struct {
	OrderID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CancelOrderRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cancelOrderRequestJSON) RawJSON() string {
	return r.raw
}

func (r CancelOrderRequest) implementsGenericEventPayload() {}

type CancelOrderRequestParam struct {
	// Order ID
	OrderID param.Field[string] `json:"orderId,required"`
}

func (r CancelOrderRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CancelOrderRequestParam) implementsGenericEventPayloadUnionParam() {}

type Order struct {
	// The total cost of this order.
	Cost float64 `json:"cost,required"`
	// Created timestamp
	CreatedAt int64 `json:"createdAt,required"`
	// Exchange account ID
	ExchangeAccountID string `json:"exchangeAccountId,required" format:"uuid"`
	// Exchange type
	ExchangeType OrderExchangeType `json:"exchangeType,required"`
	// The quantity of this order that has been filled.
	Filled float64 `json:"filled,required"`
	// Order side
	OrderSide OrderOrderSide `json:"orderSide,required"`
	// Order type
	OrderType OrderOrderType `json:"orderType,required"`
	// Quantity
	Quantity float64 `json:"quantity,required"`
	// Order status
	Status OrderStatus `json:"status,required"`
	// Symbol
	Symbol string `json:"symbol,required"`
	// Time in force
	TimeInForce OrderTimeInForce `json:"timeInForce,required"`
	// Last updated timestamp
	UpdatedAt int64 `json:"updatedAt,required"`
	// Fee
	Fee float64 `json:"fee"`
	// Fee currency
	FeeCurrency string `json:"feeCurrency"`
	OrderID     string `json:"orderId" format:"uuid"`
	// Position ID
	PositionID string `json:"positionId" format:"uuid"`
	// Price
	Price float64 `json:"price"`
	// Quote Quantity
	QuoteQuantity float64 `json:"quoteQuantity"`
	// Tenant ID
	TenantID string `json:"tenantId"`
	// User ID
	UserID string    `json:"userId" format:"uuid"`
	JSON   orderJSON `json:"-"`
}

// orderJSON contains the JSON metadata for the struct [Order]
type orderJSON struct {
	Cost              apijson.Field
	CreatedAt         apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Filled            apijson.Field
	OrderSide         apijson.Field
	OrderType         apijson.Field
	Quantity          apijson.Field
	Status            apijson.Field
	Symbol            apijson.Field
	TimeInForce       apijson.Field
	UpdatedAt         apijson.Field
	Fee               apijson.Field
	FeeCurrency       apijson.Field
	OrderID           apijson.Field
	PositionID        apijson.Field
	Price             apijson.Field
	QuoteQuantity     apijson.Field
	TenantID          apijson.Field
	UserID            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Order) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderJSON) RawJSON() string {
	return r.raw
}

func (r Order) implementsGenericEventPayload() {}

// Exchange type
type OrderExchangeType string

const (
	OrderExchangeTypeBinance       OrderExchangeType = "BINANCE"
	OrderExchangeTypeBinanceMargin OrderExchangeType = "BINANCE_MARGIN"
	OrderExchangeTypeB2C2          OrderExchangeType = "B2C2"
	OrderExchangeTypeWintermute    OrderExchangeType = "WINTERMUTE"
	OrderExchangeTypeBlockfills    OrderExchangeType = "BLOCKFILLS"
	OrderExchangeTypeStonex        OrderExchangeType = "STONEX"
)

func (r OrderExchangeType) IsKnown() bool {
	switch r {
	case OrderExchangeTypeBinance, OrderExchangeTypeBinanceMargin, OrderExchangeTypeB2C2, OrderExchangeTypeWintermute, OrderExchangeTypeBlockfills, OrderExchangeTypeStonex:
		return true
	}
	return false
}

// Order side
type OrderOrderSide string

const (
	OrderOrderSideBuy  OrderOrderSide = "BUY"
	OrderOrderSideSell OrderOrderSide = "SELL"
)

func (r OrderOrderSide) IsKnown() bool {
	switch r {
	case OrderOrderSideBuy, OrderOrderSideSell:
		return true
	}
	return false
}

// Order type
type OrderOrderType string

const (
	OrderOrderTypeMarket          OrderOrderType = "MARKET"
	OrderOrderTypeLimit           OrderOrderType = "LIMIT"
	OrderOrderTypeStopLoss        OrderOrderType = "STOP_LOSS"
	OrderOrderTypeStopLossLimit   OrderOrderType = "STOP_LOSS_LIMIT"
	OrderOrderTypeTakeProfit      OrderOrderType = "TAKE_PROFIT"
	OrderOrderTypeTakeProfitLimit OrderOrderType = "TAKE_PROFIT_LIMIT"
	OrderOrderTypeQuoted          OrderOrderType = "QUOTED"
)

func (r OrderOrderType) IsKnown() bool {
	switch r {
	case OrderOrderTypeMarket, OrderOrderTypeLimit, OrderOrderTypeStopLoss, OrderOrderTypeStopLossLimit, OrderOrderTypeTakeProfit, OrderOrderTypeTakeProfitLimit, OrderOrderTypeQuoted:
		return true
	}
	return false
}

// Order status
type OrderStatus string

const (
	OrderStatusSubmitted       OrderStatus = "SUBMITTED"
	OrderStatusAccepted        OrderStatus = "ACCEPTED"
	OrderStatusOpen            OrderStatus = "OPEN"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusPendingCancel   OrderStatus = "PENDING_CANCEL"
	OrderStatusRejected        OrderStatus = "REJECTED"
	OrderStatusExpired         OrderStatus = "EXPIRED"
	OrderStatusRevoked         OrderStatus = "REVOKED"
)

func (r OrderStatus) IsKnown() bool {
	switch r {
	case OrderStatusSubmitted, OrderStatusAccepted, OrderStatusOpen, OrderStatusPartiallyFilled, OrderStatusFilled, OrderStatusCanceled, OrderStatusPendingCancel, OrderStatusRejected, OrderStatusExpired, OrderStatusRevoked:
		return true
	}
	return false
}

// Time in force
type OrderTimeInForce string

const (
	OrderTimeInForceDay OrderTimeInForce = "DAY"
	OrderTimeInForceGtc OrderTimeInForce = "GTC"
	OrderTimeInForceGtx OrderTimeInForce = "GTX"
	OrderTimeInForceGtd OrderTimeInForce = "GTD"
	OrderTimeInForceOpg OrderTimeInForce = "OPG"
	OrderTimeInForceCls OrderTimeInForce = "CLS"
	OrderTimeInForceIoc OrderTimeInForce = "IOC"
	OrderTimeInForceFok OrderTimeInForce = "FOK"
	OrderTimeInForceGfa OrderTimeInForce = "GFA"
	OrderTimeInForceGfs OrderTimeInForce = "GFS"
	OrderTimeInForceGtm OrderTimeInForce = "GTM"
	OrderTimeInForceMoo OrderTimeInForce = "MOO"
	OrderTimeInForceMoc OrderTimeInForce = "MOC"
	OrderTimeInForceExt OrderTimeInForce = "EXT"
)

func (r OrderTimeInForce) IsKnown() bool {
	switch r {
	case OrderTimeInForceDay, OrderTimeInForceGtc, OrderTimeInForceGtx, OrderTimeInForceGtd, OrderTimeInForceOpg, OrderTimeInForceCls, OrderTimeInForceIoc, OrderTimeInForceFok, OrderTimeInForceGfa, OrderTimeInForceGfs, OrderTimeInForceGtm, OrderTimeInForceMoo, OrderTimeInForceMoc, OrderTimeInForceExt:
		return true
	}
	return false
}

type OrderParam struct {
	// The total cost of this order.
	Cost param.Field[float64] `json:"cost,required"`
	// Created timestamp
	CreatedAt param.Field[int64] `json:"createdAt,required"`
	// Exchange account ID
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId,required" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[OrderExchangeType] `json:"exchangeType,required"`
	// The quantity of this order that has been filled.
	Filled param.Field[float64] `json:"filled,required"`
	// Order side
	OrderSide param.Field[OrderOrderSide] `json:"orderSide,required"`
	// Order type
	OrderType param.Field[OrderOrderType] `json:"orderType,required"`
	// Quantity
	Quantity param.Field[float64] `json:"quantity,required"`
	// Order status
	Status param.Field[OrderStatus] `json:"status,required"`
	// Symbol
	Symbol param.Field[string] `json:"symbol,required"`
	// Time in force
	TimeInForce param.Field[OrderTimeInForce] `json:"timeInForce,required"`
	// Last updated timestamp
	UpdatedAt param.Field[int64] `json:"updatedAt,required"`
	// Fee
	Fee param.Field[float64] `json:"fee"`
	// Fee currency
	FeeCurrency param.Field[string] `json:"feeCurrency"`
	OrderID     param.Field[string] `json:"orderId" format:"uuid"`
	// Position ID
	PositionID param.Field[string] `json:"positionId" format:"uuid"`
	// Price
	Price param.Field[float64] `json:"price"`
	// Quote Quantity
	QuoteQuantity param.Field[float64] `json:"quoteQuantity"`
	// Tenant ID
	TenantID param.Field[string] `json:"tenantId"`
	// User ID
	UserID param.Field[string] `json:"userId" format:"uuid"`
}

func (r OrderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrderParam) implementsGenericEventPayloadUnionParam() {}

type PlaceOrderRequest struct {
	// Exchange account ID
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// Levarage
	Leverage int64 `json:"leverage"`
	// Order side
	OrderSide PlaceOrderRequestOrderSide `json:"orderSide"`
	// Order type
	OrderType PlaceOrderRequestOrderType `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID string `json:"positionId" format:"uuid"`
	// Price
	Price float64 `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance float64 `json:"priceSlippageTolerance"`
	// Priority list of exchange account ID in descending order
	Priority []string `json:"priority"`
	// Quantity. One of quantity or quoteQuantity must be provided. If both is
	// provided, only quantity will be used.
	Quantity float64 `json:"quantity"`
	// Quote ID used by exchange for RFQ, e.g. WINTERMUTE need this field to execute
	// QUOTED order
	QuoteID string `json:"quoteId"`
	// Quote Quantity
	QuoteQuantity float64 `json:"quoteQuantity"`
	// Quote request ID
	QuoteRequestID string `json:"quoteRequestId" format:"uuid"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy PlaceOrderRequestRoutePolicy `json:"routePolicy"`
	// Symbol
	Symbol string `json:"symbol"`
	// Tenant ID
	TenantID string `json:"tenantId"`
	// Time in force
	TimeInForce PlaceOrderRequestTimeInForce `json:"timeInForce"`
	JSON        placeOrderRequestJSON        `json:"-"`
}

// placeOrderRequestJSON contains the JSON metadata for the struct
// [PlaceOrderRequest]
type placeOrderRequestJSON struct {
	ExchangeAccountID      apijson.Field
	Leverage               apijson.Field
	OrderSide              apijson.Field
	OrderType              apijson.Field
	PositionID             apijson.Field
	Price                  apijson.Field
	PriceSlippageTolerance apijson.Field
	Priority               apijson.Field
	Quantity               apijson.Field
	QuoteID                apijson.Field
	QuoteQuantity          apijson.Field
	QuoteRequestID         apijson.Field
	RoutePolicy            apijson.Field
	Symbol                 apijson.Field
	TenantID               apijson.Field
	TimeInForce            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PlaceOrderRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placeOrderRequestJSON) RawJSON() string {
	return r.raw
}

func (r PlaceOrderRequest) implementsGenericEventPayload() {}

// Order side
type PlaceOrderRequestOrderSide string

const (
	PlaceOrderRequestOrderSideBuy  PlaceOrderRequestOrderSide = "BUY"
	PlaceOrderRequestOrderSideSell PlaceOrderRequestOrderSide = "SELL"
)

func (r PlaceOrderRequestOrderSide) IsKnown() bool {
	switch r {
	case PlaceOrderRequestOrderSideBuy, PlaceOrderRequestOrderSideSell:
		return true
	}
	return false
}

// Order type
type PlaceOrderRequestOrderType string

const (
	PlaceOrderRequestOrderTypeMarket          PlaceOrderRequestOrderType = "MARKET"
	PlaceOrderRequestOrderTypeLimit           PlaceOrderRequestOrderType = "LIMIT"
	PlaceOrderRequestOrderTypeStopLoss        PlaceOrderRequestOrderType = "STOP_LOSS"
	PlaceOrderRequestOrderTypeStopLossLimit   PlaceOrderRequestOrderType = "STOP_LOSS_LIMIT"
	PlaceOrderRequestOrderTypeTakeProfit      PlaceOrderRequestOrderType = "TAKE_PROFIT"
	PlaceOrderRequestOrderTypeTakeProfitLimit PlaceOrderRequestOrderType = "TAKE_PROFIT_LIMIT"
	PlaceOrderRequestOrderTypeQuoted          PlaceOrderRequestOrderType = "QUOTED"
)

func (r PlaceOrderRequestOrderType) IsKnown() bool {
	switch r {
	case PlaceOrderRequestOrderTypeMarket, PlaceOrderRequestOrderTypeLimit, PlaceOrderRequestOrderTypeStopLoss, PlaceOrderRequestOrderTypeStopLossLimit, PlaceOrderRequestOrderTypeTakeProfit, PlaceOrderRequestOrderTypeTakeProfitLimit, PlaceOrderRequestOrderTypeQuoted:
		return true
	}
	return false
}

// Route policy. For PRIORITY, the order request will be routed to the exchange
// account with the highest priority. For QUOTE, the system will execute the
// execution plan based on the quote. Order request with route policy QUOTE will
// only accept two parameters, quoteRequestId and priceSlippageTolerance
type PlaceOrderRequestRoutePolicy string

const (
	PlaceOrderRequestRoutePolicyPriority PlaceOrderRequestRoutePolicy = "PRIORITY"
	PlaceOrderRequestRoutePolicyQuote    PlaceOrderRequestRoutePolicy = "QUOTE"
)

func (r PlaceOrderRequestRoutePolicy) IsKnown() bool {
	switch r {
	case PlaceOrderRequestRoutePolicyPriority, PlaceOrderRequestRoutePolicyQuote:
		return true
	}
	return false
}

// Time in force
type PlaceOrderRequestTimeInForce string

const (
	PlaceOrderRequestTimeInForceDay PlaceOrderRequestTimeInForce = "DAY"
	PlaceOrderRequestTimeInForceGtc PlaceOrderRequestTimeInForce = "GTC"
	PlaceOrderRequestTimeInForceGtx PlaceOrderRequestTimeInForce = "GTX"
	PlaceOrderRequestTimeInForceGtd PlaceOrderRequestTimeInForce = "GTD"
	PlaceOrderRequestTimeInForceOpg PlaceOrderRequestTimeInForce = "OPG"
	PlaceOrderRequestTimeInForceCls PlaceOrderRequestTimeInForce = "CLS"
	PlaceOrderRequestTimeInForceIoc PlaceOrderRequestTimeInForce = "IOC"
	PlaceOrderRequestTimeInForceFok PlaceOrderRequestTimeInForce = "FOK"
	PlaceOrderRequestTimeInForceGfa PlaceOrderRequestTimeInForce = "GFA"
	PlaceOrderRequestTimeInForceGfs PlaceOrderRequestTimeInForce = "GFS"
	PlaceOrderRequestTimeInForceGtm PlaceOrderRequestTimeInForce = "GTM"
	PlaceOrderRequestTimeInForceMoo PlaceOrderRequestTimeInForce = "MOO"
	PlaceOrderRequestTimeInForceMoc PlaceOrderRequestTimeInForce = "MOC"
	PlaceOrderRequestTimeInForceExt PlaceOrderRequestTimeInForce = "EXT"
)

func (r PlaceOrderRequestTimeInForce) IsKnown() bool {
	switch r {
	case PlaceOrderRequestTimeInForceDay, PlaceOrderRequestTimeInForceGtc, PlaceOrderRequestTimeInForceGtx, PlaceOrderRequestTimeInForceGtd, PlaceOrderRequestTimeInForceOpg, PlaceOrderRequestTimeInForceCls, PlaceOrderRequestTimeInForceIoc, PlaceOrderRequestTimeInForceFok, PlaceOrderRequestTimeInForceGfa, PlaceOrderRequestTimeInForceGfs, PlaceOrderRequestTimeInForceGtm, PlaceOrderRequestTimeInForceMoo, PlaceOrderRequestTimeInForceMoc, PlaceOrderRequestTimeInForceExt:
		return true
	}
	return false
}

type PlaceOrderRequestParam struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId" format:"uuid"`
	// Levarage
	Leverage param.Field[int64] `json:"leverage"`
	// Order side
	OrderSide param.Field[PlaceOrderRequestOrderSide] `json:"orderSide"`
	// Order type
	OrderType param.Field[PlaceOrderRequestOrderType] `json:"orderType"`
	// Position ID for closing position in margin trading
	PositionID param.Field[string] `json:"positionId" format:"uuid"`
	// Price
	Price param.Field[float64] `json:"price"`
	// Price slippage tolerance, range: [0, 0.1] with 2 decimal places
	PriceSlippageTolerance param.Field[float64] `json:"priceSlippageTolerance"`
	// Priority list of exchange account ID in descending order
	Priority param.Field[[]string] `json:"priority"`
	// Quantity. One of quantity or quoteQuantity must be provided. If both is
	// provided, only quantity will be used.
	Quantity param.Field[float64] `json:"quantity"`
	// Quote ID used by exchange for RFQ, e.g. WINTERMUTE need this field to execute
	// QUOTED order
	QuoteID param.Field[string] `json:"quoteId"`
	// Quote Quantity
	QuoteQuantity param.Field[float64] `json:"quoteQuantity"`
	// Quote request ID
	QuoteRequestID param.Field[string] `json:"quoteRequestId" format:"uuid"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy param.Field[PlaceOrderRequestRoutePolicy] `json:"routePolicy"`
	// Symbol
	Symbol param.Field[string] `json:"symbol"`
	// Tenant ID
	TenantID param.Field[string] `json:"tenantId"`
	// Time in force
	TimeInForce param.Field[PlaceOrderRequestTimeInForce] `json:"timeInForce"`
}

func (r PlaceOrderRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlaceOrderRequestParam) implementsGenericEventPayloadUnionParam() {}

type TradingOrderNewParams struct {
	PlaceOrderRequest PlaceOrderRequestParam `json:"placeOrderRequest,required"`
	IdempotencyKey    param.Field[string]    `header:"Idempotency-Key"`
}

func (r TradingOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PlaceOrderRequest)
}

type TradingOrderListParams struct {
	// End time (in unix milliseconds)
	EndTime param.Field[int64] `query:"endTime"`
	// Exchange account ID
	ExchangeAccountID param.Field[string] `query:"exchangeAccountId"`
	// Limit the number of returned results.
	Limit param.Field[int64] `query:"limit"`
	// Offset of the returned results. Default: 0
	Offset param.Field[int64] `query:"offset"`
	// Order ID
	OrderID param.Field[string] `query:"orderId"`
	// Order status
	OrderStatus param.Field[TradingOrderListParamsOrderStatus] `query:"orderStatus"`
	// Start time (in unix milliseconds)
	StartTime param.Field[int64] `query:"startTime"`
	// Symbol
	Symbol param.Field[string] `query:"symbol"`
	// Tenant ID
	TenantID param.Field[string] `query:"tenantId"`
}

// URLQuery serializes [TradingOrderListParams]'s query parameters as `url.Values`.
func (r TradingOrderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order status
type TradingOrderListParamsOrderStatus string

const (
	TradingOrderListParamsOrderStatusSubmitted       TradingOrderListParamsOrderStatus = "SUBMITTED"
	TradingOrderListParamsOrderStatusAccepted        TradingOrderListParamsOrderStatus = "ACCEPTED"
	TradingOrderListParamsOrderStatusOpen            TradingOrderListParamsOrderStatus = "OPEN"
	TradingOrderListParamsOrderStatusPartiallyFilled TradingOrderListParamsOrderStatus = "PARTIALLY_FILLED"
	TradingOrderListParamsOrderStatusFilled          TradingOrderListParamsOrderStatus = "FILLED"
	TradingOrderListParamsOrderStatusCanceled        TradingOrderListParamsOrderStatus = "CANCELED"
	TradingOrderListParamsOrderStatusPendingCancel   TradingOrderListParamsOrderStatus = "PENDING_CANCEL"
	TradingOrderListParamsOrderStatusRejected        TradingOrderListParamsOrderStatus = "REJECTED"
	TradingOrderListParamsOrderStatusExpired         TradingOrderListParamsOrderStatus = "EXPIRED"
	TradingOrderListParamsOrderStatusRevoked         TradingOrderListParamsOrderStatus = "REVOKED"
)

func (r TradingOrderListParamsOrderStatus) IsKnown() bool {
	switch r {
	case TradingOrderListParamsOrderStatusSubmitted, TradingOrderListParamsOrderStatusAccepted, TradingOrderListParamsOrderStatusOpen, TradingOrderListParamsOrderStatusPartiallyFilled, TradingOrderListParamsOrderStatusFilled, TradingOrderListParamsOrderStatusCanceled, TradingOrderListParamsOrderStatusPendingCancel, TradingOrderListParamsOrderStatusRejected, TradingOrderListParamsOrderStatusExpired, TradingOrderListParamsOrderStatusRevoked:
		return true
	}
	return false
}

type TradingOrderCancelParams struct {
	CancelOrderRequest CancelOrderRequestParam `json:"cancelOrderRequest,required"`
}

func (r TradingOrderCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CancelOrderRequest)
}
