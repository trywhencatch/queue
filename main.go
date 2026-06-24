package main

import (
	"distribute-job-queue/src/broker"
	"distribute-job-queue/src/connection-manager"
)



func main() {
	broker := broker.NewBroker()
	broker.CreateQueue("test_queue")
	connectionmanager.OpenConnection(broker.HandleClient)
}