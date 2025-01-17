package rabbitmq

import (
	"github.com/FinnTew/FinnEcommerce/src/internal/service/email/conf"
	"github.com/streadway/amqp"
)

var (
	Conn    *amqp.Connection
	Channel *amqp.Channel
	err     error
)

func Init() {
	Conn, err = amqp.Dial(
		conf.GetConf().RabbitMQ.URL,
	)
	if err != nil {
		panic(err)
	}
	Channel, err = Conn.Channel()
	if err != nil {
		panic(err)
	}
}
