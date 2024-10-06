package rabbit

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"time"
)

func Send(logger *zap.SugaredLogger, login string, id int) error {
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//if err != nil {
	//	logger.Error(err, "Failed to connect to RabbitMQ")
	//	return err
	//}
	//defer conn.Close()

	//ch, err := Conn.Channel()
	//if err != nil {
	//	logger.Error(err, "Failed to open a channel")
	//	return err
	//}
	//defer ch.Close()

	q, err := ch.QueueDeclare(
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

	type message struct {
		Login string `json:"login"`
		Id    int    `json:"id"`
	}
	m := new(message)
	m.Login = login
	m.Id = id
	body, err := json.Marshal(m)

	err = ch.PublishWithContext(ctx,
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
