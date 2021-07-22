# Multiple Cluster generator

Only require data is json configs file that contains `generators` field.

```json
{
  "$schema": "https://raw.githubusercontent.com/kamontat/fthelper/main/generator/schema/generator.json",
  "generators": []
}
```

## Examples

1. [Copy](./example/copy) - Copy file/directory to output location
2. [Docker](./example/docker) - Generate docker-compose from configuration
