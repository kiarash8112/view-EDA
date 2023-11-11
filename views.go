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
		new(PaymentCodec),
	)
	if err != nil {
		panic(err)
	}

	bookingView, err := goka.NewView(brokers,
		goka.GroupTable(booking_group),
		new(BookingCodec),
	)
	if err != nil {
		panic(err)
	}

	root := mux.NewRouter()
	root.HandleFunc("/{key}", func(w http.ResponseWriter, r *http.Request) {
		view := ViewResult{}
		payment, _ := paymentView.Get(mux.Vars(r)["key"])
		paymentService := payment.(*PaymentService)

		view.BookingID = paymentService.BookingID
		view.PaymentStatus = paymentService.PaymentStatus

		booking, _ := bookingView.Get(mux.Vars(r)["key"])
		bookingService := booking.(*BookingService)

		view.HotelID = bookingService.HotelID
		view.BookingStatus = bookingService.BookingStatus

		data, err := json.Marshal(view)
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
