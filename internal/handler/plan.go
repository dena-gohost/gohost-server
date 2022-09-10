package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/internal/handler/mock"
)

func (s *Server) GetPlan(ec echo.Context) error {
	mh := time.Date(2022, 9, 11, 17, 0, 0, 0, time.Local)
	ms := "渋谷駅"
	ret := &api.GetPlanResponse{
		MeetingHour:    &mh,
		MeetingStation: &ms,
		Spot:           mock.Spots[0],
		Users:          &mock.Users,
	}
	return ec.JSON(http.StatusOK, ret)
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
