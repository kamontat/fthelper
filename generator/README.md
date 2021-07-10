```bash
ftgenerator \
  --config-dirs "/etc/configs" \
  --env-files "/hello/.env.default" \
  --env-files "/hello/.env.production" \
  --envs "live" --envs "dryrun" --envs "local"
```
