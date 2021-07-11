{
  {{ if .freqtrade.strategy.enabled }}
  "strategy": "{{ .freqtrade.strategy.current }}",
  {{ end }}

  "bid_strategy": {
    "price_side": "bid",
    "ask_last_balance": 0.0,
    "use_order_book": {{ .freqtrade.trade.orderbook }},
    "order_book_top": 1,
    "check_depth_of_market": {
      "enabled": {{ .freqtrade.trade.checkmarket.enabled }},
      "bids_to_ask_delta": {{ .freqtrade.trade.checkmarket.delta }}
    }
  },

  "ask_strategy": {
    "price_side": "ask",
    "use_order_book": {{ .freqtrade.trade.orderbook }},
    "order_book_top": 1
  }
}
