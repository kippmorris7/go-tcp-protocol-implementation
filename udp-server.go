package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	udpConn, err := net.ListenPacket("udp", "localhost:1025")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer udpConn.Close()
	fiveSecondDuration, _ := time.ParseDuration("5s")
	udpConn.SetReadDeadline(time.Now().Add(fiveSecondDuration))

	for {
		buffer := make([]byte, 64, 64)
		bytesRead, senderAddr, _ := udpConn.ReadFrom(buffer)

		if bytesRead > 0 {
			fmt.Println("Received a message from ", senderAddr, ":")
			fmt.Println(string(buffer))
		}
	}
}
