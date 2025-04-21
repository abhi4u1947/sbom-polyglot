#!/usr/bin/env bash

set -eu -o pipefail

function log() {
    printf "%s\n" "$@" >&2
}

test -z "${RUNNER_DEBUG+x}" || set -x

CONF_PROJECT_DIR="${CONF_PROJECT_DIR:-$PWD}"
CONF_ECOSYSTEMS="${CONF_ECOSYSTEMS:-}"
CONF_RESULT_PATH="${CONF_RESULT_PATH:-./bom.json}"

cmd_args=( "--recurse" "--output=$CONF_RESULT_PATH" "--profile=license-compliance" "$CONF_PROJECT_DIR" )

IFS=' ' read -r -a ecosystems <<< "$CONF_ECOSYSTEMS"
for ecosystem in "${ecosystems[@]}"; do
    cmd_args+=( "--type=$ecosystem" )
done

log "Generating SBOM using cdxgen..."
FETCH_LICENSE=true cdxgen "${cmd_args[@]}"
