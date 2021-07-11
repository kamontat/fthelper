{
  "unfilledtimeout": {
    "buy": 10,
    "sell": 30
  },
  "exchange": {
    "ccxt_config": { "enableRateLimit": true },
    "ccxt_async_config": {
      "enableRateLimit": true,
      "rateLimit": 200
    },
    {{ if .freqtrade.pair.static }}
      "pair_whitelist": {{ json .freqtrade.pair.whitelist }},
    {{ end }}
    "pair_blacklist": {{ json .freqtrade.pair.blacklist }},
    "name": "{{ .secrets.exchange.name }}",
    "key": "{{ .secrets.exchange.apikey }}",
    "secret": "{{ .secrets.exchange.secret }}"
  }
}
