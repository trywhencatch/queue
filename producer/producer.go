package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	conn, err := net.DialTimeout("tcp", "localhost:9001", 180*time.Second)
	if err != nil {
		fmt.Printf("Error while connecting %v\n", err)
		return
	}

	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter message to publish")
		input := ""

		if scanner.Scan() {
			input = scanner.Text()
			fmt.Printf("You typed: %s\n", input)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		message := input + "\n"
		_, err = conn.Write([]byte(message))

		fmt.Println("message published")

		if err != nil {
			fmt.Printf("Failed to write to queue %v", err)
		}
	}

}
