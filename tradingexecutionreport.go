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

// TradingExecutionReportService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTradingExecutionReportService] method instead.
type TradingExecutionReportService struct {
	Options []option.RequestOption
}

// NewTradingExecutionReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTradingExecutionReportService(opts ...option.RequestOption) (r *TradingExecutionReportService) {
	r = &TradingExecutionReportService{}
	r.Options = opts
	return
}

// Quote will give the best quote from all available exchange accounts
func (r *TradingExecutionReportService) List(ctx context.Context, query TradingExecutionReportListParams, opts ...option.RequestOption) (res *pagination.Offset[ExecutionReport], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/trading/listExecutionReports"
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

// Quote will give the best quote from all available exchange accounts
func (r *TradingExecutionReportService) ListAutoPaging(ctx context.Context, query TradingExecutionReportListParams, opts ...option.RequestOption) *pagination.OffsetAutoPager[ExecutionReport] {
	return pagination.NewOffsetAutoPager(r.List(ctx, query, opts...))
}

// Quote will give the best quote from all available exchange accounts
func (r *TradingExecutionReportService) GetQuoteExecutionReport(ctx context.Context, query TradingExecutionReportGetQuoteExecutionReportParams, opts ...option.RequestOption) (res *ExecutionReport, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/trading/getQuoteExecutionReport"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ExecutionReport struct {
	// Base currency
	BaseCurrency string `json:"baseCurrency,required"`
	// Cost, the total cost of the quote
	Cost float64 `json:"cost,required"`
	// Create time of the quote
	CreatedAt int64 `json:"createdAt,required"`
	// Filled quantity, the quantity of the base currency executed
	Filled float64 `json:"filled,required"`
	// Quote currency
	QuoteCurrency string `json:"quoteCurrency,required"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy ExecutionReportRoutePolicy `json:"routePolicy,required"`
	// Status of the quote execution, should only have SUBMITTED, ACCEPTED,
	// PARTIALLY_FILLED, FILLED, EXPIRED. the final status of the quote execution
	// should be either FILLED or EXPIRED
	Status ExecutionReportStatus `json:"status,required"`
	// Last updated time of the quote execution
	UpdatedAt int64 `json:"updatedAt,required"`
	// Execution Report ID
	ID string `json:"id" format:"uuid"`
	// List of executions to fulfill the order, the order status should only have
	// FILLED, REJECTED, or EXPIRED
	Executions []Order `json:"executions"`
	// Fees
	Fees  []ExecutionReportFee `json:"fees"`
	Order Order                `json:"order"`
	JSON  executionReportJSON  `json:"-"`
}

// executionReportJSON contains the JSON metadata for the struct [ExecutionReport]
type executionReportJSON struct {
	BaseCurrency  apijson.Field
	Cost          apijson.Field
	CreatedAt     apijson.Field
	Filled        apijson.Field
	QuoteCurrency apijson.Field
	RoutePolicy   apijson.Field
	Status        apijson.Field
	UpdatedAt     apijson.Field
	ID            apijson.Field
	Executions    apijson.Field
	Fees          apijson.Field
	Order         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ExecutionReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionReportJSON) RawJSON() string {
	return r.raw
}

// Route policy. For PRIORITY, the order request will be routed to the exchange
// account with the highest priority. For QUOTE, the system will execute the
// execution plan based on the quote. Order request with route policy QUOTE will
// only accept two parameters, quoteRequestId and priceSlippageTolerance
type ExecutionReportRoutePolicy string

const (
	ExecutionReportRoutePolicyPriority ExecutionReportRoutePolicy = "PRIORITY"
	ExecutionReportRoutePolicyQuote    ExecutionReportRoutePolicy = "QUOTE"
)

func (r ExecutionReportRoutePolicy) IsKnown() bool {
	switch r {
	case ExecutionReportRoutePolicyPriority, ExecutionReportRoutePolicyQuote:
		return true
	}
	return false
}

// Status of the quote execution, should only have SUBMITTED, ACCEPTED,
// PARTIALLY_FILLED, FILLED, EXPIRED. the final status of the quote execution
// should be either FILLED or EXPIRED
type ExecutionReportStatus string

const (
	ExecutionReportStatusSubmitted       ExecutionReportStatus = "SUBMITTED"
	ExecutionReportStatusAccepted        ExecutionReportStatus = "ACCEPTED"
	ExecutionReportStatusOpen            ExecutionReportStatus = "OPEN"
	ExecutionReportStatusPartiallyFilled ExecutionReportStatus = "PARTIALLY_FILLED"
	ExecutionReportStatusFilled          ExecutionReportStatus = "FILLED"
	ExecutionReportStatusCanceled        ExecutionReportStatus = "CANCELED"
	ExecutionReportStatusPendingCancel   ExecutionReportStatus = "PENDING_CANCEL"
	ExecutionReportStatusRejected        ExecutionReportStatus = "REJECTED"
	ExecutionReportStatusExpired         ExecutionReportStatus = "EXPIRED"
	ExecutionReportStatusRevoked         ExecutionReportStatus = "REVOKED"
)

func (r ExecutionReportStatus) IsKnown() bool {
	switch r {
	case ExecutionReportStatusSubmitted, ExecutionReportStatusAccepted, ExecutionReportStatusOpen, ExecutionReportStatusPartiallyFilled, ExecutionReportStatusFilled, ExecutionReportStatusCanceled, ExecutionReportStatusPendingCancel, ExecutionReportStatusRejected, ExecutionReportStatusExpired, ExecutionReportStatusRevoked:
		return true
	}
	return false
}

type ExecutionReportFee struct {
	// Asset
	Asset string `json:"asset"`
	// Quantity
	Quantity float64                `json:"quantity"`
	JSON     executionReportFeeJSON `json:"-"`
}

// executionReportFeeJSON contains the JSON metadata for the struct
// [ExecutionReportFee]
type executionReportFeeJSON struct {
	Asset       apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionReportFee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionReportFeeJSON) RawJSON() string {
	return r.raw
}

type ExecutionReportParam struct {
	// Base currency
	BaseCurrency param.Field[string] `json:"baseCurrency,required"`
	// Cost, the total cost of the quote
	Cost param.Field[float64] `json:"cost,required"`
	// Create time of the quote
	CreatedAt param.Field[int64] `json:"createdAt,required"`
	// Filled quantity, the quantity of the base currency executed
	Filled param.Field[float64] `json:"filled,required"`
	// Quote currency
	QuoteCurrency param.Field[string] `json:"quoteCurrency,required"`
	// Route policy. For PRIORITY, the order request will be routed to the exchange
	// account with the highest priority. For QUOTE, the system will execute the
	// execution plan based on the quote. Order request with route policy QUOTE will
	// only accept two parameters, quoteRequestId and priceSlippageTolerance
	RoutePolicy param.Field[ExecutionReportRoutePolicy] `json:"routePolicy,required"`
	// Status of the quote execution, should only have SUBMITTED, ACCEPTED,
	// PARTIALLY_FILLED, FILLED, EXPIRED. the final status of the quote execution
	// should be either FILLED or EXPIRED
	Status param.Field[ExecutionReportStatus] `json:"status,required"`
	// Last updated time of the quote execution
	UpdatedAt param.Field[int64] `json:"updatedAt,required"`
	// Execution Report ID
	ID param.Field[string] `json:"id" format:"uuid"`
	// List of executions to fulfill the order, the order status should only have
	// FILLED, REJECTED, or EXPIRED
	Executions param.Field[[]OrderParam] `json:"executions"`
	// Fees
	Fees  param.Field[[]ExecutionReportFeeParam] `json:"fees"`
	Order param.Field[OrderParam]                `json:"order"`
}

func (r ExecutionReportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExecutionReportFeeParam struct {
	// Asset
	Asset param.Field[string] `json:"asset"`
	// Quantity
	Quantity param.Field[float64] `json:"quantity"`
}

func (r ExecutionReportFeeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TradingExecutionReportListParams struct {
	// End time (in unix milliseconds)
	EndTime param.Field[int64] `query:"endTime"`
	// Limit the number of returned results.
	Limit param.Field[int64] `query:"limit"`
	// Offset of the returned results. Default: 0
	Offset param.Field[int64] `query:"offset"`
	// Quote Request ID
	QuoteRequestID param.Field[string] `query:"quoteRequestId"`
	// Start time (in unix milliseconds)
	StartTime param.Field[int64] `query:"startTime"`
}

// URLQuery serializes [TradingExecutionReportListParams]'s query parameters as
// `url.Values`.
func (r TradingExecutionReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TradingExecutionReportGetQuoteExecutionReportParams struct {
	// Quote Request ID
	QuoteRequestID param.Field[string] `query:"quoteRequestId"`
}

// URLQuery serializes [TradingExecutionReportGetQuoteExecutionReportParams]'s
// query parameters as `url.Values`.
func (r TradingExecutionReportGetQuoteExecutionReportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
