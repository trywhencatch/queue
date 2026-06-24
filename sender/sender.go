package main

import (
	"fmt"
	"net"
	"time"
)


func main(){

	conn, err := net.DialTimeout("tcp","localhost:9001", 180*time.Second)

	if err != nil {
		fmt.Printf("Error while connecting %v\n", err)
		return 
	}

	message := "Hello there\n"
	conn.Write([]byte(message))
}