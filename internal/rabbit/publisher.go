package rabbit

import "C"
import (
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type Publisher struct {
	Logger     *zap.SugaredLogger
	RabbitConn string
	ch         *amqp.Channel
	conn       *amqp.Connection
}

func New(logger *zap.SugaredLogger, rabbitConn string) *Publisher {
	return &Publisher{
		Logger:     logger,
		RabbitConn: rabbitConn,
	}
}

func (p *Publisher) Connect() (err error) {

	p.conn, err = amqp.Dial(p.RabbitConn)
	if err != nil {
		p.Logger.Error(err, "Failed to connect to RabbitMQ")
		return err
	}
	//defer conn.Close()

	p.ch, err = p.conn.Channel()
	if err != nil {
		p.Logger.Error(err, "Failed to open a channel")
		return err
	}
	//defer ch.Close()

	return nil
}

//func (p *Publisher) Close() {
//	// итать про * в методах
//	err := p.ch.Close()
//	if err != nil {
//		return
//	}
//	err = p.conn.Close()
//	if err != nil {
//		return
//	}
//}
