package service

import (
	"context"
	"database/sql"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/gen/daocore"
)

func daterev(t time.Time) *openapi_types.Date {
	return &openapi_types.Date{t}
}

func GetPlan(
	ctx context.Context,
	txn *sql.Tx,
	user *daocore.User,
) (*api.GetPlanResponse, error) {
	user_plan, err := daocore.SelectOneUserPlanByUserID(ctx, txn, &user.ID)

	if err != nil {
		return nil, err
	}
	plan, err := daocore.SelectOnePlanByID(ctx, txn, &user_plan.PlanID)
	if err != nil {
		return nil, err
	}
	university, err := daocore.SelectOneUniversityByID(ctx, txn, &plan.UniversityID)
	if err != nil {
		return nil, err
	}
	spot, err := daocore.SelectOneSpotByID(ctx, txn, &plan.SpotID)
	if err != nil {
		return nil, err
	}
	user_plans, err := daocore.SelectUserPlanByPlanID(ctx, txn, &plan.ID)
	if err != nil {
		return nil, err
	}
	users := make([]api.User, 0, len(user_plans))
	for _, up := range user_plans {
		user, err := daocore.SelectOneUserByID(ctx, txn, &up.UserID)
		if err != nil {
			return nil, err
		}
		users = append(users, api.User{
			BirthDate:    daterev(*user.BirthDate),
			Email:        &user.Email,
			FirstName:    &user.FirstName,
			GenderId:     &user.GenderID,
			IconUrl:      &user.IconUrl,
			Id:           &user.ID,
			InstagramId:  &user.InstagramID,
			LastName:     &user.LastName,
			Password:     &user.Password,
			UniversityId: &user.UniversityID,
			UserName:     &user.UserName,
			Year:         &user.Year,
		})
	}

	return &api.GetPlanResponse{
		MeetingHour:    plan.Date,
		MeetingStation: &university.MeetingStation,
		Spot: &api.Spot{
			Address:     &spot.Address,
			Description: &spot.Description,
			Id:          &spot.ID,
			ImageUrl:    &spot.ImageUrl,
			Name:        &spot.Name,
		},
		Users: &users,
	}, err
}
