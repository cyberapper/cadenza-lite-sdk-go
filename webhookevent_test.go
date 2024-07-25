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

func TestWebhookEventDropCopyQuoteWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhook.Event.DropCopyQuote(context.TODO(), cadenzasdk.WebhookEventDropCopyQuoteParams{
		DropCopyQuote: cadenzasdk.DropCopyQuoteParam{
			EventID:   "eventId",
			EventType: "eventType",
			Source:    "source",
			Timestamp: int64(1632933600000),
			Payload: cadenzasdk.EventPayloadPlaceOrderRequestParam{
				QuoteRequestID:    cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				BaseCurrency:      "BTC",
				QuoteCurrency:     "USDT",
				AskPrice:          42859.990000,
				AskQuantity:       int64(1),
				BidPrice:          42859.710000,
				BidQuantity:       int64(1),
				Timestamp:         int64(1632933600000),
				ValidUntil:        int64(1632933600000),
				ExchangeAccountID: cadenzasdk.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ExchangeType:      "BINANCE",
			},
		},
	})
	if err != nil {
		var apierr *cadenzasdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
