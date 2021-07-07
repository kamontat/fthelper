#!/usr/bin/env bash
# shellcheck disable=SC1000

# generate by 2.3.2
# link (https://github.com/Template-generator/script-genrating/tree/2.3.2)

set -e
# set -x #DEBUG - Display commands and their arguments as they are executed.
# set -v #VERBOSE - Display shell input lines as they are read.
# set -n #EVALUATE - Check syntax of the script but don't execute.

appname="${APPNAME:-1}"

GITHUB_ENV="${GITHUB_ENV:-.env}"
echo "Current Github Environment file: $GITHUB_ENV"

# show all tag currently
git tag --column

latest_version="$(git describe --abbrev=0 --tags --match "$appname/*")"
previous_version="$(git describe --abbrev=0 --tags --match "$appname/*" "$latest_version^")"

export VERSION="${latest_version//$appname\//}"
export PREV_VERSION="${previous_version//$appname\//}"

echo "Version: $VERSION"
echo "Previous: $PREV_VERSION"
echo "
APP_VERSION=$VERSION
" >>"$GITHUB_ENV"

# GORELEASER_CURRENT_TAG=$VERSION
# GORELEASER_PREVIOUS_TAG=$PREV_VERSION
