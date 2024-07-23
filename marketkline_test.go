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

func TestMarketKlineGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Market.Kline.Get(context.TODO(), cadenzalite.MarketKlineGetParams{
		ExchangeType: cadenzalite.F(cadenzalite.MarketKlineGetParamsExchangeTypeBinance),
		Interval:     cadenzalite.F(cadenzalite.MarketKlineGetParamsInterval1s),
		Symbol:       cadenzalite.F("BTC/USDT"),
		EndTime:      cadenzalite.F(int64(1632933600000)),
		Limit:        cadenzalite.F(int64(100)),
		StartTime:    cadenzalite.F(int64(1622505600000)),
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
