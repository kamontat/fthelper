<a name="unreleased"></a>
## [Unreleased]

### Chore
- push to unknown branch is invalid
- push change before create PR
- forget to add github token
- remove gh auth completely
- remove login

### Fix
- **ci:** look like gh is already installed
- **ci:** changelog generator fail because cannot push changes
- **shared:** remove log that expose password set in env.


<a name="v4.1.0-beta.4"></a>
## [v4.1.0-beta.4] - 2021-07-11
### Chore
- **release:** published 'v4.1.0-beta.4'

### Fix
- **deps:** invalid platform name for armv6 and armv7
- **gen:** ref in schema didn't works


<a name="v4.1.0-beta.3"></a>
## [v4.1.0-beta.3] - 2021-07-11
### Chore
- update json schema
- add merge bot
- separate unittest of each module
- add test data to strategy
- use same format for all workflow
- remove deprecated code
- auto generate changelog every week on monday
- add codeql to analysis
- update readme
- add dependabot
- **ci:** support manually trigger
- **ci:** format code
- **deps:** add chglog to changelog generator
- **doc:** add goreleaser limitation
- **gen:** named file correctly it
- **generator:** remove demo templates
- **release:** published 'v4.1.0-beta.3'
- **shared:** move utils.ToArray to datatype.ToArray

### Feat
- **gen:** support new type 'strategy' for create strategy code to output

### Fix
- **gen:** fs.variable is wrong name

### Perf
- **deps:** add armv6 and armv7 for backward close [#1](https://github.com/kamontat/fthelper/issues/1)
- **metric:** add freqtrade url in info log
- **shared:** improve fs code and reduce wrapper


<a name="v4.1.0-beta.2"></a>
## [v4.1.0-beta.2] - 2021-07-11
### Chore
- add migration plan and gitgo config
- move dotenv up to be execute before config
- comment fmt.Printf
- add template testcase
- move debug log to understand how log works
- **release:** published 'v4.1.0-beta.2'

### Feat
- **core:** support template generator
- **shared:** support convert data to array via `a,b,c` format

### Fix
- **core:** config needs dotenv to get environment information

### Perf
- **core:** add new cli hooks for after_command
- **core:** change json generator to use inputs instead of templates
- **shared:** move config builder to before command


<a name="v4.1.0-beta.1"></a>
## [v4.1.0-beta.1] - 2021-07-11
### Chore
- remove unused scripts
- just changes
- **release:** published 'v4.1.0-beta.1'

### Feat
- refactor code and make generators works with json type
- change --configs to --config-dirs, --envs to --env-files and --no-env to --no-env-file [BREAK]

### Fix
- **core:** migrated last change
- **shared:** fs.variables is not including with passing template

### Perf
- **config:** remove default value in .env.docker


<a name="v4.0.0"></a>
## [v4.0.0] - 2021-07-10
### Chore
- **doc:** update ftmetric readme
- **release:** published 'v4.0.0'

### Perf
- **core:** config command will show type of value


<a name="v4.0.0-beta.4"></a>
## [v4.0.0-beta.4] - 2021-07-10
### Chore
- publish binary script in Github release
- add simple readme
- **doc:** add ftmetric readme
- **doc:** add docker image registry changes
- **release:** published 'v4.0.0-beta.4'


<a name="v4.0.0-beta.3"></a>
## [v4.0.0-beta.3] - 2021-07-09
### Chore
- go mod tidy
- **release:** published 'v4.0.0-beta.3'


<a name="v4.0.0-beta.2"></a>
## [v4.0.0-beta.2] - 2021-07-09
### Chore
- empty message
- **release:** published 'v4.0.0-beta.2'


<a name="v4.0.0-beta.1"></a>
## [v4.0.0-beta.1] - 2021-07-09
### Chore
- change generator to v4 as well
- **release:** published 'v4.0.0-beta.1'

### Feat
- migrate ftmetric from private repository


<a name="v0.1.0-beta.12"></a>
## [v0.1.0-beta.12] - 2021-07-07
### Chore
- needs login to github container registry
- **release:** published 'v0.1.0-beta.12'


<a name="v0.1.0-beta.11"></a>
## [v0.1.0-beta.11] - 2021-07-07
### Chore
- access packages write permission
- **release:** published 'v0.1.0-beta.11'


<a name="v0.1.0-beta.10"></a>
## [v0.1.0-beta.10] - 2021-07-07
### Chore
- use buildx to build docker
- **release:** published 'v0.1.0-beta.10'


<a name="v0.1.0-beta.9"></a>
## [v0.1.0-beta.9] - 2021-07-07
### Chore
- **release:** published 'v0.1.0-beta.9'


<a name="v0.1.0-beta.8"></a>
## [v0.1.0-beta.8] - 2021-07-07
### Chore
- **release:** published 'v0.1.0-beta.8'


<a name="v0.1.0-beta.7"></a>
## [v0.1.0-beta.7] - 2021-07-07
### Chore
- **ci:** add busybox in docker manifest
- **release:** published 'v0.1.0-beta.7'


<a name="v0.1.0-beta.6"></a>
## [v0.1.0-beta.6] - 2021-07-07
### Chore
- enable debug mode in ci
- **release:** published 'v0.1.0-beta.6'


<a name="v0.1.0-beta.5"></a>
## [v0.1.0-beta.5] - 2021-07-07
### Chore
- remove unused dist and add docker for busybox and stratch
- **release:** published 'v0.1.0-beta.5'


<a name="v0.1.0-beta.4"></a>
## [v0.1.0-beta.4] - 2021-07-07
### Chore
- update gitignore config and goreleaser config
- **release:** published 'v0.1.0-beta.4'


<a name="v0.1.0-beta.3"></a>
## [v0.1.0-beta.3] - 2021-07-07
### Chore
- wrong directory name
- **release:** published 'v0.1.0-beta.3'


<a name="v0.1.0-beta.2"></a>
## [v0.1.0-beta.2] - 2021-07-07
### Chore
- remove docker hub login
- **release:** published 'v0.1.0-beta.2'


<a name="v0.1.0-beta.1"></a>
## v0.1.0-beta.1 - 2021-07-07
### Chore
- **release:** published 'v0.1.0-beta.1'

### Feat
- **init:** start new project


[Unreleased]: https://github.com/kamontat/fthelper/compare/v4.1.0-beta.4...HEAD
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
