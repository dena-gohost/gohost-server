package handler

import (
	"context"
	"net/http"

	"github.com/dena-gohost/gohost-server/internal/handler/mock"

	"github.com/labstack/echo/v4"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/internal/handler/middleware"
	"github.com/dena-gohost/gohost-server/internal/handler/validator"
	"github.com/dena-gohost/gohost-server/internal/service"
	"github.com/dena-gohost/gohost-server/pkg/echoutil"
)

func (s *Server) GetSpots(ec echo.Context, params api.GetSpotsParams) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	limit := validator.ValidateListSpotsLimit(params.Limit)

	user, err := middleware.GetUserFromSession(ec)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	spots, err := service.ListSpots(
		ctx,
		txn,
		user,
		service.WithListSpotsDate(params.Date),
		service.WithListSpotsLimit(limit))
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	_ = spots
	return ec.JSON(http.StatusOK, &mock.Spots)
}

func (s *Server) GetSpotsSpotId(ec echo.Context, spotId string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	spot, err := service.GetSpot(ctx, txn, spotId)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	_ = spot
	return ec.JSON(http.StatusOK, &mock.Spots[0])
}

func (s *Server) PostSpotsSpotIdEntry(ec echo.Context, spotId string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user, err := middleware.GetUserFromSession(ec)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	req := &api.PostSpotsSpotIdEntryJSONBody{}
	if err := ec.Bind(req); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	msg, err := service.EntrySpot(
		ctx,
		txn,
		user,
		service.WithEntrySpotDate(req.Date),
	)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	_ = msg
	ret := "エントリー完了"
	return ec.JSON(http.StatusOK, &ret)
}
