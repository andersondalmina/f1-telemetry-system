package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2Write "github.com/influxdata/influxdb-client-go/v2/api/write"
)

// Packet ID 2
// The lap data packet gives details of all the cars in the session.
// Frequency: Rate as specified in menus
// Size: 841 bytes

type LapData struct {
	LastLapTime   float32 // Last lap time in seconds
	CarPosition   uint8   // Car race position
	CurrentLapNum int8    // Current lap number
}

type LapPacket struct {
	Header  PacketHeader
	LapData [20]LapData
}

func NewLapPacket(buffer []byte) LapPacket {
	var pack LapPacket
	r := bytes.NewReader(buffer)
	if err := binary.Read(r, binary.LittleEndian, &pack); err != nil {
		log.Fatal(err)
	}

	return pack
}

func (p LapPacket) getMyCarData() LapData {
	return p.LapData[p.Header.PlayerIndex]
}

func (p LapPacket) ExportPoint() *influxdb2Write.Point {
	return influxdb2.NewPointWithMeasurement("lap").
		AddField("current", p.getMyCarData().CurrentLapNum)
}
