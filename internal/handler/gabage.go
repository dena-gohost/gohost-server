package handler

import (
	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/internal/handler/mock"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) GetRegister(ec echo.Context) error {
	ret := &api.GetRegisterResponse{
		Genders:      &mock.Genders,
		Universities: &mock.Universities,
	}
	return ec.JSON(http.StatusOK, &ret)
}
