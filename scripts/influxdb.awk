NR == 2 {
    printf "INFLUX_BUCKET=%s\n", $1
    printf "INFLUX_ORG=%s\n", $4
}

NR == 4 {
    printf "INFLUX_TOKEN=%s\n", $2
}