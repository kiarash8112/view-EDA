package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

var (
	brokers                   = []string{"127.0.0.1:9092"}
	payment_topic goka.Stream = "payment"
	booking_topic goka.Stream = "booking"
	tmc           *goka.TopicManagerConfig
)

func init() {
	tmc = goka.NewTopicManagerConfig()
	tmc.Table.Replication = 2
	tmc.Stream.Replication = 2
}

func runPaymentEmitter() {
	emitter, err := goka.NewEmitter(brokers, payment_topic,
		new(codec.String))
	if err != nil {
		panic(err)
	}

	status := rand.Intn(3)
	paymentStatus := map[int]Status{
		0: success,
		1: in_progress,
		2: faild,
	}

	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

	var i int
	for range t.C {
		key := fmt.Sprintf("user-%d", i%10)
		value := paymentStatus[status]
		emitter.EmitSync(key, value)
		i++
	}

	defer emitter.Finish()
}

func runBookingEmitter() {
	emitter, err := goka.NewEmitter(brokers, booking_topic,
		new(codec.String))
	if err != nil {
		panic(err)
	}
	status := rand.Intn(3)
	bookingStatus := map[int]Status{
		0: success,
		1: in_progress,
		2: faild,
	}

	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

	var i int
	for range t.C {
		key := fmt.Sprintf("user-%d", i%10)
		value := bookingStatus[status]
		emitter.EmitSync(key, value)
		i++
	}

	defer emitter.Finish()
}
