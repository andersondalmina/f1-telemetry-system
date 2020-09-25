package packets

import (
	"bytes"
	"encoding/binary"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2Write "github.com/influxdata/influxdb-client-go/v2/api/write"
)

// Packet ID 6
// This packet details telemetry for all the cars in the race.
// It details various values that would be recorded on the car such as speed,
// throttle application, DRS etc.

// Frequency: Rate as specified in menus
// Size: 1085 bytes

type CarTelemetryData struct {
	Speed                   uint16     // Speed of car in kilometres per hour
	Throttle                uint8      // Amount of throttle applied (0 to 100)
	Steer                   int8       // Steering (-100 (full lock left) to 100 (full lock right))
	Brake                   uint8      // Amount of brake applied (0 to 100)
	Clutch                  uint8      // Amount of clutch applied (0 to 100)
	Gear                    int8       // Gear selected (1-8, N=0, R=-1)
	RPM                     uint16     // Engine RPM
	DRS                     uint8      // 0 = off, 1 = on
	RevLightsPercent        uint8      // Rev lights indicator (percentage)
	BrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint16  // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint16  // Tyres inner temperature (celsius)
	EngineTemperature       uint16     // Engine temperature (celsius)
	TyresPressure           [4]float32 // Tyres pressure (PSI)
}

type TelemetryPacket struct {
	Header           PacketHeader
	CarTelemetryData [20]CarTelemetryData
	ButtonStatus     uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

func NewTelemetryPacket(buffer []byte) TelemetryPacket {
	var pack TelemetryPacket
	r := bytes.NewReader(buffer)
	if err := binary.Read(r, binary.LittleEndian, &pack); err != nil {
		log.Fatal(err)
	}

	return pack
}

func (p TelemetryPacket) GetPlayerData() CarTelemetryData {
	return p.CarTelemetryData[p.Header.PlayerIndex]
}

func (p TelemetryPacket) CreatePoint() *influxdb2Write.Point {
	d := p.GetPlayerData()

	return influxdb2.NewPointWithMeasurement("telemetry").
		AddField("speed", int(d.Speed)).
		AddField("gear", int(d.Gear)).
		AddField("throttle", int(d.Throttle)).
		AddField("brake", int(d.Brake)).
		AddField("rpm", int(d.RPM))
}
