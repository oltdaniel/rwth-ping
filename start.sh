#!/usr/bin/env bash

source scripts/helpers.sh

if [[ ! -f influxdb.env ]]; then
    log "ERR" "Ensure to setup your influxdb."
    exit 1
fi

COMMAND="INFLUX_HOST=http://localhost:8086 D=1 go run main.go"
COMMAND="$(cat influxdb.env | tr -d '\r' | while read r; do printf "$r "; done) $COMMAND"

log "INF" "$COMMAND"

eval $COMMAND