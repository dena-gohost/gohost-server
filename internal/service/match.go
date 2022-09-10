package service

import (
	"context"
	"database/sql"
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/gen/daocore"
	"time"
)

type listSpotsOptions struct {
	Date  *time.Time
	Limit int
}

type SetupListSpotsOptions func(options *listSpotsOptions)

func WithListSpotsDate(t types.Date) SetupListSpotsOptions {
	return func(o *listSpotsOptions) {
		o.Date = &t.Time
	}
}

func WithListSpotsLimit(l int) SetupListSpotsOptions {
	return func(o *listSpotsOptions) {
		o.Limit = l
	}
}

func ListSpots(
	ctx context.Context,
	txn *sql.Tx,
	user *daocore.User,
	setupOpts ...SetupListSpotsOptions,
) ([]*api.Spot, error) {
	return []*api.Spot{}, nil
}

func GetSpot(
	ctx context.Context,
	txn *sql.Tx,
	spotID string,
) (*api.Spot, error) {
	return &api.Spot{}, nil
}

type entrySpotOptions struct {
	Date *time.Time
}

type SetupEntrySpotOptions func(options *entrySpotOptions)

func WithEntrySpotDate(t *types.Date) SetupEntrySpotOptions {
	return func(o *entrySpotOptions) {
		o.Date = &t.Time
	}
}

func EntrySpot(
	ctx context.Context,
	txn *sql.Tx,
	user *daocore.User,
	setupOpts ...SetupEntrySpotOptions,
) (*api.Message, error) {
	return &api.Message{}, nil
}
