package main

import (
	"fmt"
	"log"
	"net"

	"github.com/andersondalmina/f1-telemetry/persist"
	"github.com/andersondalmina/f1-telemetry/services"
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

	persist.CreateClient()
	defer persist.CloseClient()

	buffer := make([]byte, 1341)
	for {
		_, err := con.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		services.ProcessPacket(buffer)
	}
}
