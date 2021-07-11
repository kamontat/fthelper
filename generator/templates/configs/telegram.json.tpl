{
  {{ if .freqtrade.telegram.enabled }}
    "telegram": {
      "enabled": true,
      "balance_dust_level": 0.0001,
      "token": "{{ .secrets.freqtrade.telegram.token }}",
      "chat_id": "{{ .secrets.freqtrade.telegram.chatid }}",
      "keyboard": {{ json .freqtrade.telegram.dashboard }},
      "notification_settings": {{ json .freqtrade.telegram.notification }}
    }
  {{ else }}
    "telegram": {
      "enabled": false,
      "token": "",
      "chat_id": "",
      "keyboard": []
    }
  {{ end }}
}
