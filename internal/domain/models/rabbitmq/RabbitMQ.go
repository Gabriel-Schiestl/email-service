package rabbitmq

import (
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp091.Connection
	ch *amqp091.Channel
	queue amqp091.Queue
}

func NewRabbitMQ(queueName string) *RabbitMQ {
	rmq := &RabbitMQ{}

	conn, err := amqp091.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatalf("Falha ao conectar ao servidor AMQP: %v", err)
	}
	rmq.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir canal: %v", err)
	}
	rmq.ch = ch

	queue, err := ch.QueueDeclare(
		queueName,
		true,
		false, 
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Falha ao declarar fila: %v", err)
	}
	rmq.queue = queue

	return rmq
}

func (r *RabbitMQ) Close() {
	if r.ch != nil {
		r.ch.Close()
	}

	if r.conn != nil {
		r.conn.Close()
	}
}

func (r *RabbitMQ) Consume() (<-chan amqp091.Delivery, error) {
	msgs, err := r.ch.Consume(
		r.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}