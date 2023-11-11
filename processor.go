package main

import (
	"context"

	"github.com/lovoo/goka"
)

var (
	payment_group goka.Group = "payment_group1"
	booking_group goka.Group = "booking_group1"
)

func paymentProcess(ctx goka.Context, msg interface{}) {
	var p1 []PaymentService
	if v := ctx.Value(); v != nil {
		p1 = v.([]PaymentService)
	}

	p := msg.(*PaymentService)
	p1 = append(p1, *p)

	ctx.SetValue(p1)
}

func runPaymentProcessor() {
	g := goka.DefineGroup(payment_group,
		goka.Input(payment_topic, new(PaymentCodec), paymentProcess),
		goka.Persist(new(PaymentListCodec)),
	)
	p, err := goka.NewProcessor(brokers,
		g,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
	)
	if err != nil {
		panic(err)
	}

	p.Run(context.Background())
}

func BookingProcess(ctx goka.Context, msg interface{}) {
	var b1 []BookingService
	if v := ctx.Value(); v != nil {
		b1 = v.([]BookingService)
	}

	b := msg.(*BookingService)
	b1 = append(b1, *b)

	ctx.SetValue(b1)
}

func runBookingProcessor() {
	g := goka.DefineGroup(booking_group,
		goka.Input(booking_topic, new(BookingCodec), BookingProcess),
		goka.Persist(new(BookingListCodec)),
	)
	p, err := goka.NewProcessor(brokers,
		g,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
	)
	if err != nil {
		panic(err)
	}

	p.Run(context.Background())
}
