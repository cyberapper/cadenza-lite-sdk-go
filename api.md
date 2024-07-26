# Utility

Methods:

- <code title="get /api/v2/health">client.Utility.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#UtilityService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExchangeAccount

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccount">ExchangeAccount</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountNewResponse">ExchangeAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountUpdateResponse">ExchangeAccountUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountRemoveResponse">ExchangeAccountRemoveResponse</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountSetExchangePriorityResponse">ExchangeAccountSetExchangePriorityResponse</a>

Methods:

- <code title="post /api/v2/exchange/addExchangeAccount">client.ExchangeAccount.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountNewParams">ExchangeAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountNewResponse">ExchangeAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/exchange/updateExchangeAccount">client.ExchangeAccount.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountUpdateParams">ExchangeAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountUpdateResponse">ExchangeAccountUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/exchange/listExchangeAccounts">client.ExchangeAccount.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccount">ExchangeAccount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/exchange/removeExchangeAccount">client.ExchangeAccount.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountRemoveParams">ExchangeAccountRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountRemoveResponse">ExchangeAccountRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/exchange/setExchangePriority">client.ExchangeAccount.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountService.SetExchangePriority">SetExchangePriority</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountSetExchangePriorityParams">ExchangeAccountSetExchangePriorityParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountSetExchangePriorityResponse">ExchangeAccountSetExchangePriorityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Market

## Instrument

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Instrument">Instrument</a>

Methods:

- <code title="get /api/v2/market/listSymbolInfo">client.Market.Instrument.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketInstrumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketInstrumentListParams">MarketInstrumentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Instrument">Instrument</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ticker

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Ticker">Ticker</a>

Methods:

- <code title="get /api/v2/market/ticker">client.Market.Ticker.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketTickerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketTickerGetParams">MarketTickerGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Ticker">Ticker</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orderbook

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#OrderbookParam">OrderbookParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Orderbook">Orderbook</a>

Methods:

- <code title="get /api/v2/market/orderbook">client.Market.Orderbook.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketOrderbookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketOrderbookGetParams">MarketOrderbookGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Orderbook">Orderbook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Kline

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#CandlesParam">CandlesParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#KlineParam">KlineParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Candles">Candles</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Kline">Kline</a>

Methods:

- <code title="get /api/v2/market/kline">client.Market.Kline.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketKlineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketKlineGetParams">MarketKlineGetParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Kline">Kline</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Trading

## Order

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#CancelOrderRequestParam">CancelOrderRequestParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#OrderParam">OrderParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PlaceOrderRequestParam">PlaceOrderRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Order">Order</a>

Methods:

