# FTMetric (part of FTHelper)

This is simple server for expose [freqtrade](https://freqtrade.io) to prometheus database. Previously, this is part of fthelper repository which is private. After several people on freqtrade community show interest on this project. I decide to open source it. Private repository on version 3 will be stopped publish now.

## Break change

[Migration plan](./MIGRATION.md)

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
  - you can disable default `.env` load via `--no-env-file` option.

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

## Setup multiclusters

I centralize code between ftgenerator and ftmetric to use same logic for clustering.
So since version `5.0.0.beta.3`, you require to change how to set cluster configuration.

Normally, all fthelper module already support custom configuration via custom (`_`) fields (more information is [here](../shared/configs/README.md)).

Below is a steps for setup multiple clusters mode

1. Setup how many clusters (and name) you want
   1. Set by config file: field name `clusters`
   2. Set by environment: field name `FTH_CLUSTERS`
   3. Set by arguments: `--clusters <name>`
2. Setup configuration for specify cluster
   1. Set by config file: `{"_": {"1A": {...}}}`
   2. Set by environment: `FTC_1A__FREQTRADE__HTTP__URL=http://localhost:8081`
   3. Set by arguments: `_.1A.freqtrade.http.url=http://localhost:8081`

You will find default value of config file [here](./configs/common.json) and example value for environment [here](./.env.default)

## Example

1. example config - [./configs/common.json](./configs/common.json)
2. example environment value - [./.env.example](./.env.example)
3. example docker-compose - [./example/docker-compose.yml](./example/docker-compose.yml)
