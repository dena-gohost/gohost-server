package handler

import (
	"context"
	"net/http"

	"github.com/dena-gohost/gohost-server/internal/handler/middleware"
	"github.com/dena-gohost/gohost-server/internal/handler/validator"
	"github.com/dena-gohost/gohost-server/internal/service"

	"github.com/labstack/echo/v4"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/pkg/echoutil"
)

func (s *Server) PostLogin(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &api.User{}
	if err := ec.Bind(&req); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	if err := (&validator.User{req}).Login(); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	msg, id, err := service.Login(ctx, txn, req)
	if msg == nil || err != nil {
		return echoutil.ErrBadPassword(ec, err)
	}

	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	ec, err = middleware.SetCookie(ec, id)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	return ec.JSON(http.StatusOK, msg)
}
