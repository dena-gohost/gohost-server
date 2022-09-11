package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/dena-gohost/gohost-server/pkg/logger"
	"github.com/dena-gohost/gohost-server/worker"

	environment "github.com/dena-gohost/gohost-server/internal/env"
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

	l := logger.New()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName))
	if err != nil {
		return err
	}

	w := worker.NewMatchingWorker(&l, db)
	w.Run()

	return nil
}
