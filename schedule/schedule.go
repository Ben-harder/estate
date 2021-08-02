package schedule

import (
	"container/list"
	"strings"
	"time"

	"github.com/Ben-harder/estate/household"
)

func NewSchedule(household household.HouseholdInterface) ScheduleInterface {
	schedule := &schedule{household: household, whoseTurn: household.First()}
	return schedule
}

type ScheduleInterface interface {
	NextJob() (string, string, string)
	SetTurn(name string)
	CheckNextJob()
	nextJob() *job
}

type schedule struct {
	jobs      *list.List
	household household.HouseholdInterface
	whoseTurn string
}

func (sch *schedule) NextJob() (responsibilities string, date string, whoseTurn string) {
	job := sch.jobs.Front().Value.(*job)
	return string(job.responsibilities), strings.Split(job.date.String(), " ")[0], sch.whoseTurn
}

func (sch *schedule) SetTurn(name string) {
	sch.whoseTurn = name
}

func (sch *schedule) nextJob() *job {
	job := sch.jobs.Front().Value.(*job)
	return job
}

// Removes next job if it's a day passed
func (sch *schedule) CheckNextJob() {
	now := time.Now()
	nowDay := now.YearDay()
	nextJob := sch.nextJob()
	nextJobDay := nextJob.date.YearDay()
	if nowDay > nextJobDay {
		sch.jobs.Remove(sch.jobs.Front())
		sch.household.Next(sch.whoseTurn)
	}
}

func (sch *schedule) deletePassedJobs() {
	oldJobsExist := true
	now := time.Now()
	nowDay := now.YearDay()
	for oldJobsExist {
		nextJob := sch.nextJob()
		nextJobDay := nextJob.date.YearDay()
		if nowDay > nextJobDay {
			sch.jobs.Remove(sch.jobs.Front())
		} else {
			oldJobsExist = false
		}
	}
}
