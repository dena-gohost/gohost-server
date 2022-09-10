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

const UserTableName = "users"

var UserAllColumns = []string{
	"id",
	"first_name",
	"last_name",
	"user_name",
	"email",
	"password",
	"university_id",
	"birth_date",
	"year",
	"gender_id",
	"icon_url",
	"instagram_id",
	"good",
	"bad",
}

var UserColumnsWOMagics = []string{
	"id",
	"first_name",
	"last_name",
	"user_name",
	"email",
	"password",
	"university_id",
	"birth_date",
	"year",
	"gender_id",
	"icon_url",
	"instagram_id",
	"good",
	"bad",
}

var UserPrimaryKeyColumns = []string{
	"id",
}

type User struct {
	ID           string
	FirstName    string
	LastName     string
	UserName     string
	Email        string
	Password     string
	UniversityID string
	BirthDate    *time.Time
	Year         int
	GenderID     string
	IconUrl      string
	InstagramID  string
	Good         int
	Bad          int
}

func (t *User) Values() []interface{} {
	return []interface{}{
		t.ID,
		t.FirstName,
		t.LastName,
		t.UserName,
		t.Email,
		t.Password,
		t.UniversityID,
		t.BirthDate,
		t.Year,
		t.GenderID,
		t.IconUrl,
		t.InstagramID,
		t.Good,
		t.Bad,
	}
}

func (t *User) SetMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            t.ID,
		"first_name":    t.FirstName,
		"last_name":     t.LastName,
		"user_name":     t.UserName,
		"email":         t.Email,
		"password":      t.Password,
		"university_id": t.UniversityID,
		"birth_date":    t.BirthDate,
		"year":          t.Year,
		"gender_id":     t.GenderID,
		"icon_url":      t.IconUrl,
		"instagram_id":  t.InstagramID,
		"good":          t.Good,
		"bad":           t.Bad,
	}
}

func (t *User) Ptrs() []interface{} {
	return []interface{}{
		&t.ID,
		&t.FirstName,
		&t.LastName,
		&t.UserName,
		&t.Email,
		&t.Password,
		&t.UniversityID,
		&t.BirthDate,
		&t.Year,
		&t.GenderID,
		&t.IconUrl,
		&t.InstagramID,
		&t.Good,
		&t.Bad,
	}
}

func IterateUser(sc interface{ Scan(...interface{}) error }) (User, error) {
	t := User{}
	if err := sc.Scan(t.Ptrs()...); err != nil {
		return User{}, dberror.MapError(err)
	}
	return t, nil
}

func SelectAllUser(ctx context.Context, txn *sql.Tx) ([]*User, error) {
	query, params, err := squirrel.
		Select(UserAllColumns...).
		From(UserTableName).
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
	res := make([]*User, 0)
	for rows.Next() {
		t, err := IterateUser(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}

func SelectOneUserByEmail(ctx context.Context, txn *sql.Tx, email *string) (User, error) {
	eq := squirrel.Eq{}
	if email != nil {
		eq["email"] = *email
	}
	query, params, err := squirrel.
		Select(UserAllColumns...).
		From(UserTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return User{}, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return User{}, dberror.MapError(err)
	}
	return IterateUser(stmt.QueryRowContext(ctx, params...))
}

func SelectUserByUniversityIDAndBirthDateAndYearAndGenderID(ctx context.Context, txn *sql.Tx, university_id *string, birth_date **time.Time, year *int, gender_id *string) ([]*User, error) {
	eq := squirrel.Eq{}
	if university_id != nil {
		eq["university_id"] = *university_id
	}
	if birth_date != nil {
		eq["birth_date"] = *birth_date
	}
	if year != nil {
		eq["year"] = *year
	}
	if gender_id != nil {
		eq["gender_id"] = *gender_id
	}
	query, params, err := squirrel.
		Select(UserAllColumns...).
		From(UserTableName).
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
	res := make([]*User, 0)
	for rows.Next() {
		t, err := IterateUser(rows)
		if err != nil {
			return nil, dberror.MapError(err)
		}
		res = append(res, &t)
	}
	return res, nil
}

func SelectOneUserByUserName(ctx context.Context, txn *sql.Tx, user_name *string) (User, error) {
	eq := squirrel.Eq{}
	if user_name != nil {
		eq["user_name"] = *user_name
	}
	query, params, err := squirrel.
		Select(UserAllColumns...).
		From(UserTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return User{}, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return User{}, dberror.MapError(err)
	}
	return IterateUser(stmt.QueryRowContext(ctx, params...))
}

func SelectOneUserByID(ctx context.Context, txn *sql.Tx, id *string) (User, error) {
	eq := squirrel.Eq{}
	if id != nil {
		eq["id"] = *id
	}
	query, params, err := squirrel.
		Select(UserAllColumns...).
		From(UserTableName).
		Where(eq).
		ToSql()
	if err != nil {
		return User{}, dberror.MapError(err)
	}
	stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
		return User{}, dberror.MapError(err)
	}
	return IterateUser(stmt.QueryRowContext(ctx, params...))
}

func InsertUser(ctx context.Context, txn *sql.Tx, records []*User) error {
	for i := range records {
		if records[i] == nil {
			records = append(records[:i], records[i+1:]...)
		}
	}
	if len(records) == 0 {
		return nil
	}
	sq := squirrel.Insert(UserTableName).Columns(UserColumnsWOMagics...)
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

func UpdateUser(ctx context.Context, txn *sql.Tx, record User) error {
	sql, params, err := squirrel.Update(UserTableName).SetMap(record.SetMap()).
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

func UpsertUser(ctx context.Context, txn *sql.Tx, record User) error {
	updateSQL, params, err := squirrel.Update(UserTableName).SetMap(record.SetMap()).ToSql()
	if err != nil {
		return err
	}
	updateSQL = strings.TrimPrefix(updateSQL, "UPDATE "+UserTableName+" SET ")
	query, params, err := squirrel.Insert(UserTableName).Columns(UserColumnsWOMagics...).Values(record.Values()...).SuffixExpr(squirrel.Expr("ON DUPLICATE KEY UPDATE "+updateSQL, params...)).ToSql()
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

func DeleteOneUserByEmail(ctx context.Context, txn *sql.Tx, email *string) error {
	eq := squirrel.Eq{}
	if email != nil {
		eq["email"] = *email
	}

	query, params, err := squirrel.
		Delete(UserTableName).
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

func DeleteUserByUniversityIDAndBirthDateAndYearAndGenderID(ctx context.Context, txn *sql.Tx, university_id *string, birth_date **time.Time, year *int, gender_id *string) error {
	eq := squirrel.Eq{}
	if university_id != nil {
		eq["university_id"] = *university_id
	}
	if birth_date != nil {
		eq["birth_date"] = *birth_date
	}
	if year != nil {
		eq["year"] = *year
	}
	if gender_id != nil {
		eq["gender_id"] = *gender_id
	}

	query, params, err := squirrel.
		Delete(UserTableName).
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

func DeleteOneUserByUserName(ctx context.Context, txn *sql.Tx, user_name *string) error {
	eq := squirrel.Eq{}
	if user_name != nil {
		eq["user_name"] = *user_name
	}

	query, params, err := squirrel.
		Delete(UserTableName).
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

func DeleteOneUserByID(ctx context.Context, txn *sql.Tx, id *string) error {
	eq := squirrel.Eq{}
	if id != nil {
		eq["id"] = *id
	}

	query, params, err := squirrel.
		Delete(UserTableName).
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
