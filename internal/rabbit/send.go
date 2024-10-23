package rabbit

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"time"
)

type message struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
}

func (p *Publisher) Send(logger *zap.SugaredLogger, login string, id int) error {

	q, err := p.ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		logger.Error(err, "Failed to declare a queue")
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := new(message)
	m.Login = login
	m.Id = id
	body, err := json.Marshal(m)

	err = p.ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		logger.Error(err, "Failed to publish a message")
		return err
	}
	log.Printf(" [x] Sent %s\n", body)
	return nil
}
