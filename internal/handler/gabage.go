package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/dena-gohost/gohost-server/gen/api"
)

func (s *Server) GetPlans(ctx echo.Context, params api.GetPlansParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetPlansPlanId(ctx echo.Context, planId string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetRegister(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetSpots(ctx echo.Context, params api.GetSpotsParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetSpotsSpotId(ctx echo.Context, spotId int) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostSpotsSpotIdEntry(ctx echo.Context, spotId string) error {
	//TODO implement me
	panic("implement me")
}
