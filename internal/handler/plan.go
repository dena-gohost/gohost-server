package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/internal/handler/middleware"
	"github.com/dena-gohost/gohost-server/internal/service"
	"github.com/dena-gohost/gohost-server/pkg/echoutil"
)

func (s *Server) GetPlan(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()
	user, err := middleware.GetUserFromSession(ec)
	plan, err := service.GetPlan(ctx, txn, user)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	return ec.JSON(http.StatusOK, plan)
}

func (s *Server) PostPlanCancel(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()
	user, err := middleware.GetUserFromSession(ec)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	err = service.PostPlanCancel(ctx, txn, user)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	msg := "successfully canceled"
	ret := &api.Message{Message: &msg}
	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	return ec.JSON(http.StatusOK, ret)
}

func (s *Server) GetPlanFinish(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()
	user, err := middleware.GetUserFromSession(ec)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	users, err := service.GetPlanUsers(ctx, txn, user)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	return ec.JSON(http.StatusOK, users)
}

func (s *Server) PostPlanFinish(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &api.PostPlanFinishJSONRequestBody{}
	if err := ec.Bind(req); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()
	user, err := middleware.GetUserFromSession(ec)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	err = service.PostPlanFinish(ctx, txn, user, req)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	msg := "glad you've returned safely!"
	ret := &api.Message{Message: &msg}
	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	return ec.JSON(http.StatusOK, ret)
}
