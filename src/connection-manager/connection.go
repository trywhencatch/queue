package connectionmanager

import (
	"fmt"
	"net"
)

func OpenConnection(handler func(net.Conn)) {
	listener, err := net.Listen("tcp","localhost:9001")

	if err != nil {
		fmt.Println("Failed to connect",err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is listening on port 9001")

	for{
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept",err)
			continue
		}

		go handler(conn)
	}
}