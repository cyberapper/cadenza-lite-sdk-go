# Health

Methods:

- <code title="get /api/v2/health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#HealthService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Clients

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ClientInfoGetResponse">ClientInfoGetResponse</a>

Methods:

- <code title="get /api/v2/client/getInfo">client.Clients.Info.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ClientInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ClientInfoGetResponse">ClientInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExchangeAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccount">ExchangeAccount</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountNewResponse">ExchangeAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountUpdateResponse">ExchangeAccountUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountRemoveResponse">ExchangeAccountRemoveResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountSetExchangePriorityResponse">ExchangeAccountSetExchangePriorityResponse</a>

Methods:

- <code title="post /api/v2/exchange/addExchangeAccount">client.ExchangeAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountNewParams">ExchangeAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountNewResponse">ExchangeAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/exchange/updateExchangeAccount">client.ExchangeAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountUpdateParams">ExchangeAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountUpdateResponse">ExchangeAccountUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/exchange/listExchangeAccounts">client.ExchangeAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccount">ExchangeAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/exchange/removeExchangeAccount">client.ExchangeAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountRemoveParams">ExchangeAccountRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountRemoveResponse">ExchangeAccountRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/exchange/setExchangePriority">client.ExchangeAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountService.SetExchangePriority">SetExchangePriority</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountSetExchangePriorityParams">ExchangeAccountSetExchangePriorityParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountSetExchangePriorityResponse">ExchangeAccountSetExchangePriorityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Market

## Instrument

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Instrument">Instrument</a>

Methods:

- <code title="get /api/v2/market/listSymbolInfo">client.Market.Instrument.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketInstrumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketInstrumentListParams">MarketInstrumentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Instrument">Instrument</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ticker

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Ticker">Ticker</a>

Methods:

- <code title="get /api/v2/market/ticker">client.Market.Ticker.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketTickerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketTickerGetParams">MarketTickerGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Ticker">Ticker</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orderbook

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#OrderbookParam">OrderbookParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Orderbook">Orderbook</a>

Methods:

- <code title="get /api/v2/market/orderbook">client.Market.Orderbook.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketOrderbookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketOrderbookGetParams">MarketOrderbookGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Orderbook">Orderbook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Kline

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#OhlcvParam">OhlcvParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Ohlcv">Ohlcv</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketKlineGetResponse">MarketKlineGetResponse</a>

Methods:

- <code title="get /api/v2/market/kline">client.Market.Kline.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketKlineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketKlineGetParams">MarketKlineGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#MarketKlineGetResponse">MarketKlineGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Trading

## Order

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#OrderParam">OrderParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Order">Order</a>

Methods:

- <code title="post /api/v2/trading/placeOrder">client.Trading.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingOrderNewParams">TradingOrderNewParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/trading/listOrders">client.Trading.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingOrderListParams">TradingOrderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go/internal/pagination#Offset">Offset</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Order">Order</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/trading/cancelOrder">client.Trading.Order.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingOrderService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingOrderCancelParams">TradingOrderCancelParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Quote

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#QuoteWithOrderCandidatesParam">QuoteWithOrderCandidatesParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#QuoteWithOrderCandidates">QuoteWithOrderCandidates</a>

Methods:

- <code title="post /api/v2/trading/fetchQuotes">client.Trading.Quote.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingQuoteService.RequestForQuote">RequestForQuote</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingQuoteRequestForQuoteParams">TradingQuoteRequestForQuoteParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#QuoteWithOrderCandidates">QuoteWithOrderCandidates</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ExecutionReport

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExecutionReportParam">ExecutionReportParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExecutionReport">ExecutionReport</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#QuoteExecutionReport">QuoteExecutionReport</a>

Methods:

- <code title="get /api/v2/trading/listExecutionReports">client.Trading.ExecutionReport.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingExecutionReportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingExecutionReportListParams">TradingExecutionReportListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExecutionReport">ExecutionReport</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/trading/getQuoteExecutionReport">client.Trading.ExecutionReport.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingExecutionReportService.GetQuoteExecutionReport">GetQuoteExecutionReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#TradingExecutionReportGetQuoteExecutionReportParams">TradingExecutionReportGetQuoteExecutionReportParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#QuoteExecutionReport">QuoteExecutionReport</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Portfolio

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountCreditParam">ExchangeAccountCreditParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountBalance">ExchangeAccountBalance</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountCredit">ExchangeAccountCredit</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountPosition">ExchangeAccountPosition</a>

Methods:

- <code title="get /api/v2/portfolio/listBalances">client.Portfolio.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#PortfolioService.ListBalances">ListBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#PortfolioListBalancesParams">PortfolioListBalancesParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountBalance">ExchangeAccountBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/portfolio/listCredit">client.Portfolio.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#PortfolioService.ListCredit">ListCredit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#PortfolioListCreditParams">PortfolioListCreditParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountCredit">ExchangeAccountCredit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/portfolio/listPositions">client.Portfolio.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#PortfolioService.ListPositions">ListPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#PortfolioListPositionsParams">PortfolioListPositionsParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#ExchangeAccountPosition">ExchangeAccountPosition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhook

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#EventParam">EventParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#WebhookPubsubResponse">WebhookPubsubResponse</a>

Methods:

- <code title="post /api/v2/webhook/pubsub">client.Webhook.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#WebhookService.Pubsub">Pubsub</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#WebhookPubsubParams">WebhookPubsubParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go">cadenzalite</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cadenza-lite-go#WebhookPubsubResponse">WebhookPubsubResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
