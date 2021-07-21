# Copy

This example will show how to copy file from `A` directory to `B` directory using docker image.

```bash
docker run --rm -it \
  -v "$PWD/configs:/fth/configs" \
  -v "$PWD/A:/fth/A" -v "$PWD/B:/fth/B" \
  "ghcr.io/kamontat/ftgenerator"
```

## Description

If you open [configs/generators.json](./configs/generators.json), you will find this config

```json
{
  "type": "copy",
  "input": "adir",
  "output": "bdir"
}
```

and on [configs/fs.json](./configs/fs.json)

```json
"adir": {
  "type": "directory",
  "mode": "single",
  "fullpath": "{{ .current }}/A"
},
"bdir": {
  "type": "directory",
  "mode": "single",
  "fullpath": "{{ .current }}/B"
}
```

Basically, `copy` type will copy from `input` to `output` where input and output is key name of json in **fs.json** file. So above config is meaning copy from `adir` to `bdir` where adir is single directory at `$PWD/A` path and bdir is single directory at `$PWD/B`.
