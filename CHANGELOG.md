# Release note of FTHelper

<a name="unreleased"></a>
## [Unreleased]

### 3. Bug Fixes
- **shared:** add more coverage on configs modules ([#34](https://github.com/kamontat/fthelper/issues/34))


<a name="v5.0.0-beta.14"></a>
## [v5.0.0-beta.14] - 2021-09-20

### 2. Performance Improvements
- **shared:** improve how cache join together with new caches.Join apis

### 3. Bug Fixes
- **metric:** http strict-mode panic always failed because non-pointer is passed
- **metric:** single mode didn't work properly because connection didn't create correctly


<a name="v5.0.0-beta.13"></a>
## [v5.0.0-beta.13] - 2021-09-20

### 1. Features
- **metric:** new 'strict-mode' to panic stop server if it cannot connect to freqtrade

### 2. Performance Improvements
- **gen:** support .cluster in fs template
- **metric:** remove annoy log when user disable warmup
- **shared:** add more information in aggreegate error message
- **shared:** improve merger log by reduce duplicate assign value
- **shared:** mask some of map:merger log to reduce security leak [#26](https://github.com/kamontat/fthelper/issues/26)
- **shared:** type and mode of fs will be case insensitive
- **shared:** improve runner to handle disabled in validator function
- **shared:** remove unused code in runner and make information as data model
- **shared:** avoid using same function on validator and executor
- **shared:** move summary name to constant

### 3. Bug Fixes
- **metric:** update example env to match with new environment variable
- **shared:** update test of error aggregator output
- **shared:** test fail on CI
- **shared:** remove unused code to reduce code size


<a name="v5.0.0-beta.12"></a>
## [v5.0.0-beta.12] - 2021-08-24

### 2. Performance Improvements
- **gen:** add disabled field to disable specify generator
- **shared:** refactor runners and reduce code size
- **shared:** add loggers.GetTable() for get table directly

### 3. Bug Fixes
- **shared:** table log didn't add space if no data


<a name="v5.0.0-beta.11"></a>
## [v5.0.0-beta.11] - 2021-08-22

### 3. Bug Fixes
- **shared:** copy directory didn't copy every file from directory


<a name="v5.0.0-beta.10"></a>
## [v5.0.0-beta.10] - 2021-08-22

### 1. Features
- **gen:** new generator 'bash' for execute shell script in generator

### 3. Bug Fixes
- **gen:** json and config generator also support single/multiple mode
- **gen:** fs should not force to be object
- **gen:** update schema to support string input/output


<a name="v5.0.0-beta.9"></a>
## [v5.0.0-beta.9] - 2021-08-22

### 3. Bug Fixes
- **gen:** compile error


<a name="v5.0.0-beta.8"></a>
## [v5.0.0-beta.8] - 2021-08-22

### 2. Performance Improvements
- **gen:** restore fs fields back to reduce some duplicate config


<a name="v5.0.0-beta.7"></a>
## [v5.0.0-beta.7] - 2021-08-22

### 3. Bug Fixes
- **metric:** remove unused struct code


<a name="v5.0.0-beta.6"></a>
## [v5.0.0-beta.6] - 2021-08-22

### 1. Features
- **gen:** not link fs data from generator to fs fields, but use directly
- **gen:** move fs.variable to variables instead
- **metric:** add new path /healthcheck for check connection with freqtrade

### 2. Performance Improvements
- **gen:** update json schema to support new config format
- **gen:** deprecate -fsvar and use --var instead
- **metric:** application will crash if freqtrade http is not available
- **shared:** fs.Build will use mapper instead of name


<a name="v5.0.0-beta.5"></a>
## [v5.0.0-beta.5] - 2021-08-19

### 1. Features
- **metric:** add new freqtrade_perf_* for calculated profit hourly, daily, and monthly

### 2. Performance Improvements
- **metric:** improve error message when user forget to setup cache duration
- **metric:** significate reduce debug log to left only what importance

### 3. Bug Fixes
- **shared:** cluster didn't return correct data


<a name="v5.0.0-beta.4"></a>
## [v5.0.0-beta.4] - 2021-08-19

### 3. Bug Fixes
- **metric:** some metric didn't include on previous version


<a name="v5.0.0-beta.3"></a>
## [v5.0.0-beta.3] - 2021-08-18

### 1. Features
- **metric:** update cluster config to match with generator module

### 2. Performance Improvements
- **shared:** update default clusters to be empty string
- **shared:** when --debug is present, we will force setup log level early
- **shared:** centralize underscroll override config to support upper and lower case

### 3. Bug Fixes
- **metric:** db initial should honor enabled flag


<a name="v5.0.0-beta.2"></a>
## [v5.0.0-beta.2] - 2021-08-18

### 2. Performance Improvements
- remove all deprecated from v4.x.x.
- **shared:** add support custom clusters in config itself "clusters" fields


<a name="v5.0.0-beta.1"></a>
## [v5.0.0-beta.1] - 2021-08-17

### 1. Features
- **metric:** add history data support
- **metric:** add database connection support
- **metric:** reimplement how metric connect to freqtrade apis and add database support

### 2. Performance Improvements
- **shared:** update logger to custom writer for log message
- **shared:** update runner to throw whole list of errors instead of only first one

### 3. Bug Fixes
- **docs:** readme didn't open correct link


<a name="v4.5.3"></a>
## [v4.5.3] - 2021-08-04

### 2. Performance Improvements
- **shared:** add logger level setup from environment `FTH_INTERNAL__LOG__LEVEL` support


<a name="v4.5.2"></a>
## [v4.5.2] - 2021-08-04

### 2. Performance Improvements
- **metric:** add cluster detail to log namespace for multicluster mode

### 3. Bug Fixes
- **metric:** metric not show when 1+ cluster is down
- **metric:** crash when http connection refuse
- **metric:** ftmetric return 5xx when one cluster is down


<a name="v4.5.1"></a>
## [v4.5.1] - 2021-08-03

### 2. Performance Improvements
- **metric:** add better handle when freqtrade return non-success response.


<a name="v4.5.0"></a>
## [v4.5.0] - 2021-08-01

### 2. Performance Improvements
- **docs:** add migration plan for v4.x.x

### 3. Bug Fixes
- **ci:** upload artifact so I will understand what happen on testing


<a name="v4.5.0-beta.1"></a>
## [v4.5.0-beta.1] - 2021-08-01

### 3. Bug Fixes
- **metric:** remove deadcode


<a name="v4.5.0-alpha.8"></a>
## [v4.5.0-alpha.8] - 2021-08-01

### 3. Bug Fixes
- **metric:** crash on previous version
- **metric:** ft_call didn't report correct number


<a name="v4.5.0-alpha.7"></a>
## [v4.5.0-alpha.7] - 2021-08-01

### 2. Performance Improvements
- **metric:** remove failure ft_call and add clusters for multicluster mode
- **metric:** support cache miss per cluster on multicluster mode


<a name="v4.5.0-alpha.6"></a>
## [v4.5.0-alpha.6] - 2021-08-01

### 3. Bug Fixes
- **shared:** joinArray separator should be comma


<a name="v4.5.0-alpha.5"></a>
## [v4.5.0-alpha.5] - 2021-08-01

### 3. Bug Fixes
- **gen:** deprecate message always print because docker fixed code
- **metric:** deprecate message always print because docker fixed code


<a name="v4.5.0-alpha.4"></a>
## [v4.5.0-alpha.4] - 2021-08-01

### 2. Performance Improvements
- **shared:** add `joinArray` to template for join array variable


<a name="v4.5.0-alpha.3"></a>
## [v4.5.0-alpha.3] - 2021-08-01

### 2. Performance Improvements
- **metric:** add info when using alpha feature


<a name="v4.5.0-alpha.2"></a>
## [v4.5.0-alpha.2] - 2021-08-01

### 1. Features
- **shared:** deprecated `--env-files` option, use `--envs` instead
- **shared:** add support auto fs type to query data from file-system and check the type

### 2. Performance Improvements
- **docs:** add metric for multiple cluster
- **shared:** reduce duplicate code in config
- **shared:** deprecated `--config-dirs` because we have auto type supported

### 3. Bug Fixes
- **metric:** whitelist didn't return correct information
- **shared:** formatted config wrapper correctly


<a name="v4.5.0-alpha.1"></a>
## [v4.5.0-alpha.1] - 2021-08-01

### 1. Features
- **docker:** add pandas-ta to docker images
- **metric:** alpha version for support multiple cluster on single ftmetric


<a name="v4.4.4"></a>
## [v4.4.4] - 2021-07-31

### 1. Features
- **metric:** add `freqtrade_perf_realized_pct` for calculated profit as percentage


<a name="v4.4.3"></a>
## [v4.4.3] - 2021-07-31

### 2. Performance Improvements
- **metric:** add `fthelper_internal_cache_size` label 'type' for local and global cache
- **metric:** add `freqtrade_perf_unrealized` and `freqtrade_perf_realized`
- **metric:** add new warmup process to refresh currency rate
- **metric:** reduce execute time in freqtrade_stat_*
- **metric:** reduce win_duration/draw_duration/loss_duration to single duration with type label

### 3. Bug Fixes
- **metric:** try to fix sometime fiat balance return zero value
- **metric:** realized and unrealized is swopping
- **metric:** `freqtrade_trade_avg_duration_seconds` not result correct information if average is more than 24 hours


<a name="v4.4.2"></a>
## [v4.4.2] - 2021-07-31

### 3. Bug Fixes
- **shared:** support version without prefix `v`


<a name="v4.4.1"></a>
## [v4.4.1] - 2021-07-31

### 1. Features
- **metric:** add new `freqtrade_stat_sell_reason` for win/draw/loss number of each sell reason
- **metric:** fully support warmup success rate for all processor
- **shared:** add version converter to convert version string to number

### 2. Performance Improvements
- **metric:** fthelper_build_info value will now change to new version when new version deployed
- **shared:** add testcase struct for build loopable test

### 3. Bug Fixes
- **metric:** crash after first called
- **metric:** fthelper_internal_warmup should be gauge.


<a name="v4.4.0"></a>
## [v4.4.0] - 2021-07-31

### 1. Features
- **metric:** add new `freqtrade_perf_daily` for daily profit percentage

### 2. Performance Improvements
- **metric:** cmd will not use same caches with global
- **metric:** add warmup for daily performance cal and log
- **metric:** change all `freqtrade_pair_performance` to `freqtrade_pair_perf`
- **metric:** remove warmup_error, and use warmup for warmup success rate instead.
- **metric:** warmup process will collect error and report
- **metric:** remove perf completely in freqtrade_pair
- **metric:** change global cache instead for warmup
- **shared:** exclude some data in cache in String() result
- **shared:** support AndD for add data and error to errors handler
- **shared:** create cache.bucket() for collect series of data
- **shared:** add Total() in error handler for calculate error percentage
- **shared:** add more error if cache key is empty
- **shared:** change some method signature and expose fetching status of cache data
- **shared:** add debug log how cache works

### 3. Bug Fixes
- **metric:** any global cache should use global cache
- **metric:** if daily balance is error it will keep trying
- **metric:** crash after first warmup occurred
- **metric:** crash when cannot get freqtrade data
- **metric:** warmup success rate should not report only when error occurred
- **shared:** total error in handler is not return correct value
- **shared:** check error in unit-test


<a name="v4.3.0"></a>
## [v4.3.0] - 2021-07-23

### 1. Features
- **gen:** add example for generate docker-compose files
- **shared:** loading .env support both files and directory

### 2. Performance Improvements
- **docs:** update example in generator readme
- **shared:** caches.SData will updated data before return value, and IsExist return false if error is exist
- **shared:** add support description log when run with `-test.v` option
- **shared:** check for nil value in xtests
- **shared:** datatype.ToInt will support all int bits
- **shared:** datatype convert will support all kind of int
- **shared:** add equal float number in xtests
- **shared:** add warning when using --list-config option
- **shared:** ForceArray will always return array even input is not
- **shared:** add new support for actual and bool
- **shared:** add testing helpers
- **shared:** remove version.go

### 3. Bug Fixes
- **gen:** docker type has been remove long time ago, use template instead
- **metric:** not checking error at some place
- **shared:** list-config deprecated should not always log
- **shared:** lint complaining
- **shared:** ForceString didn't return correct string when value is not string
- **shared:** linter warning
- **shared:** map shouldn't return null value if Gets still has value left


<a name="v4.2.0"></a>
## [v4.2.0] - 2021-07-21

### 2. Performance Improvements
- **gen:** add new info log when generators is executing


<a name="v4.2.0-beta.1"></a>
## [v4.2.0-beta.1] - 2021-07-21

### 2. Performance Improvements
- **gen:** add example code for copy type
- **gen:** refactor code and clustering runners

### 3. Bug Fixes
- **metric:** stop_buy shouldn't be default value if data is not exist


<a name="v4.1.1"></a>
## [v4.1.1] - 2021-07-20

### 2. Performance Improvements
- **docker:** cleanup do nothing, remove it
- **docker:** change base image for pg to develop_plot
- **shared:** add new function on template (toLower, toUpper, and toTitle)


<a name="v4.1.0"></a>
## [v4.1.0] - 2021-07-20

### 1. Features
- add 3 metrics at `freqtrade_stat` namespace.
- **docker:** add new buildflow for docker image with postgres support

### 2. Performance Improvements
- support build freqtrade image with postgres
- **docker:** add suffix in docker image tag name _pg
- **docker:** do some cleanup
- **docker:** instead libpg-dev for postgres usage
- **docker:** use psycopg2 instead of binary package


<a name="v4.1.0-beta.9"></a>
## [v4.1.0-beta.9] - 2021-07-14


<a name="v4.1.0-beta.8"></a>
## [v4.1.0-beta.8] - 2021-07-14

### 2. Performance Improvements
- **gen:** remove default settings


<a name="v4.1.0-beta.7"></a>
## [v4.1.0-beta.7] - 2021-07-14

### 3. Bug Fixes
- **gen:** use wrong config mapper


<a name="v4.1.0-beta.6"></a>
## [v4.1.0-beta.6] - 2021-07-14

### 1. Features
- **gen:** remove 'docker' type because it do nothing different from 'template' type
- **gen:** support generator multiple files base on input clusters
- **gen:** add config withCluster, meaning it will overide value that define on underscore fields by cluster
- **gen:** config also support multiple files/directories
- **gen:** remove all configs and template to use privately
- **metric:** add new metric for monitor warmup fails

### 2. Performance Improvements
- **gen:** add suffix in generator 'config' type
- **gen:** support override ftconfig from custom environment with specify key
- **gen:** disable withCluster in copy by default
- **gen:** remove support freqtrades in docker due to support clusters actively
- **gen:** change byEnv to byCluster in go template
- **gen:** now you can access data from generator parameters within .data map
- **gen:** add new type 'docker' for generator freqtrade docker-compsoe
- **metric:** use instead fetch expired caches instead
- **shared:** byEnv will fallback to get default value of cluster override is not exist
- **shared:** migrate envname to cluster to be the same with ftmetric
- **shared:** add new internal.meta fields in configuration
- **shared:** join function in template is support interface{} data type
- **shared:** add byEnv for get information base on envname (cluster)
- **shared:** remove default override value by cluster name

### 3. Bug Fixes
- **gen:** add default empty cluster name
- **gen:** enable 'withCluster' didn't end as expected
- **gen:** some place didn't been migrated to cluster
- **gen:** config not create correct file with suffix
- **metric:** sometimes metric return unhandler exception due to reading same data from different thread


<a name="v4.1.0-beta.5"></a>
## [v4.1.0-beta.5] - 2021-07-11

### 1. Features
- **gen:** add configuration and setting

### 2. Performance Improvements
- **config:** support override by envname from --env option
- **gen:** fully implementation of 'config' type
- **gen:** support new generator type 'config' for create config file
- **gen:** add default value of fs
- **shared:** support parsing number (int and float) in maps
- **shared:** add more function on go template

### 3. Bug Fixes
- **ci:** changelog generator fail because cannot push changes
- **ci:** look like gh is already installed
- **gen:** wrong config name
- **gen:** support `--env` in config to override data
- **gen:** remove default value with resolve strategy path, you can set '' (empty string) to disable path
- **metric:** cannot override port number due to different datatype
- **shared:** envname is not load when config with envvar
- **shared:** remove log that expose password set in env.


<a name="v4.1.0-beta.4"></a>
## [v4.1.0-beta.4] - 2021-07-11

### 3. Bug Fixes
- **deps:** invalid platform name for armv6 and armv7
- **gen:** ref in schema didn't works


<a name="v4.1.0-beta.3"></a>
## [v4.1.0-beta.3] - 2021-07-11

### 1. Features
- **gen:** support new type 'strategy' for create strategy code to output

### 2. Performance Improvements
- **deps:** add armv6 and armv7 for backward close [#1](https://github.com/kamontat/fthelper/issues/1)
- **metric:** add freqtrade url in info log
- **shared:** improve fs code and reduce wrapper

### 3. Bug Fixes
- **gen:** fs.variable is wrong name


<a name="v4.1.0-beta.2"></a>
## [v4.1.0-beta.2] - 2021-07-11

### 1. Features
- **core:** support template generator
- **shared:** support convert data to array via `a,b,c` format

### 2. Performance Improvements
- **core:** add new cli hooks for after_command
- **core:** change json generator to use inputs instead of templates
- **shared:** move config builder to before command

### 3. Bug Fixes
- **core:** config needs dotenv to get environment information


<a name="v4.1.0-beta.1"></a>
## [v4.1.0-beta.1] - 2021-07-11

### 1. Features
- refactor code and make generators works with json type
- change --configs to --config-dirs, --envs to --env-files and --no-env to --no-env-file [BREAK]

### 2. Performance Improvements
- **config:** remove default value in .env.docker

### 3. Bug Fixes
- **core:** migrated last change
- **shared:** fs.variables is not including with passing template


<a name="v4.0.0"></a>
## [v4.0.0] - 2021-07-10

### 2. Performance Improvements
- **core:** config command will show type of value


<a name="v4.0.0-beta.4"></a>
## [v4.0.0-beta.4] - 2021-07-10


<a name="v4.0.0-beta.3"></a>
## [v4.0.0-beta.3] - 2021-07-09


<a name="v4.0.0-beta.2"></a>
## [v4.0.0-beta.2] - 2021-07-09


<a name="v4.0.0-beta.1"></a>
## [v4.0.0-beta.1] - 2021-07-09

### 1. Features
- migrate ftmetric from private repository


<a name="v0.1.0-beta.12"></a>
## [v0.1.0-beta.12] - 2021-07-07


<a name="v0.1.0-beta.11"></a>
## [v0.1.0-beta.11] - 2021-07-07


<a name="v0.1.0-beta.10"></a>
## [v0.1.0-beta.10] - 2021-07-07


<a name="v0.1.0-beta.9"></a>
## [v0.1.0-beta.9] - 2021-07-07


<a name="v0.1.0-beta.8"></a>
## [v0.1.0-beta.8] - 2021-07-07


<a name="v0.1.0-beta.7"></a>
## [v0.1.0-beta.7] - 2021-07-07


<a name="v0.1.0-beta.6"></a>
## [v0.1.0-beta.6] - 2021-07-07


<a name="v0.1.0-beta.5"></a>
## [v0.1.0-beta.5] - 2021-07-07


<a name="v0.1.0-beta.4"></a>
## [v0.1.0-beta.4] - 2021-07-07


<a name="v0.1.0-beta.3"></a>
## [v0.1.0-beta.3] - 2021-07-07


<a name="v0.1.0-beta.2"></a>
## [v0.1.0-beta.2] - 2021-07-07


<a name="v0.1.0-beta.1"></a>
## v0.1.0-beta.1 - 2021-07-07

### 1. Features
- **init:** start new project


[Unreleased]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.14...HEAD
[v5.0.0-beta.14]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.13...v5.0.0-beta.14
[v5.0.0-beta.13]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.12...v5.0.0-beta.13
[v5.0.0-beta.12]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.11...v5.0.0-beta.12
[v5.0.0-beta.11]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.10...v5.0.0-beta.11
[v5.0.0-beta.10]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.9...v5.0.0-beta.10
[v5.0.0-beta.9]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.8...v5.0.0-beta.9
[v5.0.0-beta.8]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.7...v5.0.0-beta.8
[v5.0.0-beta.7]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.6...v5.0.0-beta.7
[v5.0.0-beta.6]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.5...v5.0.0-beta.6
[v5.0.0-beta.5]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.4...v5.0.0-beta.5
[v5.0.0-beta.4]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.3...v5.0.0-beta.4
[v5.0.0-beta.3]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.2...v5.0.0-beta.3
[v5.0.0-beta.2]: https://github.com/kamontat/fthelper/compare/v5.0.0-beta.1...v5.0.0-beta.2
[v5.0.0-beta.1]: https://github.com/kamontat/fthelper/compare/v4.5.3...v5.0.0-beta.1
[v4.5.3]: https://github.com/kamontat/fthelper/compare/v4.5.2...v4.5.3
[v4.5.2]: https://github.com/kamontat/fthelper/compare/v4.5.1...v4.5.2
[v4.5.1]: https://github.com/kamontat/fthelper/compare/v4.5.0...v4.5.1
[v4.5.0]: https://github.com/kamontat/fthelper/compare/v4.5.0-beta.1...v4.5.0
[v4.5.0-beta.1]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.8...v4.5.0-beta.1
[v4.5.0-alpha.8]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.7...v4.5.0-alpha.8
[v4.5.0-alpha.7]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.6...v4.5.0-alpha.7
[v4.5.0-alpha.6]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.5...v4.5.0-alpha.6
[v4.5.0-alpha.5]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.4...v4.5.0-alpha.5
[v4.5.0-alpha.4]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.3...v4.5.0-alpha.4
[v4.5.0-alpha.3]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.2...v4.5.0-alpha.3
[v4.5.0-alpha.2]: https://github.com/kamontat/fthelper/compare/v4.5.0-alpha.1...v4.5.0-alpha.2
[v4.5.0-alpha.1]: https://github.com/kamontat/fthelper/compare/v4.4.4...v4.5.0-alpha.1
[v4.4.4]: https://github.com/kamontat/fthelper/compare/v4.4.3...v4.4.4
[v4.4.3]: https://github.com/kamontat/fthelper/compare/v4.4.2...v4.4.3
[v4.4.2]: https://github.com/kamontat/fthelper/compare/v4.4.1...v4.4.2
[v4.4.1]: https://github.com/kamontat/fthelper/compare/v4.4.0...v4.4.1
[v4.4.0]: https://github.com/kamontat/fthelper/compare/v4.3.0...v4.4.0
[v4.3.0]: https://github.com/kamontat/fthelper/compare/v4.2.0...v4.3.0
[v4.2.0]: https://github.com/kamontat/fthelper/compare/v4.2.0-beta.1...v4.2.0
[v4.2.0-beta.1]: https://github.com/kamontat/fthelper/compare/v4.1.1...v4.2.0-beta.1
[v4.1.1]: https://github.com/kamontat/fthelper/compare/v4.1.0...v4.1.1
[v4.1.0]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.9...v4.1.0
[v4.1.0-beta.9]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.8...v4.1.0-beta.9
[v4.1.0-beta.8]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.7...v4.1.0-beta.8
[v4.1.0-beta.7]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.6...v4.1.0-beta.7
[v4.1.0-beta.6]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.5...v4.1.0-beta.6
[v4.1.0-beta.5]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.4...v4.1.0-beta.5
[v4.1.0-beta.4]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.3...v4.1.0-beta.4
[v4.1.0-beta.3]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.2...v4.1.0-beta.3
[v4.1.0-beta.2]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.1...v4.1.0-beta.2
[v4.1.0-beta.1]: https://github.com/kamontat/fthelper/compare/v4.0.0...v4.1.0-beta.1
[v4.0.0]: https://github.com/kamontat/fthelper/compare/v4.0.0-beta.4...v4.0.0
[v4.0.0-beta.4]: https://github.com/kamontat/fthelper/compare/v4.0.0-beta.3...v4.0.0-beta.4
[v4.0.0-beta.3]: https://github.com/kamontat/fthelper/compare/v4.0.0-beta.2...v4.0.0-beta.3
[v4.0.0-beta.2]: https://github.com/kamontat/fthelper/compare/v4.0.0-beta.1...v4.0.0-beta.2
[v4.0.0-beta.1]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.12...v4.0.0-beta.1
[v0.1.0-beta.12]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.11...v0.1.0-beta.12
[v0.1.0-beta.11]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.10...v0.1.0-beta.11
[v0.1.0-beta.10]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.9...v0.1.0-beta.10
[v0.1.0-beta.9]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.8...v0.1.0-beta.9
[v0.1.0-beta.8]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.7...v0.1.0-beta.8
[v0.1.0-beta.7]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.6...v0.1.0-beta.7
[v0.1.0-beta.6]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.5...v0.1.0-beta.6
[v0.1.0-beta.5]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.4...v0.1.0-beta.5
[v0.1.0-beta.4]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.3...v0.1.0-beta.4
[v0.1.0-beta.3]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.2...v0.1.0-beta.3
[v0.1.0-beta.2]: https://github.com/kamontat/fthelper/compare/v0.1.0-beta.1...v0.1.0-beta.2
