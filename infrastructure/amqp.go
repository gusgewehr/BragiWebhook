package infrastructure

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Amqp struct {
	Conn  *amqp.Connection
	Ch    *amqp.Channel
	Queue *amqp.Queue
}

func NewAmqp(env Env) *Amqp {
	Amqp := &Amqp{}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", env.AmqpUser, env.AmqpPassword, env.AmqpHost, env.AmqpPort))
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
		env.QueueName, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		panic(err)
	}

	Amqp.Queue = &queue

	return Amqp
}
