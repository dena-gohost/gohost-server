package service

import (
	"context"
	"database/sql"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/gen/daocore"
)

func ListUniversities(
	ctx context.Context,
	txn *sql.Tx,
) (*[]api.University, error) {
	us, err := daocore.SelectAllUniversity(ctx, txn)
	if err != nil {
		return nil, err
	}
	res := make([]api.University, 0, len(us))
	for _, u := range us {
		res = append(res, api.University{
			Id:             &u.ID,
			MeetingStation: &u.MeetingStation,
			Name:           &u.Name,
		})
	}
	return &res, nil
}
