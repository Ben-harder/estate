package schedule

import (
	"container/list"
	"log"
	"strings"
	"time"

	"github.com/Ben-harder/estate/household"
)

func NewSchedule(name string, household household.HouseholdInterface) ScheduleInterface {
	schedule := &schedule{name: name, household: household}
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
	name      string
	jobs      *list.List
	household household.HouseholdInterface
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
	nowDay := now.YearDay()
	nextJob := sch.nextJob()
	nextJobDay := nextJob.date.YearDay()
	if nowDay > nextJobDay {
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
