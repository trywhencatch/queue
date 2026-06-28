package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:9002", 120*time.Second)
	if err != nil {
		fmt.Printf("Failed to connect to localhost:9002 %v \n", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Failed to read data %v \n", err)
			fmt.Printf("Stopping consumer...\n")
			break;
		}
		fmt.Println("------")
		fmt.Println(string(data))
		fmt.Println("------")
	}
}
