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
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
)

type locationQueryProvider struct {
	service core.LocationService
	logger  logrus.FieldLogger
}

func (l *locationQueryProvider) provide(builder *schemabuilder.Schema) {
	query := builder.Query()
	query.FieldFunc("location", func(ctx context.Context, args struct{ Id *string }) ([]core.Location, error) {
		if id := args.Id; id != nil {
			location, err := l.service.Get(ctx, *id)
			if err != nil {
				return nil, err
			}
			return []core.Location{*location}, nil
		}
		return l.service.List(ctx)
	})

	inventory := builder.Object("Inventory", core.Inventory{})
	inventory.FieldFunc("location", func(ctx context.Context, i *core.Inventory) (*core.Location, error) {
		return l.service.Get(ctx, i.Location)
	}, schemabuilder.NonNullable)
}
