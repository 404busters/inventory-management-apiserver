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

package graphql

import (
	"context"

	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
)

type inventoryProvider struct {
	service core.InventoryService
}

func (p *inventoryProvider) provide(builder *schemabuilder.Schema) {
	builder.Enum(core.StatusStock, map[string]interface{}{
		"STOCK":     core.StatusStock,
		"IN_USE":    core.StatusInUse,
		"REPAIR":    core.StatusRepair,
		"TRANSPORT": core.StatusTransport,
	})

	builder.Query().FieldFunc("inventory", func(ctx context.Context, args struct{ Id string }) (*core.Inventory, error) {
		return p.service.Get(ctx, args.Id)
	})

	location := builder.Object("Location", core.Location{})
	location.FieldFunc("inventory", func(ctx context.Context, l *core.Location) ([]core.Inventory, error) {
		return p.service.LocationList(ctx, l.Id)
	})

	itemType := builder.Object("ItemType", core.ItemType{})
	itemType.FieldFunc("inventory", func(ctx context.Context, i *core.ItemType) ([]core.Inventory, error) {
		return p.service.ItemTypeList(ctx, i.Id)
	})
}
