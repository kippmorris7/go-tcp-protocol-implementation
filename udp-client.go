package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:1025")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	udpConn, err := net.ListenPacket("udp", "localhost:0")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer udpConn.Close()
	readyToExit := false
	scanner := bufio.NewScanner(os.Stdin)

	for !readyToExit {
		fmt.Println("Type a message to send to the server and then press enter (type \"exit\" as your message to end the client program):")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			readyToExit = true
		} else {
			_, err := udpConn.WriteTo([]byte(input), udpAddr)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Sent the message to the server!")
		}
	}
}
