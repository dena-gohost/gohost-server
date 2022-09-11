package worker

import (
	"context"
	"database/sql"
	"github.com/dena-gohost/gohost-server/pkg/logger"
	"time"

	"github.com/dena-gohost/gohost-server/internal/service"
)

const interval = 1 * time.Second

type MatchingWorker struct {
	logger *logger.Logger
	db     *sql.DB
}

func NewMatchingWorker(logger *logger.Logger, db *sql.DB) *MatchingWorker {
	return &MatchingWorker{
		logger: logger,
		db:     db,
	}
}

func (w *MatchingWorker) Run() {
	ticker := time.NewTicker(time.Second * 100)
	defer ticker.Stop()
	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			if err := w.run(); err != nil {
				panic(err)
			}
		}
	}
}

func (w *MatchingWorker) run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	txn, err := w.db.Begin()
	defer txn.Rollback()
	if err != nil {
		return err
	}

	if err := service.Match(ctx, txn); err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}
