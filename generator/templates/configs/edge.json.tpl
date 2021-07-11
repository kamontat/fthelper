{
  {{ if .freqtrade.edge.enabled }}
  "edge": {
    "enabled": {{ .freqtrade.edge.enabled }},
    "process_throttle_secs": {{ toDuration "6h" "s" }},
    "calculate_since_number_of_days": {{ .freqtrade.edge.caldays }},
    "allowed_risk": {{ ratio .freqtrade.edge.risk }},
    "stoploss_range_min": -{{ ratio .freqtrade.edge.stoploss.min }},
    "stoploss_range_max": -{{ ratio .freqtrade.edge.stoploss.max }},
    "stoploss_range_step": -{{ ratio .freqtrade.edge.stoploss.step }},
    "minimum_winrate": {{ ratio .freqtrade.edge.winrate }},
    "minimum_expectancy": {{ ratio .freqtrade.edge.expectancy }},
    "min_trade_number": 10,
    "max_trade_duration_minute": {{ dayToDuration 5 "m" }},
    "remove_pumps": false
  }
  {{ end }}
}
