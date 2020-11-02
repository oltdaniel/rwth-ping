#!/usr/bin/env bash

source $(dirname $0)/helpers.sh

log "INF" "Creating user, bucket and create access token."
/bin/bash -c 'docker run -it --network host quay.io/influxdb/influxdb:2.0.0-rc /bin/sh -c "influx setup -f -b demo_bucket -o demo_org -u demo_user -p baconbacon && influx bucket create -o demo_org -n rwth_ping && influx auth create -o demo_org --read-buckets --write-buckets"' \
    | tail -n +4 | awk -f `dirname $0`/influxdb.awk > `dirname $0`/../influxdb.env

log "INF" "Generated config:"
cat `dirname $0`/../influxdb.env
log "SUC" "Config written."