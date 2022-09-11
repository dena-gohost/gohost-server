// Code generated by "script/dtogen". DO NOT EDIT.
package daocore

import (
    "context"
    "database/sql"
    "strings"

    "github.com/Masterminds/squirrel"
    "github.com/dena-gohost/gohost-server/pkg/dberror"
)

const UniversityTableName = "universities"

var UniversityAllColumns = []string{
    "id",
    "name",
    "meeting_station",
}

var UniversityColumnsWOMagics = []string{
    "id",
    "name",
    "meeting_station",
}

var UniversityPrimaryKeyColumns = []string{
    "id",
}

type University struct {
    ID string
    Name string
    MeetingStation string
}

func (t *University) Values() []interface{} {
    return []interface{}{
        t.ID,
        t.Name,
        t.MeetingStation,
    }
}

func (t *University) SetMap() map[string]interface{} {
    return map[string]interface{}{
        "id": t.ID,
        "name": t.Name,
        "meeting_station": t.MeetingStation,
    }
}

func (t *University) Ptrs() []interface{} {
    return []interface{}{
        &t.ID,
        &t.Name,
        &t.MeetingStation,
    }
}

func IterateUniversity(sc interface{ Scan(...interface{}) error}) (University, error) {
    t := University{}
    if err := sc.Scan(t.Ptrs()...); err != nil {
        return University{}, dberror.MapError(err)
    }
    return t, nil
}

func SelectAllUniversity(ctx context.Context, txn *sql.Tx) ([]*University, error) {
    query, params, err := squirrel.
        Select(UniversityAllColumns...).
        From(UniversityTableName).
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
    res := make([]*University, 0)
    for rows.Next() {
        t, err := IterateUniversity(rows)
        if err != nil {
            return nil, dberror.MapError(err)
        }
        res = append(res, &t)
    }
    return res, nil
}

func SelectOneUniversityByName(ctx context.Context, txn *sql.Tx, name *string) (University, error) {
    eq := squirrel.Eq{}
    if name != nil {
        eq["name"] = *name
    }
    query, params, err := squirrel.
        Select(UniversityAllColumns...).
        From(UniversityTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return University{}, dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return University{}, dberror.MapError(err)
    }
    return IterateUniversity(stmt.QueryRowContext(ctx, params...))
}

func SelectOneUniversityByID(ctx context.Context, txn *sql.Tx, id *string) (University, error) {
    eq := squirrel.Eq{}
    if id != nil {
        eq["id"] = *id
    }
    query, params, err := squirrel.
        Select(UniversityAllColumns...).
        From(UniversityTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return University{}, dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return University{}, dberror.MapError(err)
    }
    return IterateUniversity(stmt.QueryRowContext(ctx, params...))
}



func InsertUniversity(ctx context.Context, txn *sql.Tx, records []*University) error {
    for i := range records {
        if records[i] == nil {
            records = append(records[:i], records[i+1:]...)
        }
    }
    if len(records) == 0 {
        return nil
    }
    sq := squirrel.Insert(UniversityTableName).Columns(UniversityColumnsWOMagics...)
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

func UpdateUniversity(ctx context.Context, txn *sql.Tx, record University) error {
    sql, params, err := squirrel.Update(UniversityTableName).SetMap(record.SetMap()).
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

func UpsertUniversity(ctx context.Context, txn *sql.Tx, record University) error {
    updateSQL, params, err := squirrel.Update(UniversityTableName).SetMap(record.SetMap()).ToSql()
    if err != nil {
        return err
    }
    updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+UniversityTableName+" SET ")
    query, params, err := squirrel.Insert(UniversityTableName).Columns(UniversityColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
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

func DeleteOneUniversityByName(ctx context.Context, txn *sql.Tx, name *string) error {
    eq := squirrel.Eq{}
    if name != nil {
        eq["name"] = *name
    }

    query, params, err := squirrel.
        Delete(UniversityTableName).
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

func DeleteOneUniversityByID(ctx context.Context, txn *sql.Tx, id *string) error {
    eq := squirrel.Eq{}
    if id != nil {
        eq["id"] = *id
    }

    query, params, err := squirrel.
        Delete(UniversityTableName).
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

