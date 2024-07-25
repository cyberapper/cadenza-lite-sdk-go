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

// ExchangeAccountService contains methods and other services that help with
// interacting with the cadenza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExchangeAccountService] method instead.
type ExchangeAccountService struct {
	Options []option.RequestOption
}

// NewExchangeAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExchangeAccountService(opts ...option.RequestOption) (r *ExchangeAccountService) {
	r = &ExchangeAccountService{}
	r.Options = opts
	return
}

// Add exchange account
func (r *ExchangeAccountService) New(ctx context.Context, body ExchangeAccountNewParams, opts ...option.RequestOption) (res *ExchangeAccountNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/exchange/addExchangeAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update exchange account, now only support Binance account API key and secret
func (r *ExchangeAccountService) Update(ctx context.Context, body ExchangeAccountUpdateParams, opts ...option.RequestOption) (res *ExchangeAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/exchange/updateExchangeAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List exchange accounts
func (r *ExchangeAccountService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ExchangeAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/exchange/listExchangeAccounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove exchange account, now only support Binance account API key and secret
func (r *ExchangeAccountService) Remove(ctx context.Context, body ExchangeAccountRemoveParams, opts ...option.RequestOption) (res *ExchangeAccountRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/exchange/removeExchangeAccount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Set the priority of exchanges
func (r *ExchangeAccountService) SetExchangePriority(ctx context.Context, body ExchangeAccountSetExchangePriorityParams, opts ...option.RequestOption) (res *ExchangeAccountSetExchangePriorityResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/exchange/setExchangePriority"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExchangeAccount struct {
	// Type of account (SPOT, MARGIN)
	AccountType ExchangeAccountAccountType `json:"accountType,required"`
	// Environment of the exchange account
	Environment ExchangeAccountEnvironment `json:"environment,required"`
	// Exchange account ID
	ExchangeAccountID string `json:"exchangeAccountId,required" format:"uuid"`
	// Exchange type
	ExchangeType ExchangeAccountExchangeType `json:"exchangeType,required"`
	// Name of the exchange account
	Name string `json:"name,required"`
	// Status of the exchange account
	Status ExchangeAccountStatus `json:"status,required"`
	JSON   exchangeAccountJSON   `json:"-"`
}

// exchangeAccountJSON contains the JSON metadata for the struct [ExchangeAccount]
type exchangeAccountJSON struct {
	AccountType       apijson.Field
	Environment       apijson.Field
	ExchangeAccountID apijson.Field
	ExchangeType      apijson.Field
	Name              apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ExchangeAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountJSON) RawJSON() string {
	return r.raw
}

// Type of account (SPOT, MARGIN)
type ExchangeAccountAccountType string

const (
	ExchangeAccountAccountTypeSpot   ExchangeAccountAccountType = "SPOT"
	ExchangeAccountAccountTypeMargin ExchangeAccountAccountType = "MARGIN"
)

func (r ExchangeAccountAccountType) IsKnown() bool {
	switch r {
	case ExchangeAccountAccountTypeSpot, ExchangeAccountAccountTypeMargin:
		return true
	}
	return false
}

// Environment of the exchange account
type ExchangeAccountEnvironment string

const (
	ExchangeAccountEnvironmentReal    ExchangeAccountEnvironment = "REAL"
	ExchangeAccountEnvironmentSandbox ExchangeAccountEnvironment = "SANDBOX"
	ExchangeAccountEnvironmentPaper   ExchangeAccountEnvironment = "PAPER"
)

func (r ExchangeAccountEnvironment) IsKnown() bool {
	switch r {
	case ExchangeAccountEnvironmentReal, ExchangeAccountEnvironmentSandbox, ExchangeAccountEnvironmentPaper:
		return true
	}
	return false
}

// Exchange type
type ExchangeAccountExchangeType string

const (
	ExchangeAccountExchangeTypeBinance       ExchangeAccountExchangeType = "BINANCE"
	ExchangeAccountExchangeTypeBinanceMargin ExchangeAccountExchangeType = "BINANCE_MARGIN"
	ExchangeAccountExchangeTypeB2C2          ExchangeAccountExchangeType = "B2C2"
	ExchangeAccountExchangeTypeWintermute    ExchangeAccountExchangeType = "WINTERMUTE"
	ExchangeAccountExchangeTypeBlockfills    ExchangeAccountExchangeType = "BLOCKFILLS"
	ExchangeAccountExchangeTypeStonex        ExchangeAccountExchangeType = "STONEX"
)

func (r ExchangeAccountExchangeType) IsKnown() bool {
	switch r {
	case ExchangeAccountExchangeTypeBinance, ExchangeAccountExchangeTypeBinanceMargin, ExchangeAccountExchangeTypeB2C2, ExchangeAccountExchangeTypeWintermute, ExchangeAccountExchangeTypeBlockfills, ExchangeAccountExchangeTypeStonex:
		return true
	}
	return false
}

// Status of the exchange account
type ExchangeAccountStatus string

const (
	ExchangeAccountStatusActive     ExchangeAccountStatus = "ACTIVE"
	ExchangeAccountStatusAPIError   ExchangeAccountStatus = "API_ERROR"
	ExchangeAccountStatusInvalidAPI ExchangeAccountStatus = "INVALID_API"
	ExchangeAccountStatusAPIIssue   ExchangeAccountStatus = "API_ISSUE"
	ExchangeAccountStatusNotTrusted ExchangeAccountStatus = "NOT_TRUSTED"
	ExchangeAccountStatusDeleted    ExchangeAccountStatus = "DELETED"
)

func (r ExchangeAccountStatus) IsKnown() bool {
	switch r {
	case ExchangeAccountStatusActive, ExchangeAccountStatusAPIError, ExchangeAccountStatusInvalidAPI, ExchangeAccountStatusAPIIssue, ExchangeAccountStatusNotTrusted, ExchangeAccountStatusDeleted:
		return true
	}
	return false
}

type ExchangeAccountNewResponse struct {
	Data string                         `json:"data"`
	JSON exchangeAccountNewResponseJSON `json:"-"`
}

// exchangeAccountNewResponseJSON contains the JSON metadata for the struct
// [ExchangeAccountNewResponse]
type exchangeAccountNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeAccountNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountNewResponseJSON) RawJSON() string {
	return r.raw
}

type ExchangeAccountUpdateResponse struct {
	Data string                            `json:"data"`
	JSON exchangeAccountUpdateResponseJSON `json:"-"`
}

// exchangeAccountUpdateResponseJSON contains the JSON metadata for the struct
// [ExchangeAccountUpdateResponse]
type exchangeAccountUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeAccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ExchangeAccountRemoveResponse struct {
	Data string                            `json:"data"`
	JSON exchangeAccountRemoveResponseJSON `json:"-"`
}

// exchangeAccountRemoveResponseJSON contains the JSON metadata for the struct
// [ExchangeAccountRemoveResponse]
type exchangeAccountRemoveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeAccountRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type ExchangeAccountSetExchangePriorityResponse struct {
	Data string                                         `json:"data"`
	JSON exchangeAccountSetExchangePriorityResponseJSON `json:"-"`
}

// exchangeAccountSetExchangePriorityResponseJSON contains the JSON metadata for
// the struct [ExchangeAccountSetExchangePriorityResponse]
type exchangeAccountSetExchangePriorityResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExchangeAccountSetExchangePriorityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r exchangeAccountSetExchangePriorityResponseJSON) RawJSON() string {
	return r.raw
}

type ExchangeAccountNewParams struct {
	// API key
	APIKey param.Field[string] `json:"apiKey,required"`
	// API secret
	APISecret param.Field[string] `json:"apiSecret,required"`
	// Environment(0 - real, 1 - sandbox)
	Environment param.Field[ExchangeAccountNewParamsEnvironment] `json:"environment,required"`
	// Exchange account name, Available characters: a-z, A-Z, 0-9, \_, (space)
	ExchangeAccountName param.Field[string] `json:"exchangeAccountName,required"`
	// Exchange type
	ExchangeType param.Field[ExchangeAccountNewParamsExchangeType] `json:"exchangeType,required"`
}

func (r ExchangeAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Environment(0 - real, 1 - sandbox)
type ExchangeAccountNewParamsEnvironment int64

const (
	ExchangeAccountNewParamsEnvironment0 ExchangeAccountNewParamsEnvironment = 0
	ExchangeAccountNewParamsEnvironment1 ExchangeAccountNewParamsEnvironment = 1
)

func (r ExchangeAccountNewParamsEnvironment) IsKnown() bool {
	switch r {
	case ExchangeAccountNewParamsEnvironment0, ExchangeAccountNewParamsEnvironment1:
		return true
	}
	return false
}

// Exchange type
type ExchangeAccountNewParamsExchangeType string

const (
	ExchangeAccountNewParamsExchangeTypeBinance       ExchangeAccountNewParamsExchangeType = "BINANCE"
	ExchangeAccountNewParamsExchangeTypeBinanceMargin ExchangeAccountNewParamsExchangeType = "BINANCE_MARGIN"
	ExchangeAccountNewParamsExchangeTypeB2C2          ExchangeAccountNewParamsExchangeType = "B2C2"
	ExchangeAccountNewParamsExchangeTypeWintermute    ExchangeAccountNewParamsExchangeType = "WINTERMUTE"
	ExchangeAccountNewParamsExchangeTypeBlockfills    ExchangeAccountNewParamsExchangeType = "BLOCKFILLS"
	ExchangeAccountNewParamsExchangeTypeStonex        ExchangeAccountNewParamsExchangeType = "STONEX"
)

func (r ExchangeAccountNewParamsExchangeType) IsKnown() bool {
	switch r {
	case ExchangeAccountNewParamsExchangeTypeBinance, ExchangeAccountNewParamsExchangeTypeBinanceMargin, ExchangeAccountNewParamsExchangeTypeB2C2, ExchangeAccountNewParamsExchangeTypeWintermute, ExchangeAccountNewParamsExchangeTypeBlockfills, ExchangeAccountNewParamsExchangeTypeStonex:
		return true
	}
	return false
}

type ExchangeAccountUpdateParams struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId,required" format:"uuid"`
	// API key
	APIKey param.Field[string] `json:"apiKey"`
	// API secret
	APISecret param.Field[string] `json:"apiSecret"`
	// Exchange account name, Available characters: a-z, A-Z, 0-9, \_, (space)
	ExchangeAccountName param.Field[string] `json:"exchangeAccountName"`
}

func (r ExchangeAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeAccountRemoveParams struct {
	// Exchange account ID
	ExchangeAccountID param.Field[string] `json:"exchangeAccountId,required" format:"uuid"`
}

func (r ExchangeAccountRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExchangeAccountSetExchangePriorityParams struct {
	// Priority list of exchanges in descending order
	Priority param.Field[[]string] `json:"priority,required"`
}

func (r ExchangeAccountSetExchangePriorityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
