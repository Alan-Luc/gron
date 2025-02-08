package scheduler

import (
	"sync"
	"time"

	"github.com/Alan-Luc/gron/job"
)

type Scheduler struct {
	Mu   sync.Mutex
	Jobs []*job.Job
}

func (s *Scheduler) AddJob(job *job.Job) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	currentTime := time.Now()
	job.NextRun = currentTime.Add(job.Interval)
	s.Jobs = append(s.Jobs, job)
}

func (s *Scheduler) PopNextJob() *job.Job {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	if len(s.Jobs) > 0 {
		nextJob := s.Jobs[0]
		s.Jobs = s.Jobs[1:]
		return nextJob
	}
	return nil
}

func (s *Scheduler) IsNextRun() bool {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	if len(s.Jobs) == 0 {
		return false
	}
	nextJob := s.Jobs[0]
	now := time.Now()
	return now.After(nextJob.NextRun) || now.Equal(nextJob.NextRun)
}

func (s *Scheduler) RunNextJob() {
	isTimeToRunNextJob := s.IsNextRun()

	if isTimeToRunNextJob {
		currJob := s.PopNextJob()
		go currJob.Task()
		s.AddJob(currJob)
	}
}
