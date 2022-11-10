/*
 * skogul, complex receiver tests
 *
 * Copyright (c) 2019 Telenor Norge AS
 * Author(s):
 *  - Roshini Narasimha Raghavan <roshiragavi@gmail.com>
 *  - Kristian Lyngstøl <kly@kly.no>
 *
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301  USA
 */

package receiver_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/telenornms/skogul"
	"github.com/telenornms/skogul/receiver"
	"github.com/telenornms/skogul/sender"
)

// Tests RabbitMq receiver, sender and JSON parser implicitly
func TestRabbitMq(t *testing.T) {
	// create a data container to be sent to RabbitMQ
	by := []byte("{\"metrics\":[{\"timestamp\":\"0001-01-01T00:00:00Z\",\"metadata\":{\"key\":\"value\"}}]}")
	var data_container *skogul.Container
	err := json.Unmarshal(by, &data_container)
	if err != nil {
		t.Errorf("Failed to load config: %v", err)
		return
	}
	senderq := sender.Rabbitmq{
		Sync:      false,
		Address:   "localhost:5672",
		Username:  "guest",
		Password:  "guest",
		Encoder:   skogul.EncoderRef{},
		TLS:       false,
		Queuename: "tpoll-results",
	}
	// send a data container to Rabbitmq

	//senderq.Encoder.E.Encode(data_container)
	senderq.Send(data_container)
	// prepare the receiver
	receiverq := receiver.Rabbitmq{
		Username:  "guest",
		Password:  "guest",
		Address:   "localhost:5672",
		Queuename: "tpoll-results",
	}
	// start the receiver to receive the container.
	go receiverq.Start()
	time.Sleep(time.Duration(100 * time.Millisecond))
	// check if the receiver has received the container

}
