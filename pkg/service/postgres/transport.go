package postgres

import (
	"context"
	"github.com/sirupsen/logrus"
	"gitlab.com/404busters/inventory-management/apiserver/pkg/core"
	"gitlab.com/ysitd-cloud/golang-packages/dbutils"
)

var _ core.TransportService = &TransportService{}

type TransportService struct {
	Connector *dbutils.Connector
	Logger    logrus.FieldLogger
}

func (s TransportService) List(ctx context.Context) ([]core.Transport, error) {
	conn, err := s.Connector.Connect(ctx)
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT id, person_in_charge, location, event_type, notes FROM transport WHERE deleted_at IS NULL")
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	var transportList = make([]core.Transport, 0)
	for rows.Next() {
		var transport core.Transport
		if err := rows.Scan(&transport.Id, &transport.PersonInCharge, &transport.Location, &transport.EventType, &transport.Note); err != nil {
			s.Logger.Error(err)
			return nil, err
		}
		transportList = append(transportList, transport)
	}

	if err := rows.Err(); err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return transportList, nil
}

func (s TransportService) FilterList(ctx context.Context, filter struct {
	User      string
	Inventory string
}) ([]core.Transport, error) {
	panic("implement me")
}

func (s TransportService) CheckIn(ctx context.Context, input *core.Transport) (*core.Transport, error) {
	panic("implement me")
}

func (s TransportService) CheckOut(ctx context.Context, input *core.Transport) (*core.Transport, error) {
	panic("implement me")
}
