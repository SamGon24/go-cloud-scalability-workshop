#!/usr/bin/env bash
# Fire 50 concurrent requests at a local server and report total time.
#
# Usage:
#   ./load-test.sh                          # defaults: 50 reqs to /process
#   ./load-test.sh 100                      # 100 reqs
#   ./load-test.sh 50 /upload               # different path
#   ./load-test.sh 50 /upload POST          # with method
#
# Requires: curl. No other dependencies.

set -euo pipefail

N="${1:-50}"
PATH_="${2:-/process}"
METHOD="${3:-GET}"
URL="http://localhost:8081${PATH_}"

echo "firing ${N} concurrent ${METHOD} requests at ${URL}"
start=$(date +%s)

pids=()
for i in $(seq 1 "${N}"); do
  curl -s -o /dev/null -X "${METHOD}" "${URL}" &
  pids+=($!)
done

for pid in "${pids[@]}"; do
  wait "${pid}"
done

end=$(date +%s)
echo "done in $((end - start))s"
