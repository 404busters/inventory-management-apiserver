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
	Id           string              `json:"id,omitempty"`
	ItemType     string              `json:"itemType"`
	Location     string              `json:"location"`
	LastSeenTime time.Time           `json:"lastSeenTime"`
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
