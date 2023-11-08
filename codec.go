package main

import (
	"encoding/json"
	"fmt"
)

type status string

const (
	success     status = "SUCCESS"
	in_progress status = "IN_PROGRESS"
	faild       status = "FAILD"
)

type BookingService struct {
	BookingID     string
	PaymentStatus status
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
