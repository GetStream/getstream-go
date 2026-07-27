#!/usr/bin/env bash

set -euo pipefail

title=""
body=""
body_file=""
output=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --title)
      title="${2:-}"
      shift 2
      ;;
    --body)
      body="${2:-}"
      shift 2
      ;;
    --body-file)
      body_file="${2:-}"
      shift 2
      ;;
    --output)
      output="${2:-}"
      shift 2
      ;;
    *)
      echo "unknown argument: $1" >&2
      exit 1
      ;;
  esac
done

if [[ -n "${body_file}" ]]; then
  body="$(<"${body_file}")"
fi

determine_bump() {
  local pr_title="$1"
  local pr_body="$2"

  if [[ "${pr_body}" =~ BREAKING[[:space:]-]CHANGE ]]; then
    echo "major"
    return
  fi

  if ! echo "${pr_title}" | grep -Eq '^([a-zA-Z]+)(\([^)]+\))?(!)?:'; then
    echo "none"
    return
  fi

  if echo "${pr_title}" | grep -Eq '^([a-zA-Z]+)(\([^)]+\))?!:'; then
    echo "major"
    return
  fi

  local type
  type="$(echo "${pr_title}" | sed -E 's/^([a-zA-Z]+).*/\1/' | tr '[:upper:]' '[:lower:]')"
  if [[ "${type}" == "feat" ]]; then
    echo "minor"
    return
  fi
  if [[ "${type}" == "fix" || "${type}" == "bug" ]]; then
    echo "patch"
    return
  fi

  echo "none"
}

bump="$(determine_bump "${title}" "${body}")"

write_output() {
  local key="$1"
  local value="$2"
  if [[ -n "${output}" ]]; then
    echo "${key}=${value}" >> "${output}"
  else
    echo "${key}=${value}"
  fi
}

if [[ "${bump}" == "none" ]]; then
  write_output "should_release" "false"
  write_output "bump" "none"
  exit 0
fi

latest_version="$(git tag --list | sed 's/^v//' | grep -E '^[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n 1)"
if [[ -z "${latest_version}" ]]; then
  latest_version="0.0.0"
fi

IFS='.' read -r major minor patch <<< "${latest_version}"
case "${bump}" in
  major)
    next_version="$((major + 1)).0.0"
    ;;
  minor)
    next_version="${major}.$((minor + 1)).0"
    ;;
  patch)
    next_version="${major}.${minor}.$((patch + 1))"
    ;;
esac

sed -i "s|versionName = \".*\"|versionName = \"v${next_version}\"|" version.go

write_output "should_release" "true"
write_output "bump" "${bump}"
write_output "previous_version" "${latest_version}"
write_output "version" "${next_version}"
write_output "tag" "v${next_version}"
