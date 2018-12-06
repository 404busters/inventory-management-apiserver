/*
	Copyright 2018 Carmen Chan & Tony Yip

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package core

import (
	"context"
	"time"
)

type Inventory struct {
	Id           string          `json:"id,omitempty"`
	ItemType     string          `json:"itemType" graphql:"-"`
	Location     string          `json:"location" graphql:"-"`
	LastSeenTime time.Time       `json:"lastSeenTime,omitempty" graphql:"-"`
	Status       InventoryStatus `json:"status"`
}

type InventoryStatus string

const (
	StatusStock     InventoryStatus = "STOCK"
	StatusInUse     InventoryStatus = "IN_USE"
	StatusRepair    InventoryStatus = "REPAIR"
	StatusTransport InventoryStatus = "TRANSPORT"
)

type InventoryService interface {
	ItemTypeList(ctx context.Context, itemTypeId string) ([]Inventory, error)
	LocationList(ctx context.Context, locationId string) ([]Inventory, error)
	Get(ctx context.Context, id string) (*Inventory, error)
	Create(ctx context.Context, input *Inventory) (*Inventory, error)
	Update(ctx context.Context, id string, input *Inventory) (*Inventory, error)
	Delete(ctx context.Context, id string) error
}
