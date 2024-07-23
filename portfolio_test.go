// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzalite_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/cadenza-lite-go"
	"github.com/stainless-sdks/cadenza-lite-go/internal/testutil"
	"github.com/stainless-sdks/cadenza-lite-go/option"
)

func TestPortfolioListBalancesWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenzalite.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Portfolio.ListBalances(context.TODO(), cadenzalite.PortfolioListBalancesParams{
		ExchangeAccountID: cadenzalite.F("exchangeAccountId"),
		HideEmptyValue:    cadenzalite.F(true),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortfolioListCreditWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenzalite.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Portfolio.ListCredit(context.TODO(), cadenzalite.PortfolioListCreditParams{
		ExchangeAccountID: cadenzalite.F("exchangeAccountId"),
		HideEmptyValue:    cadenzalite.F(true),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortfolioListPositionsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenzalite.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Portfolio.ListPositions(context.TODO(), cadenzalite.PortfolioListPositionsParams{
		ExchangeAccountID: cadenzalite.F("exchangeAccountId"),
		HideEmptyValue:    cadenzalite.F(true),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
