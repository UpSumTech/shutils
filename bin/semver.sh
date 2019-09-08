#!/usr/bin/env bash

MAJOR=
MINOR=
PATCH=

ok() {
  echo -n ''
}

parseSemver() {
  local semver="$1"
  local regex="v([0-9]+)\.([0-9]+)\.([0-9]+)"

  if [[ "$semver" =~ $regex ]]; then
    MAJOR=${BASH_REMATCH[1]}
    MINOR=${BASH_REMATCH[2]}
    PATCH=${BASH_REMATCH[3]}
  fi
  ok
}

bumpSemver() {
  local semver="$1"
  local version="$2"

  parseSemver "$semver"

  case "$version" in
    major)
      MAJOR=$((MAJOR + 1))
      MINOR=0
      PATCH=0
      ;;
    minor)
      MINOR=$((MINOR + 1))
      PATCH=0
      ;;
    patch)
      PATCH=$((PATCH + 1))
      ;;
    *)
      echo "ERROR >> Not a valid version type" >/dev/stderr
      exit 1
      ;;
  esac
  ok
}

main() {
  local version="$1"
  local semver="$2"
  [[ "x$semver" == "x" ]] && semver='v0.0.0'
  bumpSemver "$semver" "$version"
  echo "v${MAJOR}.${MINOR}.${PATCH}"
}

[[ "$BASH_SOURCE" == "$0" ]] && main "$@"
