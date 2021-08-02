package schedule

import (
	"container/list"
	"time"
)

type ScheduleInterface interface {
	NextJob() (string, string)
	checkNextJob()
	nextJob() *job
}

type schedule struct {
	jobs *list.List
}

func (sch *schedule) NextJob() (responsibilities string, date string) {
	job := sch.jobs.Front().Value.(*job)
	return string(job.responsibilities), job.date.String()
}

func (sch *schedule) nextJob() *job {
	job := sch.jobs.Front().Value.(*job)
	return job
}

// Removes next job if it's a day passed
func (sch *schedule) checkNextJob() {
	now := time.Now()
	nowDay := now.YearDay()
	nextJob := sch.nextJob()
	nextJobDay := nextJob.date.YearDay()
	if nowDay > nextJobDay {
		sch.jobs.Remove(sch.jobs.Front())
	}
}
