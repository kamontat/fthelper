# FTMetric

## v4.4.0 -> v4.5.0 (1 Aug 2021)

- `job` no longer unique, use `cluster` instead
- remove `freqtrade_stat_win_duration`, `freqtrade_stat_draw_duration`, `freqtrade_stat_loss_duration`, use `freqtrade_stat_duration` with `type` label instead
- @deprecated --env-files, use `--env` instead
- @deprecated --config-dirs, use `--configs` instead

## v4.0.0 -> v4.4.0 (20 Jul 2021)

- change metric name `freqtrade_pair_performance_profit_pct` to `freqtrade_pair_profit_pct`
- change metric name `freqtrade_pair_performance_profit_abs` to `freqtrade_pair_profit_abs`
- add `fthelper_internal_warmup` for warmup success rate instead
- remove `fthelper_internal_warmup_error` (use `fthelper_internal_warmup`)

## v3.x.x -> v4.0.0 (11 Jul 2021)

- Change environment prefix from `FTP` to `FTH`
- Change several option name (checking from `ftmetric --help`)
- Change metric name from `ftpersonal_` to `fthelper_`
- Change docker registry from DockerHub to Github Packages
- Update image name from `kamontat/ftmetric` to `ghcr.io/kamontat/ftmetric`
- @deprecated --list-config, use `ftmetric config` instead
