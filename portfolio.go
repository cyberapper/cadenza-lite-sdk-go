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

// PortfolioService contains methods and other services that help with interacting
// with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortfolioService] method instead.
type PortfolioService struct {
	Options []option.RequestOption
}

// NewPortfolioService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPortfolioService(opts ...option.RequestOption) (r *PortfolioService) {
	r = &PortfolioService{}
	r.Options = opts
	return
}

// List Portfolio Summary
func (r *PortfolioService) List(ctx context.Context, query PortfolioListParams, opts ...option.RequestOption) (res *[]ExchangeAccountPortfolio, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/portfolio/listSummaries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List balances
func (r *PortfolioService) ListBalances(ctx context.Context, query PortfolioListBalancesParams, opts ...option.RequestOption) (res *[]ExchangeAccountBalance, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/portfolio/listBalances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List credit
func (r *PortfolioService) ListCredit(ctx context.Context, query PortfolioListCreditParams, opts ...option.RequestOption) (res *[]ExchangeAccountCredit, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/portfolio/listCredit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List positions
func (r *PortfolioService) ListPositions(ctx context.Context, query PortfolioListPositionsParams, opts ...option.RequestOption) (res *[]ExchangeAccountPosition, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/portfolio/listPositions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ExchangeAccountBalance struct {
	// List of balances
	Balances []ExchangeAccountBalanceBalance `json:"balances,required"`
	// Exchange account ID
	ExchangeAccountID string                     `json:"exchangeAccountId,required" format:"uuid"`
	JSON              exchangeAccountBalanceJSON `json:"-"`
}

// exchangeAccountBalanceJSON contains the JSON metadata for the struct
// [ExchangeAccountBalance]
type exchangeAccountBalanceJSON struct {
	Balances          apijson.Field
	ExchangeAccountID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ExchangeAccountBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountBalanceJSON) RawJSON() string {
	return r.raw
}

type ExchangeAccountBalanceBalance struct {
	// Asset
	Asset string `json:"asset,required"`
	// Borrowed balance from exchange
	Borrowed float64 `json:"borrowed,required"`
	// Free balance
	Free float64 `json:"free,required"`
	// Locked balance
	Locked float64 `json:"locked,required"`
	// Net Balance, net = total - borrowed
	Net float64 `json:"net,required"`
	// Total available balance
	Total float64                           `json:"total,required"`
	JSON  exchangeAccountBalanceBalanceJSON `json:"-"`
}

// exchangeAccountBalanceBalanceJSON contains the JSON metadata for the struct
// [ExchangeAccountBalanceBalance]
type exchangeAccountBalanceBalanceJSON struct {
	Asset       apijson.Field
	Borrowed    apijson.Field
	Free        apijson.Field
	Locked      apijson.Field
	Net         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeAccountBalanceBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountBalanceBalanceJSON) RawJSON() string {
	return r.raw
}

// Exchange Account Credit Info
type ExchangeAccountCredit struct {
	// Type of account (SPOT, MARGIN)
	AccountType ExchangeAccountCreditAccountType `json:"accountType"`
	// The amount of credit available to the account from the broker or exchange
	Credit            float64 `json:"credit"`
	Currency          string  `json:"currency"`
	ExchangeAccountID string  `json:"exchangeAccountId"`
	// Exchange type
	ExchangeType ExchangeAccountCreditExchangeType `json:"exchangeType"`
	// The maximum leverage the account have
	Leverage int64 `json:"leverage"`
	// The amount of collateral that the investor has deposited in the account to cover
	// potential losses
	Margin float64 `json:"margin"`
	// The rate between equity and margin requirement
	MarginLevel float64 `json:"marginLevel"`
	// The amount of money borrowed from the broker to purchase securities
	MarginLoan float64 `json:"marginLoan"`
	// The amount of collateral required to maintain the current positions
	MarginRequirement float64 `json:"marginRequirement"`
	// The rate to which the available margin is being utilized
	MarginUsage float64 `json:"marginUsage"`
	// The maximum value of risk exposure that the account can handle, set to manage
	// risk and avoid excessive exposure to market volatility
	MaxRiskExposure float64 `json:"maxRiskExposure"`
	// The total value of positions held in the account, indicating the level of market
	// exposure
	RiskExposure float64 `json:"riskExposure"`
	// The rate between risk exposure and max risk exposure
	RiskExposureRate float64                   `json:"riskExposureRate"`
	JSON             exchangeAccountCreditJSON `json:"-"`
}

// exchangeAccountCreditJSON contains the JSON metadata for the struct
// [ExchangeAccountCredit]
type exchangeAccountCreditJSON struct {
	AccountType       apijson.Field
	Credit            apijson.Field
	Currency          apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Leverage          apijson.Field
	Margin            apijson.Field
	MarginLevel       apijson.Field
	MarginLoan        apijson.Field
	MarginRequirement apijson.Field
	MarginUsage       apijson.Field
	MaxRiskExposure   apijson.Field
	RiskExposure      apijson.Field
	RiskExposureRate  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ExchangeAccountCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountCreditJSON) RawJSON() string {
	return r.raw
}

// Type of account (SPOT, MARGIN)
type ExchangeAccountCreditAccountType string

const (
	ExchangeAccountCreditAccountTypeSpot   ExchangeAccountCreditAccountType = "SPOT"
	ExchangeAccountCreditAccountTypeMargin ExchangeAccountCreditAccountType = "MARGIN"
)

func (r ExchangeAccountCreditAccountType) IsKnown() bool {
	switch r {
	case ExchangeAccountCreditAccountTypeSpot, ExchangeAccountCreditAccountTypeMargin:
		return true
	}
	return false
}

// Exchange type
type ExchangeAccountCreditExchangeType string

const (
	ExchangeAccountCreditExchangeTypeBinance       ExchangeAccountCreditExchangeType = "BINANCE"
	ExchangeAccountCreditExchangeTypeBinanceMargin ExchangeAccountCreditExchangeType = "BINANCE_MARGIN"
	ExchangeAccountCreditExchangeTypeB2C2          ExchangeAccountCreditExchangeType = "B2C2"
	ExchangeAccountCreditExchangeTypeWintermute    ExchangeAccountCreditExchangeType = "WINTERMUTE"
	ExchangeAccountCreditExchangeTypeBlockfills    ExchangeAccountCreditExchangeType = "BLOCKFILLS"
	ExchangeAccountCreditExchangeTypeStonex        ExchangeAccountCreditExchangeType = "STONEX"
)

func (r ExchangeAccountCreditExchangeType) IsKnown() bool {
	switch r {
	case ExchangeAccountCreditExchangeTypeBinance, ExchangeAccountCreditExchangeTypeBinanceMargin, ExchangeAccountCreditExchangeTypeB2C2, ExchangeAccountCreditExchangeTypeWintermute, ExchangeAccountCreditExchangeTypeBlockfills, ExchangeAccountCreditExchangeTypeStonex:
		return true
	}
	return false
}

// Exchange Account Credit Info
type ExchangeAccountCreditParam struct {
	// Type of account (SPOT, MARGIN)
	AccountType param.Field[ExchangeAccountCreditAccountType] `json:"accountType"`
	// The amount of credit available to the account from the broker or exchange
	Credit            param.Field[float64] `json:"credit"`
	Currency          param.Field[string]  `json:"currency"`
	ExchangeAccountID param.Field[string]  `json:"exchangeAccountId"`
	// Exchange type
	ExchangeType param.Field[ExchangeAccountCreditExchangeType] `json:"exchangeType"`
	// The maximum leverage the account have
	Leverage param.Field[int64] `json:"leverage"`
	// The amount of collateral that the investor has deposited in the account to cover
	// potential losses
	Margin param.Field[float64] `json:"margin"`
	// The rate between equity and margin requirement
	MarginLevel param.Field[float64] `json:"marginLevel"`
	// The amount of money borrowed from the broker to purchase securities
	MarginLoan param.Field[float64] `json:"marginLoan"`
	// The amount of collateral required to maintain the current positions
	MarginRequirement param.Field[float64] `json:"marginRequirement"`
	// The rate to which the available margin is being utilized
	MarginUsage param.Field[float64] `json:"marginUsage"`
	// The maximum value of risk exposure that the account can handle, set to manage
	// risk and avoid excessive exposure to market volatility
	MaxRiskExposure param.Field[float64] `json:"maxRiskExposure"`
	// The total value of positions held in the account, indicating the level of market
	// exposure
	RiskExposure param.Field[float64] `json:"riskExposure"`
	// The rate between risk exposure and max risk exposure
	RiskExposureRate param.Field[float64] `json:"riskExposureRate"`
}

func (r ExchangeAccountCreditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeAccountPortfolio struct {
	// Exchange Account Credit Info
	Credit ExchangeAccountCredit `json:"credit,required"`
	// The unique identifier for the account.
	ExchangeAccountID string `json:"exchangeAccountId,required" format:"uuid"`
	// Exchange type
	ExchangeType ExchangeAccountPortfolioExchangeType `json:"exchangeType,required"`
	// The timestamp when the portfolio information was updated.
	UpdatedAt int64                              `json:"updatedAt,required"`
	Balances  []ExchangeAccountPortfolioBalance  `json:"balances"`
	Positions []ExchangeAccountPortfolioPosition `json:"positions"`
	JSON      exchangeAccountPortfolioJSON       `json:"-"`
}

// exchangeAccountPortfolioJSON contains the JSON metadata for the struct
// [ExchangeAccountPortfolio]
type exchangeAccountPortfolioJSON struct {
	Credit            apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	UpdatedAt         apijson.Field
	Balances          apijson.Field
	Positions         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ExchangeAccountPortfolio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountPortfolioJSON) RawJSON() string {
	return r.raw
}

func (r ExchangeAccountPortfolio) implementsEventPayload() {}

// Exchange type
type ExchangeAccountPortfolioExchangeType string

const (
	ExchangeAccountPortfolioExchangeTypeBinance       ExchangeAccountPortfolioExchangeType = "BINANCE"
	ExchangeAccountPortfolioExchangeTypeBinanceMargin ExchangeAccountPortfolioExchangeType = "BINANCE_MARGIN"
	ExchangeAccountPortfolioExchangeTypeB2C2          ExchangeAccountPortfolioExchangeType = "B2C2"
	ExchangeAccountPortfolioExchangeTypeWintermute    ExchangeAccountPortfolioExchangeType = "WINTERMUTE"
	ExchangeAccountPortfolioExchangeTypeBlockfills    ExchangeAccountPortfolioExchangeType = "BLOCKFILLS"
	ExchangeAccountPortfolioExchangeTypeStonex        ExchangeAccountPortfolioExchangeType = "STONEX"
)

func (r ExchangeAccountPortfolioExchangeType) IsKnown() bool {
	switch r {
	case ExchangeAccountPortfolioExchangeTypeBinance, ExchangeAccountPortfolioExchangeTypeBinanceMargin, ExchangeAccountPortfolioExchangeTypeB2C2, ExchangeAccountPortfolioExchangeTypeWintermute, ExchangeAccountPortfolioExchangeTypeBlockfills, ExchangeAccountPortfolioExchangeTypeStonex:
		return true
	}
	return false
}

type ExchangeAccountPortfolioBalance struct {
	// Asset
	Asset string `json:"asset,required"`
	// Borrowed balance from exchange
	Borrowed float64 `json:"borrowed,required"`
	// Free balance
	Free float64 `json:"free,required"`
	// Locked balance
	Locked float64 `json:"locked,required"`
	// Net Balance, net = total - borrowed
	Net float64 `json:"net,required"`
	// Total available balance
	Total float64                             `json:"total,required"`
	JSON  exchangeAccountPortfolioBalanceJSON `json:"-"`
}

// exchangeAccountPortfolioBalanceJSON contains the JSON metadata for the struct
// [ExchangeAccountPortfolioBalance]
type exchangeAccountPortfolioBalanceJSON struct {
	Asset       apijson.Field
	Borrowed    apijson.Field
	Free        apijson.Field
	Locked      apijson.Field
	Net         apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeAccountPortfolioBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountPortfolioBalanceJSON) RawJSON() string {
	return r.raw
}

type ExchangeAccountPortfolioPosition struct {
	// Amount
	Amount float64 `json:"amount,required"`
	// Position side
	PositionSide ExchangeAccountPortfolioPositionsPositionSide `json:"positionSide,required"`
	// Status
	Status ExchangeAccountPortfolioPositionsStatus `json:"status,required"`
	// Symbol
	Symbol string `json:"symbol,required"`
	// Cost
	Cost float64 `json:"cost"`
	// Entry price
	EntryPrice float64                              `json:"entryPrice"`
	JSON       exchangeAccountPortfolioPositionJSON `json:"-"`
}

// exchangeAccountPortfolioPositionJSON contains the JSON metadata for the struct
// [ExchangeAccountPortfolioPosition]
type exchangeAccountPortfolioPositionJSON struct {
	Amount       apijson.Field
	PositionSide apijson.Field
	Status       apijson.Field
	Symbol       apijson.Field
	Cost         apijson.Field
	EntryPrice   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ExchangeAccountPortfolioPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountPortfolioPositionJSON) RawJSON() string {
	return r.raw
}

// Position side
type ExchangeAccountPortfolioPositionsPositionSide string

const (
	ExchangeAccountPortfolioPositionsPositionSideLong  ExchangeAccountPortfolioPositionsPositionSide = "LONG"
	ExchangeAccountPortfolioPositionsPositionSideShort ExchangeAccountPortfolioPositionsPositionSide = "SHORT"
)

func (r ExchangeAccountPortfolioPositionsPositionSide) IsKnown() bool {
	switch r {
	case ExchangeAccountPortfolioPositionsPositionSideLong, ExchangeAccountPortfolioPositionsPositionSideShort:
		return true
	}
	return false
}

// Status
type ExchangeAccountPortfolioPositionsStatus string

const (
	ExchangeAccountPortfolioPositionsStatusOpen ExchangeAccountPortfolioPositionsStatus = "OPEN"
)

func (r ExchangeAccountPortfolioPositionsStatus) IsKnown() bool {
	switch r {
	case ExchangeAccountPortfolioPositionsStatusOpen:
		return true
	}
	return false
}

type ExchangeAccountPortfolioParam struct {
	// Exchange Account Credit Info
	Credit param.Field[ExchangeAccountCreditParam] `json:"credit,required"`
	// The unique identifier for the account.
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId,required" format:"uuid"`
	// Exchange type
	ExchangeType param.Field[ExchangeAccountPortfolioExchangeType] `json:"exchangeType,required"`
	// The timestamp when the portfolio information was updated.
	UpdatedAt param.Field[int64]                                   `json:"updatedAt,required"`
	Balances  param.Field[[]ExchangeAccountPortfolioBalanceParam]  `json:"balances"`
	Positions param.Field[[]ExchangeAccountPortfolioPositionParam] `json:"positions"`
}

func (r ExchangeAccountPortfolioParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExchangeAccountPortfolioParam) implementsEventPayloadUnionParam() {}

type ExchangeAccountPortfolioBalanceParam struct {
	// Asset
	Asset param.Field[string] `json:"asset,required"`
	// Borrowed balance from exchange
	Borrowed param.Field[float64] `json:"borrowed,required"`
	// Free balance
	Free param.Field[float64] `json:"free,required"`
	// Locked balance
	Locked param.Field[float64] `json:"locked,required"`
	// Net Balance, net = total - borrowed
	Net param.Field[float64] `json:"net,required"`
	// Total available balance
	Total param.Field[float64] `json:"total,required"`
}

func (r ExchangeAccountPortfolioBalanceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeAccountPortfolioPositionParam struct {
	// Amount
	Amount param.Field[float64] `json:"amount,required"`
	// Position side
	PositionSide param.Field[ExchangeAccountPortfolioPositionsPositionSide] `json:"positionSide,required"`
	// Status
	Status param.Field[ExchangeAccountPortfolioPositionsStatus] `json:"status,required"`
	// Symbol
	Symbol param.Field[string] `json:"symbol,required"`
	// Cost
	Cost param.Field[float64] `json:"cost"`
	// Entry price
	EntryPrice param.Field[float64] `json:"entryPrice"`
}

func (r ExchangeAccountPortfolioPositionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeAccountPosition struct {
	// Exchange account ID
	ExchangeAccountID string `json:"exchangeAccountId" format:"uuid"`
	// List of positions
	Positions []ExchangeAccountPositionPosition `json:"positions"`
	JSON      exchangeAccountPositionJSON       `json:"-"`
}

// exchangeAccountPositionJSON contains the JSON metadata for the struct
// [ExchangeAccountPosition]
type exchangeAccountPositionJSON struct {
	ExchangeAccountID apijson.Field
	Positions         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ExchangeAccountPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountPositionJSON) RawJSON() string {
	return r.raw
}

// List of positions
type ExchangeAccountPositionPosition struct {
	// Amount
	Amount float64 `json:"amount,required"`
	// Position side
	PositionSide ExchangeAccountPositionPositionsPositionSide `json:"positionSide,required"`
	// Status
	Status ExchangeAccountPositionPositionsStatus `json:"status,required"`
	// Symbol
	Symbol string `json:"symbol,required"`
	// Cost
	Cost float64 `json:"cost"`
	// Entry price
	EntryPrice float64                             `json:"entryPrice"`
	JSON       exchangeAccountPositionPositionJSON `json:"-"`
}

// exchangeAccountPositionPositionJSON contains the JSON metadata for the struct
// [ExchangeAccountPositionPosition]
type exchangeAccountPositionPositionJSON struct {
	Amount       apijson.Field
	PositionSide apijson.Field
	Status       apijson.Field
	Symbol       apijson.Field
	Cost         apijson.Field
	EntryPrice   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ExchangeAccountPositionPosition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountPositionPositionJSON) RawJSON() string {
	return r.raw
}

// Position side
type ExchangeAccountPositionPositionsPositionSide string

const (
	ExchangeAccountPositionPositionsPositionSideLong  ExchangeAccountPositionPositionsPositionSide = "LONG"
	ExchangeAccountPositionPositionsPositionSideShort ExchangeAccountPositionPositionsPositionSide = "SHORT"
)

func (r ExchangeAccountPositionPositionsPositionSide) IsKnown() bool {
	switch r {
	case ExchangeAccountPositionPositionsPositionSideLong, ExchangeAccountPositionPositionsPositionSideShort:
		return true
	}
	return false
}

// Status
type ExchangeAccountPositionPositionsStatus string

const (
	ExchangeAccountPositionPositionsStatusOpen ExchangeAccountPositionPositionsStatus = "OPEN"
)

func (r ExchangeAccountPositionPositionsStatus) IsKnown() bool {
	switch r {
	case ExchangeAccountPositionPositionsStatusOpen:
		return true
	}
	return false
}

type PortfolioListParams struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `query:"exchangeAccountId"`
	// Hide small account
	HideEmptyValue param.Field[bool] `query:"hideEmptyValue"`
}

// URLQuery serializes [PortfolioListParams]'s query parameters as `url.Values`.
func (r PortfolioListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortfolioListBalancesParams struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `query:"exchangeAccountId"`
	// Hide small account
	HideEmptyValue param.Field[bool] `query:"hideEmptyValue"`
}

// URLQuery serializes [PortfolioListBalancesParams]'s query parameters as
// `url.Values`.
func (r PortfolioListBalancesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortfolioListCreditParams struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `query:"exchangeAccountId"`
	// Hide small account
	HideEmptyValue param.Field[bool] `query:"hideEmptyValue"`
}

// URLQuery serializes [PortfolioListCreditParams]'s query parameters as
// `url.Values`.
func (r PortfolioListCreditParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortfolioListPositionsParams struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `query:"exchangeAccountId"`
	// Hide small account
	HideEmptyValue param.Field[bool] `query:"hideEmptyValue"`
}

// URLQuery serializes [PortfolioListPositionsParams]'s query parameters as
// `url.Values`.
func (r PortfolioListPositionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
