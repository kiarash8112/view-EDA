package main

import (
	"fmt"
	"time"

	"github.com/lovoo/goka"
)

var (
	brokers                   = []string{"127.0.0.1:9092"}
	payment_topic goka.Stream = "payment1"
	booking_topic goka.Stream = "booking1"
	tmc           *goka.TopicManagerConfig
)

func init() {
	tmc = goka.NewTopicManagerConfig()
	tmc.Table.Replication = 2
	tmc.Stream.Replication = 2
}

func runPaymentEmitter() {
	emitter, err := goka.NewEmitter(brokers, payment_topic,
		new(PaymentCodec))
	if err != nil {
		panic(err)
	}

	t := time.NewTicker(time.Millisecond*100 + time.Millisecond*10)
	defer t.Stop()

	var i int
	for range t.C {
		key := fmt.Sprintf("user-%d", i%10)
		emitter.EmitSync(key, &_PaymentService)
		i++
	}

	defer emitter.Finish()
}

func runBookingEmitter() {
	emitter, err := goka.NewEmitter(brokers, booking_topic,
		new(BookingCodec))
	if err != nil {
		panic(err)
	}

	t := time.NewTicker(time.Millisecond*100 + time.Millisecond*10)
	defer t.Stop()

	var i int
	for range t.C {
		key := fmt.Sprintf("user-%d", i%10)
		emitter.EmitSync(key, &_BookingService)
		i++
	}

	defer emitter.Finish()
}
