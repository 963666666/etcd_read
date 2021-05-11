#!/usr/bin/env bash
set -e



function hello {
  local tool="$1"
  local nowExit="$0"
  echo "echo dengdeng" ${nowExit}
  echo "hello world" ${tool}
}

RESULT=$(hello world)

echo ${RESULT}


function run_for_module {
  local module=${1:-"."}
  echo ${module}
  shift 1
#  (
#    cd "${ETCD_ROOT_DIR}/${module}" && "$@"
#  )
}

RESULT=$(run_for_module ./tools/mod run gobin ${GOBINARGS:-} -p -m --mod=readonly "${tool}")

echo ${RESULT}