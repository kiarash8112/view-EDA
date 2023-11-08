package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hashicorp/go-uuid"
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
		new(PaymentCodec))
	if err != nil {
		panic(err)
	}

	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()

	var i int
	for range t.C {
		key := fmt.Sprintf("user-%d", i%10)
		value := createNewPaymentService()
		emitter.EmitSync(key, &value)
		i++
	}

	defer emitter.Finish()
}

func createNewPaymentService() PaymentService {
	random := rand.Intn(3)
	paymentStatus := map[int]Status{
		0: success,
		1: in_progress,
		2: faild,
	}

	bookingID, err := uuid.GenerateUUID()
	if err != nil {
		log.Panic("can't create random bookingID")
	}

	return PaymentService{
		BookingID:     bookingID,
		PaymentStatus: paymentStatus[random],
	}

}

func runBookingEmitter() {
	emitter, err := goka.NewEmitter(brokers, booking_topic,
		new(codec.String))
	if err != nil {
		panic(err)
	}
	random := rand.Intn(3)
	bookingStatus := map[int]Status{
		0: success,
		1: in_progress,
		2: faild,
	}

	bookingID, err := uuid.GenerateUUID()
	if err != nil {
		log.Panic("can't create random bookingID")
	}

	hotelID, err := uuid.GenerateUUID()
	if err != nil {
		log.Panic("can't create random bookingID")
	}

	b := BookingService{
		HotelID:       hotelID,
		BookingID:     bookingID,
		BookingStatus: bookingStatus[random],
	}

	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

	var i int
	for range t.C {
		key := fmt.Sprintf("user-%d", i%10)
		emitter.EmitSync(key, &b)
		i++
	}

	defer emitter.Finish()
}
