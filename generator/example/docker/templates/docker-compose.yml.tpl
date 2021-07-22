version: '3'
volumes:{{ if .docker.postgres.enabled }}
  {{ join .docker.postgres.name "data" }}:{{ end }}{{ if .docker.grafana.enabled }}
  {{ join .docker.grafana.name "data" }}:{{ end }}{{ if .docker.prometheus.enabled }}
  {{ join .docker.prometheus.name "data" }}:{{ end }}

services:
  {{- $root := . }}
  {{- range $val := .data.clusters }}
  {{ $root.docker.freqtrade.name }}{{ $val }}:
    image: {{ $root.docker.freqtrade.image.name }}:{{ $root.docker.freqtrade.image.tag }}{{ if $root.docker.postgres.enabled -}}_pg{{- end }}
    container_name: {{ $root.docker.freqtrade.name }}{{ $val }}
    restart: unless-stopped
    ports:
      - {{ byCluster $root $val "docker.freqtrade.port" }}:{{ byCluster $root $val "docker.freqtrade.port" }}
    {{- if $root.docker.postgres.enabled }}
    depends_on:
      - {{ $root.docker.postgres.name }}
    {{- end }}
    volumes:
      - ./user_data:/freqtrade/user_data
    command: >
      trade
        --logfile "/freqtrade/user_data/logs/{{ join $root.docker.freqtrade.name $val }}.log"
        --config "/freqtrade/user_data/{{ join "config" $val }}.json"
    healthcheck:
      test: ["CMD-SHELL", "python3 /freqtrade/scripts/rest_client.py --config /freqtrade/user_data/config-rest-{{ $val }}.json version | grep -q version"]
      interval: 30s
      timeout: 4s
      retries: 3
      start_period: 1m
  {{- if $root.docker.ftmetric.enabled }}
  {{ $root.docker.ftmetric.name }}{{ $val }}:
    image: {{ $root.docker.ftmetric.image.name }}:{{ if not (eq $root.docker.ftmetric.image.tag "") }}{{ $root.docker.ftmetric.image.tag }}{{ else }}{{ if eq $root.internal.meta.version "dev" }}latest{{ else }}v{{ $root.internal.meta.version }}{{ end }}{{ end }}
    container_name: {{ $root.docker.ftmetric.name }}{{ $val }}
    restart: on-failure
    environment:
      - FTH_FREQTRADE__URL=http://{{ $root.docker.freqtrade.name }}{{ $val }}:{{ byCluster $root $val "docker.freqtrade.port" }}
      - FTH_FREQTRADE__USERNAME={{ byCluster $root $val "secrets.freqtrade.username" }}
      - FTH_FREQTRADE__PASSWORD={{ byCluster $root $val "secrets.freqtrade.password" }}
      - FTH_FREQTRADE__CLUSTER={{ $val }}
      - FTH_SERVER__PORT={{ $root.docker.ftmetric.port }}
    depends_on:
      - {{ $root.docker.freqtrade.name }}{{ $val }}
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:{{ $root.docker.ftmetric.port }}/version"]
      interval: 20s
      timeout: 1s
      retries: 3
      start_period: 20s
  {{ end }}
  {{- end -}}
