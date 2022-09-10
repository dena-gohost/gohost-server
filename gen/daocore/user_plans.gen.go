// Code generated by "script/dtogen". DO NOT EDIT.
package daocore

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/dena-gohost/gohost-server/pkg/dberror"
)

const UserPlanTableName = "user_plans"

var UserPlanAllColumns = []string{
	"id",
	"plan_id",
	"canceled",
	"finished",
	"created_at",
	"updated_at",
}

var UserPlanColumnsWOMagics = []string{
	"id",
	"plan_id",
	"canceled",
	"finished",
}

var UserPlanPrimaryKeyColumns = []string{
	"id",
}

type UserPlan struct {
	ID        string
	PlanID    string
	Canceled  *time.Time
	Finished  *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (t *UserPlan) Values() []interface{} {
	return []interface{}{
		t.ID,
		t.PlanID,
		t.Canceled,
		t.Finished,
	}
}

func (t *UserPlan) SetMap() map[string]interface{} {
	return map[string]interface{}{
		"id":       t.ID,
		"plan_id":  t.PlanID,
		"canceled": t.Canceled,
		"finished": t.Finished,
	}
}

func (t *UserPlan) Ptrs() []interface{} {
	return []interface{}{
		&t.ID,
		&t.PlanID,
		&t.Canceled,
		&t.Finished,
		&t.CreatedAt,
		&t.UpdatedAt,
	}
}

func IterateUserPlan(sc interface{ Scan(...interface{}) error }) (UserPlan, error) {
	t := UserPlan{}
	if err := sc.Scan(t.Ptrs()...); err != nil {
		return UserPlan{}, dberror.MapError(err)
	}
	return t, nil
}

func SelectAllUserPlan(ctx context.Context, txn *sql.Tx) ([]*UserPlan, error) {
	query, params, err := squirrel.
		Select(UserPlanAllColumns...).
		From(UserPlanTableName).
		ToSql()
	if err != nil {
		return nil, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return nil, dberror.MapError(err)
	}

	rows, err := stmt.QueryContext(ctx, params...)
	if err != nil {
		return nil, dberror.MapError(err)
	}
	res := make([]*UserPlan, 0)
	for rows.Next() {
		t, err := IterateUserPlan(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}

func SelectUserPlanByPlanID(ctx context.Context, txn *sql.Tx, plan_id *string) ([]*UserPlan, error) {
	eq := squirrel.Eq{}
	if plan_id != nil {
		eq["plan_id"] = *plan_id
	}
	query, params, err := squirrel.
		Select(UserPlanAllColumns...).
		From(UserPlanTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return nil, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return nil, dberror.MapError(err)
	}
	rows, err := stmt.QueryContext(ctx, params...)
	if err != nil {
		return nil, dberror.MapError(err)
	}
	res := make([]*UserPlan, 0)
	for rows.Next() {
		t, err := IterateUserPlan(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}

func SelectOneUserPlanByID(ctx context.Context, txn *sql.Tx, id *string) (UserPlan, error) {
	eq := squirrel.Eq{}
	if id != nil {
		eq["id"] = *id
	}
	query, params, err := squirrel.
		Select(UserPlanAllColumns...).
		From(UserPlanTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return UserPlan{}, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return UserPlan{}, dberror.MapError(err)
	}
	return IterateUserPlan(stmt.QueryRowContext(ctx, params...))
}

func InsertUserPlan(ctx context.Context, txn *sql.Tx, records []*UserPlan) error {
	for i := range records {
		if records[i] == nil {
			records = append(records[:i], records[i+1:]...)
		}
	}
	if len(records) == 0 {
		return nil
	}
	sq := squirrel.Insert(UserPlanTableName).Columns(UserPlanColumnsWOMagics...)
	for _, r := range records {
		if r == nil {
			continue
		}
		sq = sq.Values(r.Values()...)
	}
	query, params, err := sq.ToSql()
	if err != nil {
		return err
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return dberror.MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return dberror.MapError(err)
	}
	return nil
}

func UpdateUserPlan(ctx context.Context, txn *sql.Tx, record UserPlan) error {
	sql, params, err := squirrel.Update(UserPlanTableName).SetMap(record.SetMap()).
		Where(squirrel.Eq{
			"id": record.ID,
		}).
		ToSql()
	if err != nil {
		return err
	}
	stmt, err := txn.PrepareContext(ctx, sql)
	if err != nil {
		return dberror.MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return dberror.MapError(err)
	}
	return nil
}

func UpsertUserPlan(ctx context.Context, txn *sql.Tx, record UserPlan) error {
	updateSQL, params, err := squirrel.Update(UserPlanTableName).SetMap(record.SetMap()).ToSql()
	if err != nil {
		return err
	}
	updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+UserPlanTableName+" SET ")
	query, params, err := squirrel.Insert(UserPlanTableName).Columns(UserPlanColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
	if err != nil {
		return err
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return dberror.MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return dberror.MapError(err)
	}
	return nil
}

func DeleteUserPlanByPlanID(ctx context.Context, txn *sql.Tx, plan_id *string) error {
	eq := squirrel.Eq{}
	if plan_id != nil {
		eq["plan_id"] = *plan_id
	}

	query, params, err := squirrel.
		Delete(UserPlanTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return dberror.MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return dberror.MapError(err)
	}
	return nil
}

func DeleteOneUserPlanByID(ctx context.Context, txn *sql.Tx, id *string) error {
	eq := squirrel.Eq{}
	if id != nil {
		eq["id"] = *id
	}

	query, params, err := squirrel.
		Delete(UserPlanTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return dberror.MapError(err)
	}
	if _, err = stmt.Exec(params...); err != nil {
		return dberror.MapError(err)
	}
	return nil
}