package core

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Inventory struct {
	Id           string              `json:"id,omitempty"`
	ItemType     string              `json:"itemType"`
	Location     string              `json:"location"`
	LastSeenTime timestamp.Timestamp `json:"lastSeenTime"`
	Status       StatusEnum          `json:"status"`
}

type StatusEnum string

const (
	STOCK     StatusEnum = "STOCK"
	IN_USE    StatusEnum = "IN_USE"
	REPAIR    StatusEnum = "REPAIR"
	TRANSPORT StatusEnum = "TRANSPORT"
)

type InventoryService interface {
	Get(ctx context.Context, id string) (*Inventory, error)
	Create(ctx context.Context, input *Inventory) (*Inventory, error)
	Update(ctx context.Context, id string, input *Inventory) (*Inventory, error)
	Delete(ctx context.Context, id string) error
}
