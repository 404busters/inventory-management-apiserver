package core

import "context"

type Transport struct {
	Id             string    `json:"id,omitempty"`
	PersonInCharge string    `json:"person_in_charge"`
	Location       string    `json:"location"`
	EventType      EventType `json:"event_type"`
	Note           string    `json:"note"`
}

type EventType string

const (
	CheckIn  EventType = "CHECK_IN"
	CheckOut EventType = "CHECK_OUT"
)

type TransportService interface {
	List(ctx context.Context) ([]Transport, error)
	FilterList(ctx context.Context, filter struct {
		User      string
		Inventory string
	}) ([]Transport, error)
	CheckIn(ctx context.Context, input *Transport) (*Transport, error)
	CheckOut(ctx context.Context, input *Transport) (*Transport, error)
}
