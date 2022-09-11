package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"net/http"
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

	go startHealthCheck(env.Port)

	w := worker.NewMatchingWorker(&l, db)
	w.Run()

	return nil
}

func startHealthCheck(port int) error {
	http.Handle("/healthcheck", http.HandlerFunc(healthCheckHandler))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return err
	}

	return nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"status": ok}`)
}
