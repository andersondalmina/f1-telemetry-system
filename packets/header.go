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
	PacketFormat    uint16  // 2018
	PacketVersion   uint8   // Version of this packet type, all start from 1
	PacketID        uint8   // Identifier for the packet type, see below
	SessionUID      uint64  // Unique identifier for the session
	SessionTime     float32 // Session timestamp
	FrameIdentifier uint32  // Identifier for the frame the data was retrieved on
	PlayerIndex     uint8   // Index of player's car in the array
}

func NewPacketHeader(buffer []byte) PacketHeader {
	var pack PacketHeader
	r := bytes.NewReader(buffer)
	if err := binary.Read(r, binary.LittleEndian, &pack); err != nil {
		log.Fatal(err)
	}
	return pack
}
