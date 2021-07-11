{
  "api_server": {
    "enabled": {{ .freqtrade.server.enabled }},
    "enable_openapi": {{ .freqtrade.server.openapi }},
    "listen_ip_address": "{{ .freqtrade.server.url }}",
    "listen_port": {{ .freqtrade.server.port }},
    "verbosity": "error",
    "jwt_secret_key": "{{ .secrets.freqtrade.server.jwt }}",
    "username": "{{ .secrets.freqtrade.server.username }}",
    "password": "{{ .secrets.freqtrade.server.password }}",
    "CORS_origins": [
      "localhost", 
      "127.0.0.1", 
      "0.0.0.0", 
      "{{ .server.domain }}",
      "{{ .server.domains.freqtrade }}"
    ]
  }
}
