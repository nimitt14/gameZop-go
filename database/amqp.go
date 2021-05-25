package database

import (
    "log"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
  }

func NewAmqpClient() *amqp.Channel {
	conn, err := amqp.Dial("amqps://nmvgccnx:3FTTOuT2OGx9L9MJcpLVWDtVJIW9al9g@baboon.rmq.cloudamqp.com/nmvgccnx")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch	
}

