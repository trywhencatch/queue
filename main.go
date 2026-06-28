package main

import (
	"distributed-job-queue/internals"
	"distributed-job-queue/broker"
)

func main() {
	exchange := broker.NewExchange()
	exchange.CreateQueue("test_queue")
	go internals.OpenProducerConnection(exchange.HandleProducer)
	internals.OpenConsumerConnection(exchange.HandleConsumer)
}
