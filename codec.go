package main

import (
	"encoding/json"
	"fmt"
)

type Status string

const (
	success     Status = "SUCCESS"
	in_progress Status = "IN_PROGRESS"
	faild       Status = "FAILD"
)

type BookingService struct {
	BookingID     string
	HotelID       string
	BookingStatus Status
}

type BookingCodec struct{}

// Encodes a BookingService into []byte
func (*BookingCodec) Encode(value interface{}) ([]byte, error) {
	if _, ok := value.(*BookingService); !ok {
		return nil, fmt.Errorf("Codec requires value *BookingService, got %T", value)
	}
	return json.Marshal(value)
}

// Decodes a BookingService from []byte to it's go representation.
func (*BookingCodec) Decode(data []byte) (interface{}, error) {
	var (
		b   BookingService
		err error
	)
	err = json.Unmarshal(data, &b)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling BookingService: %v", err)
	}
	return &b, nil
}

type PaymentService struct {
	BookingID     string
	PaymentStatus Status
}

type PaymentCodec struct{}

// Encodes a BookingService into []byte
func (*PaymentCodec) Encode(value interface{}) ([]byte, error) {
	if _, ok := value.(*PaymentService); !ok {
		return nil, fmt.Errorf("Codec requires value *PaymentService, got %T", value)
	}
	return json.Marshal(value)
}

// Decodes a BookingService from []byte to it's go representation.
func (*PaymentCodec) Decode(data []byte) (interface{}, error) {
	var (
		p   PaymentService
		err error
	)
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling PaymentService: %v", err)
	}
	return &p, nil
}

type ViewResult struct {
	BookingID     string
	HotelID       string
	BookingStatus Status
	PaymentStatus Status
}
