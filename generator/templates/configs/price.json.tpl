{
  {{ if not .freqtrade.config.live }}
    "dry_run_wallet": {{ .freqtrade.wallet.amount }},
    "fee": {{ .freqtrade.wallet.fee }},
  {{ end }}
  "max_open_trades": {{ json .freqtrade.trade.max }},
  "stake_currency": "{{ .freqtrade.wallet.currency }}",
  "stake_amount": {{ json .freqtrade.trade.amount }},
  "tradable_balance_ratio": {{ ratio "100%" }},
  "fiat_display_currency": "{{ .freqtrade.wallet.fiat }}"
  {{ if .freqtrade.config.timeframe.value }}
  ,"timeframe": "{{ .freqtrade.config.timeframe.value }}"
  {{ end }}
}
