package service

import (
	"context"
	"database/sql"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/gen/daocore"
)

func ListGender(
	ctx context.Context,
	txn *sql.Tx,
) (*[]api.Gender, error) {
	gs, err := daocore.SelectAllGender(ctx, txn)
	if err != nil {
		return nil, err
	}
	res := make([]api.Gender, 0, len(gs))
	for _, g := range gs {
		res = append(res, api.Gender{
			Id:   &g.ID,
			Name: &g.Name,
		})
	}
	return &res, nil
}
