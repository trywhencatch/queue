package internals

import (
	"fmt"
	"net"
	"strconv"
)

func OpenConnection(handler func(net.Conn), port int) {

	url := "localhost:" + strconv.Itoa(port)
	listener, err := net.Listen("tcp",url)

	if err != nil {
		fmt.Println("Failed to connect",err)
		return
	}

	defer listener.Close()
	fmt.Println("Server is listening on port: ",port)

	for{
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept",err)
			continue
		}

		go handler(conn)
	}
}

func OpenProducerConnection(handler func(net.Conn)){
	OpenConnection(handler, 9001)
}

func OpenConsumerConnection(handler func(net.Conn)){
	OpenConnection(handler, 9002)
}