- <code title="post /api/v2/trading/placeOrder">client.Trading.Order.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingOrderNewParams">TradingOrderNewParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/trading/listOrders">client.Trading.Order.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingOrderListParams">TradingOrderListParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go/internal/pagination#Offset">Offset</a>[<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Order">Order</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/trading/cancelOrder">client.Trading.Order.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingOrderService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingOrderCancelParams">TradingOrderCancelParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Quote

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#QuoteParam">QuoteParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#QuoteRequestParam">QuoteRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Quote">Quote</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#QuoteRequest">QuoteRequest</a>

Methods:

- <code title="post /api/v2/trading/fetchQuotes">client.Trading.Quote.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingQuoteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingQuoteGetParams">TradingQuoteGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Quote">Quote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ExecutionReport

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExecutionReportParam">ExecutionReportParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExecutionReport">ExecutionReport</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#QuoteExecutionReport">QuoteExecutionReport</a>

Methods:

- <code title="get /api/v2/trading/listExecutionReports">client.Trading.ExecutionReport.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingExecutionReportService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingExecutionReportListParams">TradingExecutionReportListParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExecutionReport">ExecutionReport</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/trading/getQuoteExecutionReport">client.Trading.ExecutionReport.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingExecutionReportService.GetQuoteExecutionReport">GetQuoteExecutionReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TradingExecutionReportGetQuoteExecutionReportParams">TradingExecutionReportGetQuoteExecutionReportParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#QuoteExecutionReport">QuoteExecutionReport</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Portfolio

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#BalanceEntryParam">BalanceEntryParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountCreditParam">ExchangeAccountCreditParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountPortfolioParam">ExchangeAccountPortfolioParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PositionEntryParam">PositionEntryParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#BalanceEntry">BalanceEntry</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountBalance">ExchangeAccountBalance</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountCredit">ExchangeAccountCredit</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountPortfolio">ExchangeAccountPortfolio</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountPosition">ExchangeAccountPosition</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PositionEntry">PositionEntry</a>

Methods:

- <code title="get /api/v2/portfolio/listSummaries">client.Portfolio.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioListParams">PortfolioListParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountPortfolio">ExchangeAccountPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/portfolio/listBalances">client.Portfolio.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioService.ListBalances">ListBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioListBalancesParams">PortfolioListBalancesParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountBalance">ExchangeAccountBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/portfolio/listCredit">client.Portfolio.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioService.ListCredit">ListCredit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioListCreditParams">PortfolioListCreditParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountCredit">ExchangeAccountCredit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v2/portfolio/listPositions">client.Portfolio.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioService.ListPositions">ListPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#PortfolioListPositionsParams">PortfolioListPositionsParams</a>) ([]<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#ExchangeAccountPosition">ExchangeAccountPosition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhook

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#WebhookPubsubResponse">WebhookPubsubResponse</a>

Methods:

- <code title="post /api/v2/webhook/pubsub">client.Webhook.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#WebhookService.Pubsub">Pubsub</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#WebhookPubsubParams">WebhookPubsubParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#WebhookPubsubResponse">WebhookPubsubResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CloudScheduler

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#WebhookCloudSchedulerUpdatePortfolioRoutineResponse">WebhookCloudSchedulerUpdatePortfolioRoutineResponse</a>

Methods:

- <code title="post /api/v2/webhook/cloudScheduler/updatePortfolioRoutine">client.Webhook.CloudScheduler.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#WebhookCloudSchedulerService.UpdatePortfolioRoutine">UpdatePortfolioRoutine</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#WebhookCloudSchedulerUpdatePortfolioRoutineResponse">WebhookCloudSchedulerUpdatePortfolioRoutineResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Event

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventParam">EventParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Event">Event</a>

Methods:

- <code title="post /api/v2/webhook/pubsub/event">client.Event.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventNewParams">EventNewParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#Event">Event</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Task

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TaskQuoteParam">TaskQuoteParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TaskQuote">TaskQuote</a>

Methods:

- <code title="post /api/v2/webhook/pubsub/task/quote">client.Event.Task.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventTaskService.TaskQuote">TaskQuote</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventTaskTaskQuoteParams">EventTaskTaskQuoteParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#TaskQuote">TaskQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DropCopy

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyExecutionReportParam">DropCopyExecutionReportParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyOrderParam">DropCopyOrderParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyPortfolioParam">DropCopyPortfolioParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyQuoteParam">DropCopyQuoteParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyExecutionReport">DropCopyExecutionReport</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyOrder">DropCopyOrder</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyPortfolio">DropCopyPortfolio</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyQuote">DropCopyQuote</a>

Methods:

- <code title="post /api/v2/webhook/pubsub/dropCopy/executionReport">client.Event.DropCopy.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyService.DropCopyExecutionReport">DropCopyExecutionReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyDropCopyExecutionReportParams">EventDropCopyDropCopyExecutionReportParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyExecutionReport">DropCopyExecutionReport</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/webhook/pubsub/dropCopy/order">client.Event.DropCopy.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyService.DropCopyOrder">DropCopyOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyDropCopyOrderParams">EventDropCopyDropCopyOrderParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyOrder">DropCopyOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/webhook/pubsub/dropCopy/portfolio">client.Event.DropCopy.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyService.DropCopyPortfolio">DropCopyPortfolio</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyDropCopyPortfolioParams">EventDropCopyDropCopyPortfolioParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyPortfolio">DropCopyPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/webhook/pubsub/dropCopy/quote">client.Event.DropCopy.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyService.DropCopyQuote">DropCopyQuote</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventDropCopyDropCopyQuoteParams">EventDropCopyDropCopyQuoteParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#DropCopyQuote">DropCopyQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## MarketData

Params Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketDataKlineParam">MarketDataKlineParam</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketDataOrderBookParam">MarketDataOrderBookParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketDataKline">MarketDataKline</a>
- <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketDataOrderBook">MarketDataOrderBook</a>

Methods:

- <code title="post /api/v2/webhook/pubsub/marketData/kline">client.Event.MarketData.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventMarketDataService.MarketDataKline">MarketDataKline</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventMarketDataMarketDataKlineParams">EventMarketDataMarketDataKlineParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketDataKline">MarketDataKline</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v2/webhook/pubsub/marketData/orderBook">client.Event.MarketData.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventMarketDataService.MarketDataOrderBook">MarketDataOrderBook</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#EventMarketDataMarketDataOrderBookParams">EventMarketDataMarketDataOrderBookParams</a>) (<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go">cadenzasdk</a>.<a href="https://pkg.go.dev/github.com/cyberapper/cadenza-lite-sdk-go#MarketDataOrderBook">MarketDataOrderBook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
