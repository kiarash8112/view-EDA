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
		payment, _ := paymentView.Get(mux.Vars(r)["key"])
		data, _ := json.Marshal(payment)
		w.Write(data)

		payment, _ = bookingView.Get(mux.Vars(r)["key"])
		data, _ = json.Marshal(payment)
		w.Write(data)
	})
	fmt.Println("View opened at http://localhost:9095/")
	go http.ListenAndServe(":9095", root)

	go paymentView.Run(context.Background())
	bookingView.Run(context.Background())
}
