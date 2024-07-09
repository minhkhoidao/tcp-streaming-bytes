package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func generateRandomInt16Array(size int, min, max int16) []int16 {
	rand.Seed(time.Now().UnixNano()) // Initialize the global random generator with a time-based seed.

	var arr []int16
	for i := 0; i < size; i++ {
		// Generate a random int16 value between min and max.
		num := min + int16(rand.Intn(int(max-min+1)))
		arr = append(arr, num)
	}

	return arr
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for {
		int16Rand := generateRandomInt16Array(480000, -2426, 2812)
		int16Rand2 := []int16{8, 9, 10, 11, 12, 13}
		int16Rand3 := []int16{14, 15, 16, 17, 18, 19}
		int16Rand4 := []int16{8, 9, 10, 11, 12, 13}
		byteData := int16ToBytes(int16Rand)
		byteData3 := int16ToBytes(int16Rand2)
		byteData4 := int16ToBytes(int16Rand3)
		byteData5 := int16ToBytes(int16Rand4)

		// Send each byte slice once
		if _, err := conn.Write(byteData3); err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
		if _, err := conn.Write(byteData4); err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
		if _, err := conn.Write(byteData5); err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
		if _, err := conn.Write(byteData); err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
		time.Sleep(3 * time.Second)
	}
}

func int16ToBytes(int16s []int16) []byte {
	byteData := make([]byte, len(int16s)*2)
	for i, v := range int16s {
		binary.LittleEndian.PutUint16(byteData[i*2:], uint16(v))
	}
	return byteData
}
