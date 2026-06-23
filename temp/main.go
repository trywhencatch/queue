package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}
		fmt.Printf("Recieved %s", data)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		fmt.Println("err ", err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is listening on port 9001")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err", err)
			continue
		}

		go handleClient(conn)
	}

}
