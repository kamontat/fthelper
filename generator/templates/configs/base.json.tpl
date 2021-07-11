{
  "dry_run": {{ not .freqtrade.config.live }},
  "cancel_open_orders_on_exit": false,
  "bot_name": "{{ join .freqtrade.name .common.suffix }}",
  "initial_state": "stopped",
  "forcebuy_enable": false,
  "internals": {
    "process_throttle_secs": 20,
    "heartbeat_interval": {{ toDuration "7m" "s" }}
  },
  "dataformat_ohlcv": "{{ .freqtrade.config.formatter }}",
  "dataformat_trades": "{{ .freqtrade.config.formatter }}",
{{ if eq .freqtrade.database "sqlite" }}
  "db_url": "sqlite:///{{ join .freqtrade.name .common.suffix }}.sqlite"
{{ else if eq .freqtrade.database "postgres" }}
  "db_url": "postgresql+psycopg2://{{ .secrets.postgres.username }}:{{ .secrets.postgres.password }}@{{ .postgres.name }}:{{ .postgres.port }}/{{ join .freqtrade.name .common.suffix }}"
{{ end }}
}
