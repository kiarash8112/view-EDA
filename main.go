package main

import (
	"log"
	"time"

	"github.com/lovoo/goka"
)

func main() {
	tm, err := goka.NewTopicManager(brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	defer tm.Close()
	err = tm.EnsureStreamExists(string(payment_topic), 8)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", payment_topic, err)
	}
	err = tm.EnsureStreamExists(string(booking_topic), 8)
	if err != nil {
		log.Printf("Error creating kafka topic %s: %v", payment_topic, err)
	}
	go runPaymentEmitter()
	go runPaymentProcessor()
	go runBookingEmitter()
	go runBookingProcessor()
	time.Sleep(time.Second * 2)

	runView()
}
