{
  "pairlists": [
    {{ if .freqtrade.pair.static }}
        { "method": "StaticPairList" }
    {{ else }}
        {
          "method": "VolumePairList",
          "number_assets": {{ .freqtrade.pair.size }},
          "sort_key": "quoteVolume",
          "refresh_period": {{ toDuration "1h" "s" }}
        }
    {{ end }}{{ if .freqtrade.pair.filter.enabled }}

    {{ if .freqtrade.pair.filter.age.enabled }}
    ,{ "method": "AgeFilter", "min_days_listed": {{ .freqtrade.pair.filter.age.days }} }
    {{ end }}
    {{ if .freqtrade.pair.filter.precision.enabled }}
    ,{ "method": "PrecisionFilter" }
    {{ end }}
    {{ if .freqtrade.pair.filter.stability.enabled }}
    ,{
      "method": "RangeStabilityFilter",
      "lookback_days": {{ .freqtrade.pair.filter.stability.lookback }},
      "min_rate_of_change": {{ ratio .freqtrade.pair.filter.stability.changerate }},
      "refresh_period": {{ toDuration "1h" "s" }}
    }
    {{ end }}
    {{ if .freqtrade.pair.filter.volatility.enabled }}
    ,{
      "method": "VolatilityFilter",
      "lookback_days": {{ .freqtrade.pair.filter.volatility.lookback }},
      "min_volatility": {{ ratio .freqtrade.pair.filter.volatility.min }},
      "max_volatility": {{ ratio .freqtrade.pair.filter.volatility.max }},
      "refresh_period": {{ dayToDuration 1 "s" }}
    }
    {{ end }}
    {{ if eq .freqtrade.pair.sort "performance" }}
    ,{ "method": "PerformanceFilter" }
    {{ else if eq .freqtrade.pair.sort "shuffle" }}
    ,{ "method": "ShuffleFilter" }
    {{ end }}

    {{ end }}
  ]
}
