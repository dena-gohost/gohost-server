package service

import (
	"context"
	"database/sql"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"

	"github.com/dena-gohost/gohost-server/gen/api"
	"github.com/dena-gohost/gohost-server/gen/daocore"
)

var (
	defaultMeetingHour = time.Date(2000, 1, 1, 19, 0, 0, 0, time.Local)
)

func daterev(t time.Time) *openapi_types.Date {
	return &openapi_types.Date{t}
}

func GetPlan(
	ctx context.Context,
	txn *sql.Tx,
	user *daocore.User,
) (*api.GetPlanResponse, error) {
	userPlan, err := daocore.SelectOneUserPlanByUserID(ctx, txn, &user.ID)

	if err != nil {
		return nil, err
	}
	plan, err := daocore.SelectOnePlanByID(ctx, txn, &userPlan.PlanID)
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
	userPlans, err := daocore.SelectUserPlanByPlanID(ctx, txn, &plan.ID)
	if err != nil {
		return nil, err
	}
	users := make([]api.User, 0, len(userPlans))
	for _, up := range userPlans {
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
			LastName:     &user.LastName,
			UniversityId: &user.UniversityID,
			UserName:     &user.UserName,
			Year:         &user.Year,
		})
	}
	meetingHour := time.Date(plan.Date.Year(), plan.Date.Month(), plan.Date.Day(), defaultMeetingHour.Hour(), defaultMeetingHour.Minute(), defaultMeetingHour.Second(), defaultMeetingHour.Nanosecond(), time.Local)
	return &api.GetPlanResponse{
		MeetingHour:    &meetingHour,
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

const minMatchNum = 3

type spotUnivDate struct {
	spotID       string
	universityID string
	date         time.Time
}

func Match(ctx context.Context, txn *sql.Tx) error {
	entries, err := daocore.SelectAllEntry(ctx, txn)
	if err != nil {
		return err
	}

	sud2entries := map[spotUnivDate][]string{}
	for _, entry := range entries {
		e := spotUnivDate{
			spotID:       entry.SpotID,
			universityID: entry.UniversityID,
			date:         *entry.Date,
		}
		if _, ok := sud2entries[e]; !ok {
			sud2entries[e] = make([]string, 0)
		}
		sud2entries[e] = append(sud2entries[e], entry.UserID)
	}

	plans := make([]*daocore.Plan, 0)
	userPlans := make([]*daocore.UserPlan, 0)
	for sud, userIDs := range sud2entries {
		if len(userIDs) < minMatchNum {
			break
		}
		id := uuid.NewString()
		plans = append(plans, &daocore.Plan{
			ID:           id,
			SpotID:       sud.spotID,
			UniversityID: sud.universityID,
			Date:         &sud.date,
		})
		for _, userID := range userIDs {
			userPlans = append(userPlans, &daocore.UserPlan{
				UserID: userID,
				PlanID: id,
			})
		}
	}

	if err := daocore.InsertPlan(ctx, txn, plans); err != nil {
		return err
	}
	if err := daocore.InsertUserPlan(ctx, txn, userPlans); err != nil {
		return err
	}

	return nil
}
