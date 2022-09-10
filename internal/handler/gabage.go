package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/internal/handler/mock"
)

func (s *Server) GetRegister(ec echo.Context) error {
	ret := &api.GetRegisterResponse{
		Genders:      &mock.Genders,
		Universities: &mock.Universities,
	}
	return ec.JSON(http.StatusOK, &ret)
}
