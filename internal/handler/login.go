package handler

import (
	"context"
	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/pkg/echoutil"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) PostLogin(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &api.User{}
	if err := ec.Bind(&req); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	_ = ctx

	msg := "successfully logged in"

	return ec.JSON(http.StatusOK, &msg)
}
