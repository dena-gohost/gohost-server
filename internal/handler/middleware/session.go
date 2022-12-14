package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/dena-gohost/gohost-server/gen/daocore"

	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
)

const (
	SessionStoreKey     = "session"
	SessionIDKey        = "id"
	SessionExpiresAtKey = "expires_at"
	SessionDuration     = 60 * time.Minute
	ContextUserKey      = "user"
)

func SetCookie(ec echo.Context, id string) (echo.Context, error) {
	sess, _ := session.Get(SessionStoreKey, ec)
	sess.Values[SessionIDKey] = id
	sess.Values[SessionExpiresAtKey] = time.Now().Add(SessionDuration).Unix()

	// sessionに保存
	if err := sess.Save(ec.Request(), ec.Response()); err != nil {
		return ec, err
	}
	return ec, nil
}

func SessionMiddleware(db *sql.DB, ignorePaths []string) func(next echo.HandlerFunc) echo.HandlerFunc {
	ignores := map[string]bool{}
	for _, path := range ignorePaths {
		ignores[path] = true
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			// 認証が必要ないパス
			if ignores[ec.Path()] {
				return next(ec)
			}
			// sessionの取得
			sess, _ := session.Get(SessionStoreKey, ec)
			if sess == nil {
				return ec.NoContent(http.StatusUnauthorized)
			}
			// sessionからIDの取得
			id, ok := sess.Values[SessionIDKey].(string)
			if !ok {
				return ec.NoContent(http.StatusUnauthorized)
			}
			// sessionから期限の取得
			exp, ok := sess.Values[SessionExpiresAtKey].(int64)
			if !ok {
				return ec.NoContent(http.StatusUnauthorized)
			}
			// sessionの期限の検証
			if time.Now().Unix() > exp {
				return ec.NoContent(http.StatusUnauthorized)
			}
			// sessionの期限の延長
			sess.Values[SessionExpiresAtKey] = time.Unix(exp, 0).Add(SessionDuration).Unix()
			if err := sess.Save(ec.Request(), ec.Response()); err != nil {
				return err
			}

			// DB のトランザクション準備
			ctx := context.Background()
			txn, err := db.Begin()
			if err != nil {
				return ec.NoContent(http.StatusInternalServerError)
			}
			defer txn.Rollback()

			// session idからuser idの取得
			s, err := daocore.SelectOneUserSessionByID(ctx, txn, &id)
			if err != nil {
				return ec.NoContent(http.StatusUnauthorized)
			}
			// user idからuserの取得
			u, err := daocore.SelectOneUserByID(ctx, txn, &s.UserID)
			if err != nil {
				return ec.NoContent(http.StatusUnauthorized)
			}

			// contextにuser情報を付与
			ec.Set(ContextUserKey, u)

			return next(ec)
		}
	}
}

func GetUserFromSession(ec echo.Context) (*daocore.User, error) {
	user, ok := ec.Get(ContextUserKey).(daocore.User)
	if !ok {
		return nil, errors.New("cannot dump context into user struct")
	}
	return &user, nil
}
