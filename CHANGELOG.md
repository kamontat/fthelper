<a name="unreleased"></a>
## [Unreleased]

### 1. Features
- **gen:** add config withCluster, meaning it will overide value that define on underscore fields by cluster
- **gen:** config also support multiple files/directories
- **gen:** remove all configs and template to use privately

### 2. Performance Improvements
- **gen:** add suffix in generator 'config' type
- **gen:** add new type 'docker' for generator freqtrade docker-compsoe
- **gen:** now you can access data from generator parameters within .data map
- **shared:** byEnv will fallback to get default value of cluster override is not exist
- **shared:** migrate envname to cluster to be the same with ftmetric
- **shared:** add byEnv for get information base on envname (cluster)
- **shared:** remove default override value by cluster name

### 3. Bug Fixes
- **gen:** some place didn't been migrated to cluster
- **gen:** config not create correct file with suffix


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


[Unreleased]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.5...HEAD
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
