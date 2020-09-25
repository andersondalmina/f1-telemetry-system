package persist

import (
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2Write "github.com/influxdata/influxdb-client-go/v2/api/write"
)

var client influxdb2.Client

func CreateClient() {
	client = influxdb2.NewClient("http://0.0.0.0:8086", "")
}

func CloseClient() {
	client.Close()
}

func WritePoint(p *influxdb2Write.Point) {
	writeAPI := client.WriteAPIBlocking("f1", "f1")

	p.SetTime(time.Now())

	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Println(err)
	}
}
