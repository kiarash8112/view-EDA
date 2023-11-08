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
	p.BookingID = msg.(*PaymentService).BookingID
	p.PaymentStatus = msg.(*PaymentService).PaymentStatus
	ctx.SetValue(p)
}

func runPaymentProcessor() {
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
	b.BookingID = msg.(*BookingService).BookingID
	b.BookingStatus = msg.(*BookingService).BookingStatus
	b.HotelID = msg.(*BookingService).HotelID
	ctx.SetValue(b)
}

func runBookingProcessor() {
	g := goka.DefineGroup(booking_group,
		goka.Input(booking_topic, new(BookingCodec), BookingProcess),
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
