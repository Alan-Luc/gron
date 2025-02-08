package job

import "time"

type Job struct {
	Name     string
	Interval time.Duration
	NextRun  time.Time
	Task     func()
}
