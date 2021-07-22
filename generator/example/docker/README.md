# Docker compose

This example will show how to use ftgenerator to generate docker-compose with cluster support.

```bash
docker run --rm -it \
  -v "$PWD/.env:/fth/.env" \
  -v "$PWD/configs:/fth/configs" \
  -v "$PWD/templates:/fth/templates" \
  -v "$PWD/output:/fth/output" \
  "ghcr.io/kamontat/ftgenerator" \
  --clusters 1A --clusters 2A
```

## Require steps

1. Create directory name `.env` that contains `.env.secret` file
2. Copy `.env.default` content to `.env.secret`
3. Add correct information in `.env` files
4. Double check config data at [configs/docker.json](./configs/docker.json)
