package jobs

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type Job struct {
	CronExpression string
	Cron           func()
}

type Cron struct {
	Conn *cron.Cron
	Mu   sync.Mutex
	Jobs []*Job
}
