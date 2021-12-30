package schedule

import (
	"container/list"
	"log"
	"strings"
	"time"
)

func NewSchedule(name string) ScheduleInterface {
	schedule := &schedule{name: name}
	return schedule
}

type ScheduleInterface interface {
	NextJob() (string, string)
	RemoveNextJob()
	Name() string
	nextJob() *job
	IsNextJobOld() bool
}

type schedule struct {
	name string
	jobs *list.List
}

func (sch *schedule) Name() string {
	return sch.name
}

// NextJob returns a schedules upcoming job's responsibilities and date.
func (sch *schedule) NextJob() (responsibilities string, date string) {
	job := sch.jobs.Front().Value.(*job)
	return string(job.responsibilities), strings.Split(job.date.String(), " ")[0]
}

func (sch *schedule) nextJob() *job {
	job := sch.jobs.Front().Value.(*job)
	return job
}

func (sch *schedule) IsNextJobOld() bool {
	now := time.Now()
	nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	nextJobTime := sch.nextJob().date
	nextJobDay := time.Date(nextJobTime.Year(), nextJobTime.Month(), nextJobTime.Day(), 0, 0, 0, 0, nextJobTime.Location())
	if nextJobDay.Before(nowDay) {
		return true
	}
	return false
}

func (sch *schedule) deletePassedJobs() {
	oldJobsExist := true
	now := time.Now()
	nowDay := now.YearDay()
	for oldJobsExist {
		nextJob := sch.nextJob()
		nextJobDay := nextJob.date.YearDay()
		if nowDay > nextJobDay {
			sch.RemoveNextJob()
		} else {
			oldJobsExist = false
		}
	}
}

func (sch *schedule) RemoveNextJob() {
	nextJob := sch.nextJob()
	log.Printf("removing job{%v, %v}", nextJob.responsibilities, nextJob.date.String())
	sch.jobs.Remove(sch.jobs.Front())
}
