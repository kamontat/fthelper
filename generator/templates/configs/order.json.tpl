{
  "order_types": {
    "buy": "limit",
    "sell": "limit",
    "emergencysell": "market",
    "forcesell": "limit",
    "forcebuy": "limit",
    "stoploss": "market",
    "stoploss_on_exchange": false,
    "stoploss_on_exchange_interval": {{ toDuration "5m" "s" }},
    "stoploss_on_exchange_limit_ratio": 0.99
  }
}
