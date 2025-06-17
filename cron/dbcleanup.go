package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/mattg1243/willow-server/db"
)

type DbCleanupCron struct {
	q *db.Queries
}

func New(q *db.Queries) *DbCleanupCron {
	return &DbCleanupCron{
		q: q,
	}
}

func (d *DbCleanupCron) DeleteExpiredResetPasswordTokens() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := d.q.DeleteExpiredResetTokens(ctx)
	if err != nil {
		fmt.Println("Error deleting expired reset password tokens: %w", err)
	}
}
