package cron

import (
	"fmt"
	"log"

	"github.com/mattg1243/willow-server/db"
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	Func func()
	Time string
}

func StartCronJobs(q *db.Queries) error {
	c := cron.New()
	dbCleaner := New(q)

	var jobs = []CronJob{
		{
			Time: "0 0 * * 0",
			Func: dbCleaner.DeleteExpiredResetPasswordTokens,
		},
	}

	for _, job := range jobs {
		_, err := c.AddFunc(job.Time, job.Func)
		if err != nil {
			return fmt.Errorf("failed to start cron job: %s", job.Time)
		}
	}

	c.Start()
	log.Println("Cron scheduler started")
	return nil
}
