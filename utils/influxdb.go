package utils

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/rsms/go-log"
)

var org = os.Getenv("INFLUX_ORG")
var bucket = os.Getenv("INFLUX_BUCKET")

var influx = getInfluxClient()

func getInfluxClient() influxdb.Client {
	client := influxdb.NewClient(os.Getenv("INFLUX_HOST"), os.Getenv("INFLUX_TOKEN"))
	online, err := client.Ready(context.TODO())
	if !online || err != nil {
		log.RootLogger.GoLogger().Fatal("Influxdb not running.")
	}

	return client
}

var influxWriteApi = influx.WriteAPI(org, bucket)

func InsertMeasurement(name string, tags map[string]interface{}, values map[string]interface{}) {
	tagsString := ""
	valuesString := ""
	for k, v := range tags {
		tagsString += fmt.Sprintf(",%v=%v", k, v)
	}
	for k, v := range values {
		valuesString += fmt.Sprintf("%v=%v,", k, v)
	}
	if strings.HasSuffix(valuesString, ",") {
		valuesString = valuesString[:len(valuesString)-1]
	}
	influxWriteApi.WriteRecord(fmt.Sprintf("%v%v %v %d", name, tagsString, valuesString, time.Now().UnixNano()))
	influxWriteApi.Flush()
}
