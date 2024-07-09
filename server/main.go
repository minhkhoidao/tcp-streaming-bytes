package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	var buffer []byte
	tempBuf := make([]byte, 480000)

	for {
		n, err := reader.Read(tempBuf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		buffer = append(buffer, tempBuf[:n]...)
		int16Data := bytesToInt16(buffer)

		// Print the data and reset the buffer
		if len(int16Data) > 0 {
			fmt.Printf("Received: %v\n", int16Data)
			buffer = buffer[len(int16Data)*2:]
		}
	}
}

func bytesToInt16(b []byte) []int16 {
	length := len(b) / 2
	int16s := make([]int16, length)
	for i := 0; i < length; i++ {
		int16s[i] = int16(binary.LittleEndian.Uint16(b[i*2:]))
	}
	return int16s
}
