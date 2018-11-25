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
}
