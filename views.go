package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
)

func runView() {
	paymentView, err := goka.NewView(brokers,
		goka.GroupTable(payment_group),
		new(PaymentListCodec),
	)
	if err != nil {
		panic(err)
	}

	bookingView, err := goka.NewView(brokers,
		goka.GroupTable(booking_group),
		new(BookingListCodec),
	)
	if err != nil {
		panic(err)
	}

	root := mux.NewRouter()
	root.HandleFunc("/{key}", func(w http.ResponseWriter, r *http.Request) {
		views := []ViewResult{}
		payments, _ := paymentView.Get(mux.Vars(r)["key"])
		paymentInstances := payments.([]PaymentService)

		bookings, _ := bookingView.Get(mux.Vars(r)["key"])
		bookingInstances := bookings.([]BookingService)

		fmt.Println(len(paymentInstances), len(bookingInstances))
		for index, payment := range paymentInstances {
			view := ViewResult{}
			view.BookingID = payment.BookingID
			view.PaymentStatus = payment.PaymentStatus
			if !(payment.PaymentStatus == faild) {
				view.HotelID = bookingInstances[index].HotelID
				view.BookingStatus = bookingInstances[index].BookingStatus
			} else {
				view.HotelID = bookingInstances[index].HotelID
				view.BookingStatus = faild
			}
			views = append(views, view)
		}

		data, err := json.Marshal(views)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		w.Write(data)
	})
	fmt.Println("View opened at http://localhost:9095/")
	go http.ListenAndServe(":9095", root)

	go paymentView.Run(context.Background())
	bookingView.Run(context.Background())
}
