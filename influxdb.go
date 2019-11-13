package main

import "github.com/influxdata/influxdb-client-go"

func CreateDb() {
	_, err := influxdb.New("http://influxdb:8086", "")
	if err != nil {
		panic(err)
	}
}
