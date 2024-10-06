package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Publisher struct {
	logger     *zap.SugaredLogger
	rabbitConn string
}

func New(logger *zap.SugaredLogger, rabbitConn string) *Publisher {
	return &Publisher{
		logger:     logger,
		rabbitConn: rabbitConn,
	}
}

var (
	ch *amqp.Channel
)

func PublisherRun(logger *zap.SugaredLogger, p *Publisher) error {
	conn, err := amqp.Dial(p.rabbitConn)
	if err != nil {
		logger.Error(err, "Failed to connect to RabbitMQ")
		return err
	}
	//defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		logger.Error(err, "Failed to open a channel")
		return err
	}
	//defer ch.Close()

	return nil
}
