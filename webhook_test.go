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

func TestWebhookPubsubWithOptionalParams(t *testing.T) {
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
	_, err := client.Webhook.Pubsub(context.TODO(), cadenzalite.WebhookPubsubParams{
		Event: cadenzalite.EventParam{
			EventID:   cadenzalite.F("eventId"),
			EventType: cadenzalite.F("eventType"),
			Source:    cadenzalite.F("source"),
			Timestamp: cadenzalite.F(int64(1632933600000)),
			Payload: cadenzalite.F[cadenzalite.EventPayloadUnionParam](cadenzalite.EventPayloadQuoteRequestParam{
				BaseCurrency:      cadenzalite.F("baseCurrency"),
				QuoteCurrency:     cadenzalite.F("quoteCurrency"),
				OrderSide:         cadenzalite.F("orderSide"),
				Quantity:          cadenzalite.F(0.000000),
				QuoteQuantity:     cadenzalite.F(0.000000),
				ExchangeAccountID: cadenzalite.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
		},
	})
	if err != nil {
		var apierr *cadenzalite.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
