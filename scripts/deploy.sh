#!/usr/bin/env bash
# shellcheck disable=SC1000

# generate by 2.3.2
# link (https://github.com/Template-generator/script-genrating/tree/2.3.2)

set -e
# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

# 1. Select application(s) to deploy
#    1. Validate repository
#    2. Enter release version number
#    3. Commit changes and Create tag
#    4. Push change to Github

error() {
  echo "$@" >&2
  exit 1
}

echo "1. list all tags that already exist:"
git tag -l --column "v*"

version=""
valid_tag=false
while ! $valid_tag; do
  printf "2. enter version (v0.0.0): "
  read -r version
  if [[ $version == "" ]]; then
    echo "exit"
    exit 0
  fi

  if [[ "$version" =~ ^v ]]; then # must has prefix v
    if ! git tag | grep -q "^$version$"; then
      valid_tag=true
    fi
  fi
done

echo "3. commit with release message"
git add .
git commit --allow-empty -m "chore(release): published '$version'"

echo "4. create git tag called '$version'"
git tag "$version"

echo "5. push all changes and tag to Github repository"
git push && git push --tag
