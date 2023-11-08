package main

import (
	"context"

	"github.com/lovoo/goka"
)

var (
	payment_group goka.Group = "payment_group"
	booking_group goka.Group = "booking_group"
)

func paymentProcess(ctx goka.Context, msg interface{}) {
	var p *PaymentService
	if val := ctx.Value(); val != nil {
		p = val.(*PaymentService)
	} else {
		p = new(PaymentService)
	}

	ctx.SetValue(p)
}

func runPaymentProcessor(initialized chan struct{}) {
	g := goka.DefineGroup(payment_group,
		goka.Input(payment_topic, new(PaymentCodec), paymentProcess),
		goka.Persist(new(PaymentCodec)),
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
	var b *BookingService
	if val := ctx.Value(); val != nil {
		b = val.(*BookingService)
	} else {
		b = new(BookingService)
	}

	ctx.SetValue(b)
}

func runBookingProcessor(initialized chan struct{}) {
	g := goka.DefineGroup(payment_group,
		goka.Input(payment_topic, new(BookingCodec), paymentProcess),
		goka.Persist(new(BookingCodec)),
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
