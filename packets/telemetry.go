package packets

import (
	"bytes"
	"encoding/binary"
	"log"
)

// This packet details telemetry for all the cars in the race.
// It details various values that would be recorded on the car such as speed,
// throttle application, DRS etc.

// Frequency: Rate as specified in menus
// Size: 1085 bytes

type CarTelemetryData struct {
	MSpeed                   uint16     // Speed of car in kilometres per hour
	MThrottle                uint8      // Amount of throttle applied (0 to 100)
	MSteer                   int8       // Steering (-100 (full lock left) to 100 (full lock right))
	MBrake                   uint8      // Amount of brake applied (0 to 100)
	MClutch                  uint8      // Amount of clutch applied (0 to 100)
	MGear                    int8       // Gear selected (1-8, N=0, R=-1)
	MEngineRPM               uint16     // Engine RPM
	MDrs                     uint8      // 0 = off, 1 = on
	MRevLightsPercent        uint8      // Rev lights indicator (percentage)
	MBrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	MTyresSurfaceTemperature [4]uint16  // Tyres surface temperature (celsius)
	MTyresInnerTemperature   [4]uint16  // Tyres inner temperature (celsius)
	MEngineTemperature       uint16     // Engine temperature (celsius)
	MTyresPressure           [4]float32 // Tyres pressure (PSI)
}

type PacketCarTelemetryData struct {
	MHeader           PacketHeader
	MCarTelemetryData [20]CarTelemetryData
	MButtonStatus     uint32 // Bit flags specifying which buttons are being pressed currently - see appendices
}

func NewPacketCarTelemetryData(buf []byte) PacketCarTelemetryData {
	var pack PacketCarTelemetryData
	r := bytes.NewReader(buf)
	if err := binary.Read(r, binary.LittleEndian, &pack); err != nil {
		log.Fatal(err)
	}
	return pack
}
