package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/dena-gohost/gohost-server/gen/daocore"
)

const minMatchNum = 3

type spotUnivDate struct {
	spotID       string
	universityID string
	date         *time.Time
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
			date:         entry.Date,
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
			Date:         sud.date,
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
