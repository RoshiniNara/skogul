package receiver

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/telenornms/skogul"
)

var rabbitmqLog = skogul.Logger("sender", "rabbitmq")

type Rabbitmq struct {
	Address     string
	Username    string
	Password    string
	dial_string string
	TLS         bool
	conn        *amqp.Connection
	Queuename   string
	Handler     *skogul.HandlerRef
}

/* Start the rabbitmq and never return */

func (r *Rabbitmq) Start() {
	if r.Username != "" && r.Password != "" {
		r.dial_string = "amqp://" + r.Username + ":" + r.Password + "@" + r.Address
	} else {
		rabbitmqLog.Warnf("Username %s or Password %s for the rabbit mq is empty", r.Username, r.Password)
	}
	var err error
	r.conn, err = amqp.Dial(r.dial_string)
	if err != nil {
		rabbitmqLog.Warnf("Dial error %s", err)
	}

	ch, err := r.conn.Channel()
	if err != nil {
		rabbitmqLog.Warnf("channel creation error %s", err)
	}

	/* Queue declaration is needed on the receiver as well since it is possible that receiver gets started before the sender */
	q, err := ch.QueueDeclare(
		r.Queuename,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		rabbitmqLog.Warnf("Queue declaration error %s", err)
	} else {
		rabbitmqLog.Info("Queue already exists or created successfully")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		rabbitmqLog.WithError(err).Warn("Unable to consume")
	}

	var forever chan struct{}

	go func() {
		var err error
		for d := range msgs {
			err = r.Handler.H.Handle(d.Body)
			if err != nil {
				rabbitmqLog.WithError(err).Warn("Unable to handle RabbitMQ message")
			}
		}
	}()
	<-forever
	defer r.conn.Close()
	defer ch.Close()
}
