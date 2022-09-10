package service

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"github.com/google/uuid"

	"github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/gen/daocore"
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
) (*[]api.Spot, error) {
	opts := &listSpotsOptions{}
	for _, o := range setupOpts {
		o(opts)
	}

	su, err := daocore.SelectSpotUniversityByUniversityID(
		ctx,
		txn,
		&user.UniversityID)
	if err != nil {
		return nil, err
	}

	// 重み付けをする
	// 同じ大学&日程でエントリー者が多い場所に
	entries, err := daocore.SelectEntryByDateAndSpotIDAndUniversityID(
		ctx,
		txn,
		&opts.Date,
		nil,
		&user.UniversityID)
	if err != nil {
		return nil, err
	}

	spots2entries := map[string]int{}
	for _, s := range su {
		spots2entries[s.SpotID] = 0
	}

	for _, entry := range entries {
		spots2entries[entry.SpotID]++
	}

	// 降順に並び替え
	sort.Slice(su, func(i, j int) bool {
		ip := spots2entries[su[i].SpotID]
		jp := spots2entries[su[j].SpotID]
		return ip > jp
	})

	// FIXME : this would be N + 1
	res := make([]api.Spot, 0, opts.Limit)
	for i, s := range su {
		if i >= opts.Limit {
			break
		}
		spot, err := daocore.SelectOneSpotByID(ctx, txn, &s.SpotID)
		if err != nil {
			return nil, err
		}
		res = append(res, api.Spot{
			Address:     &spot.Address,
			Description: &spot.Description,
			Id:          &spot.ID,
			ImageUrl:    &spot.ImageUrl,
			Name:        &spot.Name,
		})
	}

	return &res, nil
}

func GetSpot(
	ctx context.Context,
	txn *sql.Tx,
	spotID string,
) (*api.Spot, error) {
	spot, err := daocore.SelectOneSpotByID(ctx, txn, &spotID)
	if err != nil {
		return nil, err
	}

	return &api.Spot{
		Address:     &spot.Address,
		Description: &spot.Description,
		Id:          &spot.ID,
		ImageUrl:    &spot.ImageUrl,
		Name:        &spot.Name,
	}, nil
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
	spotID string,
	user *daocore.User,
	setupOpts ...SetupEntrySpotOptions,
) (*api.Message, error) {
	opts := &entrySpotOptions{}
	for _, o := range setupOpts {
		o(opts)
	}

	id := uuid.NewString()

	inp := &daocore.Entry{
		ID:           id,
		UserID:       user.ID,
		UniversityID: user.UniversityID,
		Date:         opts.Date,
		SpotID:       spotID,
	}
	if err := daocore.InsertEntry(ctx, txn, []*daocore.Entry{inp}); err != nil {
		return nil, err
	}

	msg := "entry has successfully applied"

	return &api.Message{&msg}, nil
}
