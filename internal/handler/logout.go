package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) PostLogout(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_ = ctx

	msg := "successfully logged out"

	return ec.JSON(http.StatusOK, &msg)
}

