package infrastructure

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Amqp struct {
	Conn  *amqp.Connection
	Ch    *amqp.Channel
	Queue *amqp.Queue
}

func NewAmqp() *Amqp {
	Amqp := &Amqp{}

	conn, err := amqp.Dial("amqp://root:root@localhost:5672/")
	if err != nil {
		panic(err)
	}

	Amqp.Conn = conn

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	Amqp.Ch = ch

	queue, err := ch.QueueDeclare(
		"ReceivedMessage", // name
		false,             // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		panic(err)
	}

	Amqp.Queue = &queue

	return Amqp
}
