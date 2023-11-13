package main

import (
	"fmt"
	"net"
)

func main() {
	// Create a listener on port 8080.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept new connections and handle them in separate goroutines.
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Read data from the client.
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write data back to the client.
	conn.Write(buffer[:n])

	// Close the connection.
	conn.Close()
}
