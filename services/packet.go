package services

import (
	"github.com/andersondalmina/f1-telemetry/packets"
	"github.com/andersondalmina/f1-telemetry/persist"
)

func ProcessPacket(buffer []byte) {
	header := packets.NewPacketHeader(buffer)

	switch header.PacketID {
	case 2:
		pack := packets.NewLapPacket(buffer)
		persist.WritePoint(pack.CreatePoint())
	case 6:
		pack := packets.NewTelemetryPacket(buffer)
		persist.WritePoint(pack.CreatePoint())
	}
}
