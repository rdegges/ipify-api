#!/bin/sh
set -e

if [ "$1" != "./ipify-api" ]; then
  exec "$@"
  exit $?
fi

exec "$@" -logtostderr=true -stderrthreshold=${LOGGING_LEVEL}
