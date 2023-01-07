package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	buf := make([]byte, 1024)

	if _, err := conn.Read(buf); err != nil {
		fmt.Println("error reading from client: ", err.Error())
		os.Exit(1)
	}

	// let's ingore the client's inpt for now and hardcode a response.
	// we'll implement a proper Redis Protocol parse in later stages.
	conn.Write([]byte("+PONG\r\n"))

}
