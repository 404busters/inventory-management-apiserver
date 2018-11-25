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

package http

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/http/graphql"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/http/restful"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/logging"
)

type contextBinder func(ctx context.Context) context.Context

var allowedHost = strings.Split(os.Getenv("CORS_ALLOWED_HOST"), ",")

func CreateHandler() (h http.Handler) {
	ctx := context.Background()

	binders := []contextBinder{
		bindLogger,
		bindService,
	}

	for _, binder := range binders {
		ctx = binder(ctx)
	}

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphql.CreateHandler(ctx))
	mux.Handle("/graphiql/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/graphiql/index.html")
	}))
	mux.Handle("/", restful.CreateHandler(ctx))

	h = mux
	h = handlers.CombinedLoggingHandler(logging.GetRoot().Writer(), h)
	h = handlers.CORS(
		handlers.AllowedOrigins(allowedHost),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
	)(h)

	return
}
