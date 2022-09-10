package handler

import (
	"context"
	"net/http"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/internal/handler/validator"
	"github.com/dena-gohost/gohost-server/internal/service"
	"github.com/dena-gohost/gohost-server/pkg/echoutil"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetRegister(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	gs, err := service.ListGender(ctx, txn)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	us, err := service.ListUniversities(ctx, txn)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	ret := &api.GetRegisterResponse{
		Genders:      gs,
		Universities: us,
	}
	return ec.JSON(http.StatusOK, &ret)
}

func (s *Server) PostRegister(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &api.User{}
	if err := ec.Bind(req); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	if err := (&validator.User{req}).Register(); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	msg, err := service.Register(ctx, txn, req)
	if msg == nil || err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	return ec.JSON(http.StatusOK, msg)
}
