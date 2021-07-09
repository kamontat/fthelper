```bash
ftgenerator \
  --config-dir "/etc/configs" \
  --env-file "/hello/.env.default" \
  --env-file "/hello/.env.production" \
  --envs "live" --envs "dryrun" --envs "local"
```
