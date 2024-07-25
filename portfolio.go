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
	// Free balance
	Free float64 `json:"free,required"`
	// Locked balance
	Locked float64 `json:"locked,required"`
	// Total balance
	Total float64                           `json:"total,required"`
	JSON  exchangeAccountBalanceBalanceJSON `json:"-"`
}

// exchangeAccountBalanceBalanceJSON contains the JSON metadata for the struct
// [ExchangeAccountBalanceBalance]
type exchangeAccountBalanceBalanceJSON struct {
	Asset       apijson.Field
	Free        apijson.Field
	Locked      apijson.Field
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
