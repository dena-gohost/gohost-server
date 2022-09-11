package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/internal/handler/middleware"
	"github.com/dena-gohost/gohost-server/internal/handler/mock"
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
	msg := "successfully canceled"
	ret := &api.Message{Message: &msg}
	return ec.JSON(http.StatusOK, ret)
}

func (s *Server) GetPlanFinish(ec echo.Context) error {
	return ec.JSON(http.StatusOK, &mock.Users)
}

func (s *Server) PostPlanFinish(ec echo.Context) error {
	msg := "glad you've returned safely!"
	ret := &api.Message{Message: &msg}
	return ec.JSON(http.StatusOK, ret)
}
