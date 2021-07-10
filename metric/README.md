# FTMetric (part of FTHelper)

This is simple server for expose [freqtrade](https://freqtrade.io) to prometheus database. Previously, this is part of ftpersonal repository which is private. After several people on freqtrade community show interest on this project. I decide to open source it. Private repository on version 3 will be stopped publish now.

## Break change

1. All environment prefix change from `FTP` to `FTH` and `FTPC` to `FTC`
2. Minor change on option. More detail will be on CHANGELOG
3. Some option is changed. More detail on --help command
4. All metric with `ftpersonal_` name, changed to `fthelper_`
5. Docker image now been published to Github packages instead
    - change image name from `kamontat/ftmetric` to `ghcr.io/kamontat/ftmetric`

## Installation

I provide 2 ways to run server. Either by docker image or download compiled scripts.

### Docker

Docker images are published to Github [packages](https://github.com/kamontat/fthelper/pkgs/container/ftmetric).

Each version will publish 2 tag name `<version>` and `<version>-scratch` as well as dynamic version `latest` and `scratch`.

1. Normal version is based from `busybox`. It contains some default shell commands for debug and healthcheck.
2. Busybox version is based from `scratch`. It not contains anything, meaning you cannot do anything inside the container.

### Compiled scripts

You will found compiled script for your os in [Release](https://github.com/kamontat/fthelper/releases) tab.

## Setup

After install scripts in your machine. You need to configure freqtrade to connect and other relate settings. You can config application with following method. ftmetric will load configuration by following order `files > environment > argument`

### Directory

- ftmetric get base settings from [configs](./configs) directory.
- all files in directory must be `json`; otherwise, it will throw error or crash
- you can change directories name via `--configs <path>` option.
  - this option can use multiple time to load multiple directories
  - on each directory also support multiple json file, each json will be merge together
  - warning: this will disable loading from default configs directory

### Environment

- ftmetric support loading configuration from environment variable / **.env** files
- every environment must prefix with `FTH_` or `FTC_`
- you can list all possible configuration and name via `ftmetric config` command at `Environment` column
- by default, ftmetric will try to load data from `$pwd/.env` file and warning if not found
- you can change files path via `--envs <path>` option.
  - this option can use multiple time to load multiple file.
  - this load as override, meaning if you have multiple setting in same name, last one will be use.
  - you can disable default `.env` load via `--no-env` option.

### Arguments

- ftmetric support argument config as well
- argument must be formatted as `<key>=<value>` (e.g. `ftmetric data.internal=true`)
- listed configuration is from `ftmetric config` command at `Key` column

## Example commands

```bash
# show help message
ftmetric --help
# show current version
ftmetric --version
# list configable settings with optional `--data` and `--all`
# --data will show current value of each config
# --all will show all settings including internal
ftmetric config [--data] [--all]
```
