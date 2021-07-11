{
  {{ if .freqtrade.protection.enabled }}
  "protections": [
    {
      "method": "StoplossGuard",
      "lookback_period": {{ toDuration .freqtrade.protection.stoploss.lookback "m" }},
      "trade_limit": {{ .freqtrade.protection.stoploss.limit }},
      "stop_duration": {{ toDuration .freqtrade.protection.stoploss.stop "m" }}
    }

    {{ if .freqtrade.protection.cooldown.enabled }}
    ,{ 
      "method": "CooldownPeriod",
      "stop_duration": {{ toDuration .freqtrade.protection.cooldown.stop "m" }}
    }
    {{ end }}

    {{ if .freqtrade.protection.drawdown.enabled }}
    ,{
      "method": "MaxDrawdown",
      "lookback_period": {{ toDuration .freqtrade.protection.drawdown.lookback "m" }},
      "trade_limit": {{ .freqtrade.protection.drawdown.limit }},
      "stop_duration": {{ toDuration .freqtrade.protection.drawdown.stop "m" }},
      "max_allowed_drawdown": {{ ratio .freqtrade.protection.drawdown.drawdown }}
    }
    {{ end }}

    {{ if .freqtrade.protection.lowprofit.enabled }}
    ,{
      "method": "LowProfitPairs",
      "lookback_period": {{ toDuration .freqtrade.protection.lowprofit.lookback "m" }},
      "trade_limit": {{ .freqtrade.protection.lowprofit.limit }},
      "stop_duration": {{ toDuration .freqtrade.protection.lowprofit.stop "m" }},
      "required_profit": {{ ratio .freqtrade.protection.lowprofit.profit }}
    }
    {{ end }}
  ]
  {{ end }}
}
