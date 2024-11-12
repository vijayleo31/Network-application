package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Start the server on localhost:8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on localhost:8080")

	for {
		// Accept new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		// Handle connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewScanner(conn)

	for reader.Scan() {
		message := reader.Text()
		fmt.Println("Received message:", message)

		// Echo message back to client
		_, err := conn.Write([]byte("Echo: " + message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}

	if err := reader.Err(); err != nil {
		fmt.Println("Connection error:", err)
	}
}
