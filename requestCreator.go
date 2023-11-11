package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hashicorp/go-uuid"
)

var _BookingService BookingService
var _PaymentService PaymentService

func createRequest() {
	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()

	for range t.C {
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

		_BookingService = BookingService{
			HotelID:       hotelID,
			BookingID:     bookingID,
			BookingStatus: bookingStatus[random],
		}

		random = rand.Intn(3)
		paymentStatus := map[int]Status{
			0: success,
			1: in_progress,
			2: faild,
		}

		_PaymentService = PaymentService{
			BookingID:     bookingID,
			PaymentStatus: paymentStatus[random],
		}
	}
}
