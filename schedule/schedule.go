package schedule

import (
	"container/list"
	"time"

	"github.com/Ben-harder/estate/household"
)

func NewSchedule(household household.HouseholdInterface) ScheduleInterface {
	schedule := &schedule{household: household}
	return schedule
}

type ScheduleInterface interface {
	NextJob() (string, string)
	checkNextJob()
	nextJob() *job
}

type schedule struct {
	jobs      *list.List
	household household.HouseholdInterface
	whoseTurn string
}

func (sch *schedule) NextJob() (responsibilities string, date string) {
	job := sch.jobs.Front().Value.(*job)
	return string(job.responsibilities), job.date.String()
}

func (sch *schedule) SetTurn(name string) {

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
