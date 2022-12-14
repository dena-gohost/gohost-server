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

const SpotUniversityTableName = "spot_universities"

var SpotUniversityAllColumns = []string{
	"id",
	"spot_id",
	"university_id",
	"created_at",
	"updated_at",
}

var SpotUniversityColumnsWOMagics = []string{
	"id",
	"spot_id",
	"university_id",
}

var SpotUniversityPrimaryKeyColumns = []string{
	"id",
}

type SpotUniversity struct {
	ID           int
	SpotID       string
	UniversityID string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (t *SpotUniversity) Values() []interface{} {
	return []interface{}{
		t.ID,
		t.SpotID,
		t.UniversityID,
	}
}

func (t *SpotUniversity) SetMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            t.ID,
		"spot_id":       t.SpotID,
		"university_id": t.UniversityID,
	}
}

func (t *SpotUniversity) Ptrs() []interface{} {
	return []interface{}{
		&t.ID,
		&t.SpotID,
		&t.UniversityID,
		&t.CreatedAt,
		&t.UpdatedAt,
	}
}

func IterateSpotUniversity(sc interface{ Scan(...interface{}) error }) (SpotUniversity, error) {
	t := SpotUniversity{}
	if err := sc.Scan(t.Ptrs()...); err != nil {
		return SpotUniversity{}, dberror.MapError(err)
	}
	return t, nil
}

func SelectAllSpotUniversity(ctx context.Context, txn *sql.Tx) ([]*SpotUniversity, error) {
	query, params, err := squirrel.
		Select(SpotUniversityAllColumns...).
		From(SpotUniversityTableName).
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
	res := make([]*SpotUniversity, 0)
	for rows.Next() {
		t, err := IterateSpotUniversity(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}

func SelectOneSpotUniversityBySpotIDAndUniversityID(ctx context.Context, txn *sql.Tx, spot_id *string, university_id *string) (SpotUniversity, error) {
	eq := squirrel.Eq{}
	if spot_id != nil {
		eq["spot_id"] = *spot_id
	}
	if university_id != nil {
		eq["university_id"] = *university_id
	}
	query, params, err := squirrel.
		Select(SpotUniversityAllColumns...).
		From(SpotUniversityTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return SpotUniversity{}, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return SpotUniversity{}, dberror.MapError(err)
	}
	return IterateSpotUniversity(stmt.QueryRowContext(ctx, params...))
}

func SelectSpotUniversityBySpotID(ctx context.Context, txn *sql.Tx, spot_id *string) ([]*SpotUniversity, error) {
	eq := squirrel.Eq{}
	if spot_id != nil {
		eq["spot_id"] = *spot_id
	}
	query, params, err := squirrel.
		Select(SpotUniversityAllColumns...).
		From(SpotUniversityTableName).
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
	res := make([]*SpotUniversity, 0)
	for rows.Next() {
		t, err := IterateSpotUniversity(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}

func SelectSpotUniversityByUniversityID(ctx context.Context, txn *sql.Tx, university_id *string) ([]*SpotUniversity, error) {
	eq := squirrel.Eq{}
	if university_id != nil {
		eq["university_id"] = *university_id
	}
	query, params, err := squirrel.
		Select(SpotUniversityAllColumns...).
		From(SpotUniversityTableName).
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
	res := make([]*SpotUniversity, 0)
	for rows.Next() {
		t, err := IterateSpotUniversity(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}

func SelectOneSpotUniversityByID(ctx context.Context, txn *sql.Tx, id *int) (SpotUniversity, error) {
	eq := squirrel.Eq{}
	if id != nil {
		eq["id"] = *id
	}
	query, params, err := squirrel.
		Select(SpotUniversityAllColumns...).
		From(SpotUniversityTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return SpotUniversity{}, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return SpotUniversity{}, dberror.MapError(err)
	}
	return IterateSpotUniversity(stmt.QueryRowContext(ctx, params...))
}

func InsertSpotUniversity(ctx context.Context, txn *sql.Tx, records []*SpotUniversity) error {
	for i := range records {
		if records[i] == nil {
			records = append(records[:i], records[i+1:]...)
		}
	}
	if len(records) == 0 {
		return nil
	}
	sq := squirrel.Insert(SpotUniversityTableName).Columns(SpotUniversityColumnsWOMagics...)
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

func UpdateSpotUniversity(ctx context.Context, txn *sql.Tx, record SpotUniversity) error {
	sql, params, err := squirrel.Update(SpotUniversityTableName).SetMap(record.SetMap()).
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

func UpsertSpotUniversity(ctx context.Context, txn *sql.Tx, record SpotUniversity) error {
	updateSQL, params, err := squirrel.Update(SpotUniversityTableName).SetMap(record.SetMap()).ToSql()
	if err != nil {
		return err
	}
	updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+SpotUniversityTableName+" SET ")
	query, params, err := squirrel.Insert(SpotUniversityTableName).Columns(SpotUniversityColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
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

func DeleteOneSpotUniversityBySpotIDAndUniversityID(ctx context.Context, txn *sql.Tx, spot_id *string, university_id *string) error {
	eq := squirrel.Eq{}
	if spot_id != nil {
		eq["spot_id"] = *spot_id
	}
	if university_id != nil {
		eq["university_id"] = *university_id
	}

	query, params, err := squirrel.
		Delete(SpotUniversityTableName).
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

func DeleteSpotUniversityBySpotID(ctx context.Context, txn *sql.Tx, spot_id *string) error {
	eq := squirrel.Eq{}
	if spot_id != nil {
		eq["spot_id"] = *spot_id
	}

	query, params, err := squirrel.
		Delete(SpotUniversityTableName).
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

func DeleteSpotUniversityByUniversityID(ctx context.Context, txn *sql.Tx, university_id *string) error {
	eq := squirrel.Eq{}
	if university_id != nil {
		eq["university_id"] = *university_id
	}

	query, params, err := squirrel.
		Delete(SpotUniversityTableName).
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

func DeleteOneSpotUniversityByID(ctx context.Context, txn *sql.Tx, id *int) error {
	eq := squirrel.Eq{}
	if id != nil {
		eq["id"] = *id
	}

	query, params, err := squirrel.
		Delete(SpotUniversityTableName).
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
