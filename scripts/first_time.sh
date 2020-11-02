#!/usr/bin/env bash

source $(dirname $0)/helpers.sh

cd $(dirname $0)/..

log "INF" "Starting influxdb."
docker-compose up -d influxdb

log "INF" "Waiting for influxdb to be ready."
sleep 3

bash $(dirname $0)/setup_influxdb.sh

log "INF" "Start all other services."
docker-compose up -d
log "SUC" "Done."