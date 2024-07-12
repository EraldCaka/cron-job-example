package jobs

import (
	"github.com/robfig/cron/v3"
)

func ConnectionSkeleton() *Cron {
	return &Cron{
		Conn: cron.New(),
		Jobs: make([]*Job, 0),
	}
}
