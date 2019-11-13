package main

import (
	"fmt"
	"log"
	"net"

	"github.com/andersondalmina/f1-telemetry/packets"
)

func main() {
	fmt.Println("Starting the server at port 20777")
	laddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:20777")
	if err != nil {
		log.Fatal(err)
	}

	con, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	fmt.Println("Server started")

	buf := make([]byte, 1289)

	for {
		_, err := con.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		// tp, err := NewTelemetryPack(buf)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		header := packets.NewPacketHeader(buf)

		if header.MPacketID == 6 {
			teste := packets.NewPacketCarTelemetryData(buf)
			fmt.Println(teste)
		}
		// fmt.Println(header)
	}
}
