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

type userProvider struct {
	service core.UserService
}

func (p *userProvider) provide(builder *schemabuilder.Schema) {
	builder.Query().FieldFunc("user", func(ctx context.Context, args struct{ Id *string }) ([]core.User, error) {
		if args.Id != nil {
			user, err := p.service.Get(ctx, *args.Id)
			if err != nil {
				return nil, err
			}
			return []core.User{*user}, nil
		}
		return p.service.List(ctx)
	})
}
