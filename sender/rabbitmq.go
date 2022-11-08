package sender

import (
	"context"
	"encoding/json"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/telenornms/skogul"
	"github.com/telenornms/skogul/encoder"
)

var rabbitmqLog = skogul.Logger("sender", "rabbitmq")

type Rabbitmq struct {
	Sync        bool
	Address     string
	Username    string
	Password    string
	Encoder     skogul.EncoderRef
	once        sync.Once
	dial_string string
	TLS         bool
	conn        *amqp.Connection
	Queuename   string
}

/* TLS is not used in the initial version (MVP) but to go to production we may want to consider using TLS */

func (r *Rabbitmq) init() {
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

}

func (r *Rabbitmq) Send(c *skogul.Container) error {
	r.once.Do(func() {
		r.init()
	})
	ch, err := r.conn.Channel()
	if err != nil {
		rabbitmqLog.Warnf("channel creation error %s", err)
	}

	defer ch.Close()

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
		rabbitmqLog.Info("Queue created successfully %s", q)
	}

	if r.Encoder.Name == "" {
		r.Encoder.E = encoder.JSON{}
	}

	ctx := context.Background()
	b, err := json.Marshal(*c)
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
	defer r.conn.Close()
	if err != nil {
		return err
	} else {
		return nil
	}

}
