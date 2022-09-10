package handler

import (
	"context"
	"net/http"

	"github.com/dena-gohost/gohost-server/internal/handler/middleware"
	"github.com/dena-gohost/gohost-server/internal/service"
	"github.com/dena-gohost/gohost-server/pkg/echoutil"

	"github.com/labstack/echo/v4"
)

func (s *Server) PostLogout(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user, err := middleware.GetUserFromSession(ec)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	msg, err := service.Logout(ctx, txn, user)
	if msg == nil || err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	return ec.JSON(http.StatusOK, msg)
}
