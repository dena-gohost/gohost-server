package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dena-gohost/gohost-server/gen/api"
	environment "github.com/dena-gohost/gohost-server/internal/env"
	"github.com/dena-gohost/gohost-server/internal/handler"
	httpmiddleware "github.com/dena-gohost/gohost-server/internal/handler/middleware"
	"github.com/dena-gohost/gohost-server/pkg/echoutil"
	"github.com/dena-gohost/gohost-server/pkg/logger"
)

func main() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run server. err:%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	env, err := environment.Process()
	if err != nil {
		return err
	}

	e := echo.New()
	e.HideBanner = true

	l := logger.New()

	// db
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName))
	if err != nil {
		return err
	}

	ss, err := base64.StdEncoding.DecodeString(env.SessionSecret)
	if err != nil {
		return err
	}
	sessStore := sessions.NewCookieStore(ss)
	sessStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   env.SessionMaxAge,
		Secure:   !env.SessionCookieInsecure,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	ignorePaths := []string{"/login", "/register"}
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.RequestID(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
		}),
		session.Middleware(sessStore),
		httpmiddleware.SessionMiddleware(db, ignorePaths),
		middleware.BodyDump(func(ec echo.Context, req, res []byte) {
			if ec.Response().Status < 400 {
				return
			}
			var reqj, resj interface{}
			json.Unmarshal(req, &reqj)
			json.Unmarshal(res, &resj)
			l.Infow("",
				"id", echoutil.RequestID(ec),
				"request", reqj,
				"response", resj,
			)
		}),
	)

	server := handler.NewServer(db)
	api.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", env.Port)))
	return nil
}
