# Configuration

The idea of configuration is a setting value that easy change and support variant of data source.

## Value

Value from environment and argument are always `String`. We have 2 type of parser. If that value has default value on config file/directory, we will try to pass data base on default value type, else the value will be passed base on nature of value. For example, if you has numeric string `1` without default value it will pass to number 1 which might not be true. So we recommend every value need default value in configuration file/directory.

## Data source

We support 3 data source, ranked by

1. `directory/file` - This consider as base configuration
2. `environment` - This will override `directory/file`, usually for secret key or token
3. `argument` - This will override everything else

Each of datasource can specify cluster override value.

## Cluster

We have special key in json/argument is `_` or for environment is `FTC` for dynamic override configuration base on input clusters name.

## Directories / Files

All configuration files must be json format; otherwise, it will crash when loading. If you enter the directory, all files inside that directory must be json. Following key is forbidded for internal use only

1. `internal` - for internal configuration
2. `fs` - for file-system configuration
3. `_` - for overrided cluster configuration

> Be aware when you add config with uppercase, the environment will not works because environment result always return lower case.

### Example configuration

- Normal config

```json
{
  "value": true,
  "languages": ["EN", "TH", "AB"],
  "image": {
    "name": "hello",
    "version": "v1.0.0"
  }
}
```

- Override config with cluster 1A

```json
{
  "value": "test",
  "_": {
    "1A": {
      "value": "prod"
    }
  }
}
```

## Environment

Every configuration must prefix with either `FTH_` or `FTC_`. Data type that environment support are string, int, float, bool, array, and map. Each key must separated by `__`. All key will always transform to lower case.

### Example configuration

```
FTH_TEST=true -> {"test": true}
FTH_CLUSTER__NAME=1A -> {"cluster": {"name": "1A"}}
FTH_NAMES=array:50,60,70 -> {"cluster": ["50", "60", "70"]}
FTC_1A__CLUSTER__NAME=1A -> {"_": {"1a": {"cluster": {"name": "1A"}}}}
```

## Argument

Each of argument must formatted as `<key>=<value>` (e.g. **test.enabled=true**). The data type if determine by configuration at json file.

### Example configuration

```bash
# This mean test.enabled is false if cluster is 1A
command_cli \
  test.enabled=true \
  _.1A.test.enabled=false

command_cli \
  "hello=message with space" \
  number=10
```

## Limitation

1. You should not use underscroll in any place except for clustering purpose
     - We cannot determine whether that underscroll is actual underscroll or dash