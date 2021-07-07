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

directories=()

echo "1. ftgenerator"
echo "2. ftmetric"
echo "3. all above"
printf "Choose Application to release (1|2|3): "
read -rn 1 app
echo

if [[ "$app" == 1 ]]; then
  directories+=("generator")
elif [[ "$app" == 2 ]]; then
  directories+=("metric")
elif [[ "$app" == 3 ]]; then
  directories+=("generator" "metric")
else
  error "your input is not match is supported number, try again"
fi

root="$PWD"
for directory in "${directories[@]}"; do
  name="ft$directory"
  echo "Publishing $name:"
  cd "$root/$directory" || exit 1

  echo "  1. list all tags that already exist:"
  git tag -l --column "$name/*"

  version=""
  valid_tag=false
  while ! $valid_tag; do
    printf "  2. select version for %s (v0.0.0): " "$name"
    read -r version
    if [[ $version == "" ]]; then
      echo "exit"
      exit 0
    fi

    if [[ "$version" =~ ^v ]]; then # must has prefix v
      if ! git tag | grep -q "^$name/$version$"; then
        valid_tag=true
      fi
    fi
  done
  version="$name/$version"

  echo "  3. commit with release message"
  git add .
  git commit --allow-empty -m "chore(release): published '$version'"

  echo "  4. create git tag called '$version'"
  git tag "$version"

  cd "$root" || exit 1
done

echo "push all changes and tag to Github repository"
git push && git push --tag
