// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenzasdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cyberapper/cadenza-lite-sdk-go"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/testutil"
	"github.com/cyberapper/cadenza-lite-sdk-go/option"
)

func TestMarketKlineGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenzasdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Market.Kline.Get(context.TODO(), cadenzasdk.MarketKlineGetParams{
		ExchangeType: cadenzasdk.F(cadenzasdk.MarketKlineGetParamsExchangeTypeBinance),
		Interval:     cadenzasdk.F(cadenzasdk.MarketKlineGetParamsInterval1s),
		Symbol:       cadenzasdk.F("BTC/USDT"),
		EndTime:      cadenzasdk.F(int64(1632933600000)),
		Limit:        cadenzasdk.F(int64(100)),
		StartTime:    cadenzasdk.F(int64(1622505600000)),
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
