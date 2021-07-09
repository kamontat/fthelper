# FTMetric (part of FTHelper)

This is simple server for expose [freqtrade](https://freqtrade.io) to prometheus database. Previously, this is part of ftpersonal repository which is private. After several people on freqtrade community show interest on this project. I decide to open source it. Private repository on version 3 will be stopped publish now.

## Break change

1. All environment prefix change from `FTP` to `FTH` and `FTPC` to `FTC`
2. Minor change on option. More detail will be on CHANGELOG
3. Some option is changed. More detail on --help command
4. All metric with `ftpersonal_` name, changed to `fthelper_`
5. Docker image now been published to Github packages instead
    - change image name from `kamontat/ftmetric` to `ghcr.io/kamontat/ftmetric`

## Usage

I provide 2 ways to run server. Either by docker image or download compiled package.

### Docker

Docker images are published to Github [packages](https://github.com/kamontat/fthelper/pkgs/container/ftmetric).

Each version will publish 2 tag name `<version>` and `<version>-scratch` as well as dynamic version `latest` and `scratch`.

1. Normal version is based from `busybox`. It contains some default shell commands for debug and healthcheck.
2. Busybox version is based from `scratch`. It not contains anything, meaning you cannot do anything inside the container.
