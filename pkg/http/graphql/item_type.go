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

type itemTypeQueryProvider struct {
	service core.ItemTypeService
}

func (p *itemTypeQueryProvider) provide(builder *schemabuilder.Schema) {
	query := builder.Query()
	query.FieldFunc("itemType", func(ctx context.Context, args struct{ Id *string }) ([]core.ItemType, error) {
		if id := args.Id; id != nil {
			entry, err := p.service.Get(ctx, *id)
			if err != nil {
				return nil, err
			} else if entry == nil {
				return nil, nil
			}
			return []core.ItemType{*entry}, nil
		}
		return p.service.List(ctx)
	})
}
