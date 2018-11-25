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
	"net/http"

	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/http/inject"
)

type graphQLProvider interface {
	provide(builder *schemabuilder.Schema)
}

func CreateHandler(ctx context.Context) http.Handler {
	builder := schemabuilder.NewSchema()

	locationService := inject.GetLocationServiceFromContext(ctx)
	logger := inject.GetLoggerFromContext(ctx)

	providers := []graphQLProvider{
		&locationQueryProvider{
			service: locationService,
			logger:  logger.WithField("query", "location"),
		},
	}

	for _, provider := range providers {
		provider.provide(builder)
	}

	schema, err := builder.Build()
	if err != nil {
		panic(err)
	}

	introspection.AddIntrospectionToSchema(schema)

	return graphql.HTTPHandler(schema)
}
