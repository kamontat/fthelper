# FTMetric (part of FTHelper)

This is simple server for expose [freqtrade](https://freqtrade.io) to prometheus database. Previously, this is part of ftpersonal repository which is private. After several people on freqtrade community show interest on this project. I decide to open source it. I still cannot open source all mono repository contains on ftpersonal project. Private repository on version 3 will be stopped publish now

## Break change

1. All environment prefix change from `FTP` to `FTH` and `FTPC` to `FTC`
2. Minor change on option. More detail will be on CHANGELOG
3. Some option is changed. More detail on --help command
4. All metric with `ftpersonal_` name, changed to `fthelper_`
5. Docker image now been published to Github packages instead
    - change image name from `kamontat/ftmetric` to `ghcr.io/kamontat/ftmetric`
