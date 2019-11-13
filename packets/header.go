package packets

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Packages
// 0 - Motion
// 1 - Session
// 2 - Lap Data
// 3 - Event
// 4 - Participants
// 5 - Car Setups
// 6 - Car Telemetry
// 7 - Car Status

type PacketHeader struct {
	MPacketFormat  uint16  // 2018
	MPacketVersion uint8   // Version of this packet type, all start from 1
	MPacketID      uint8   // Identifier for the packet type, see below
	MSessionUID    uint64  // Unique identifier for the session
	MSessionTime   float32 // Session timestamp
	// MFrameIdentifier uint  // Identifier for the frame the data was retrieved on
	MPlayerCarIndex uint8 // Index of player's car in the array
}

func NewPacketHeader(buf []byte) PacketHeader {
	var pack PacketHeader
	r := bytes.NewReader(buf)
	if err := binary.Read(r, binary.LittleEndian, &pack); err != nil {
		log.Fatal(err)
	}
	return pack
}
