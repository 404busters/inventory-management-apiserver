package http

import (
	"net/http"

	"github.com/gorilla/handlers"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/http/graphql"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/http/restful"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/logging"
)

func CreateHandler() (h http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/graphql", graphql.CreateHandler())
	mux.Handle("/", restful.CreateHandler())

	h = mux
	h = handlers.CombinedLoggingHandler(logging.GetRoot().Writer(), h)

	return
}
