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

const GenderTableName = "genders"

var GenderAllColumns = []string{
    "id",
    "name",
}

var GenderColumnsWOMagics = []string{
    "id",
    "name",
}

var GenderPrimaryKeyColumns = []string{
    "id",
}

type Gender struct {
    ID string
    Name string
}

func (t *Gender) Values() []interface{} {
    return []interface{}{
        t.ID,
        t.Name,
    }
}

func (t *Gender) SetMap() map[string]interface{} {
    return map[string]interface{}{
        "id": t.ID,
        "name": t.Name,
    }
}

func (t *Gender) Ptrs() []interface{} {
    return []interface{}{
        &t.ID,
        &t.Name,
    }
}

func IterateGender(sc interface{ Scan(...interface{}) error}) (Gender, error) {
    t := Gender{}
    if err := sc.Scan(t.Ptrs()...); err != nil {
        return Gender{}, dberror.MapError(err)
    }
    return t, nil
}

func SelectAllGender(ctx context.Context, txn *sql.Tx) ([]*Gender, error) {
    query, params, err := squirrel.
        Select(GenderAllColumns...).
        From(GenderTableName).
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
    res := make([]*Gender, 0)
    for rows.Next() {
        t, err := IterateGender(rows)
        if err != nil {
            return nil, dberror.MapError(err)
        }
        res = append(res, &t)
    }
    return res, nil
}

func SelectOneGenderByName(ctx context.Context, txn *sql.Tx, name *string) (Gender, error) {
    eq := squirrel.Eq{}
    if name != nil {
        eq["name"] = *name
    }
    query, params, err := squirrel.
        Select(GenderAllColumns...).
        From(GenderTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return Gender{}, dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return Gender{}, dberror.MapError(err)
    }
    return IterateGender(stmt.QueryRowContext(ctx, params...))
}

func SelectOneGenderByID(ctx context.Context, txn *sql.Tx, id *string) (Gender, error) {
    eq := squirrel.Eq{}
    if id != nil {
        eq["id"] = *id
    }
    query, params, err := squirrel.
        Select(GenderAllColumns...).
        From(GenderTableName).
        Where(eq).
        ToSql()
    if err != nil {
        return Gender{}, dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return Gender{}, dberror.MapError(err)
    }
    return IterateGender(stmt.QueryRowContext(ctx, params...))
}



func InsertGender(ctx context.Context, txn *sql.Tx, records []*Gender) error {
    for i := range records {
        if records[i] == nil {
            records = append(records[:i], records[i+1:]...)
        }
    }
    if len(records) == 0 {
        return nil
    }
    sq := squirrel.Insert(GenderTableName).Columns(GenderColumnsWOMagics...)
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

func UpdateGender(ctx context.Context, txn *sql.Tx, record Gender) error {
    sql, params, err := squirrel.Update(GenderTableName).SetMap(record.SetMap()).
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

func UpsertGender(ctx context.Context, txn *sql.Tx, record Gender) error {
    updateSQL, params, err := squirrel.Update(GenderTableName).SetMap(record.SetMap()).ToSql()
    if err != nil {
        return err
    }
    updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+GenderTableName+" SET ")
    query, params, err := squirrel.Insert(GenderTableName).Columns(GenderColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
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

func DeleteOneGenderByName(ctx context.Context, txn *sql.Tx, name *string) error {
    eq := squirrel.Eq{}
    if name != nil {
        eq["name"] = *name
    }

    query, params, err := squirrel.
        Delete(GenderTableName).
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

func DeleteOneGenderByID(ctx context.Context, txn *sql.Tx, id *string) error {
    eq := squirrel.Eq{}
    if id != nil {
        eq["id"] = *id
    }

    query, params, err := squirrel.
        Delete(GenderTableName).
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

