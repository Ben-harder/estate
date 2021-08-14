package schedule

import (
	"container/list"

	"github.com/Ben-harder/estate/household"
)

func NewGarbageSchedule(name string, household household.HouseholdInterface, pathToSchedule string) (ScheduleInterface, error) {
	sch := &schedule{name: name, household: household}
	sch.jobs = list.New()
	err := sch.parseEvents(pathToSchedule)
	if err != nil {
		return nil, err
	}
	sch.deletePassedJobs()
	return sch, nil
}

type garbageSchedule struct {
	schedule
}

func (sch *schedule) parseEvents(path string) error {
	jobs, err := parseICS(path)
	if err != nil {
		return err
	}
	for _, job := range jobs {
		sch.jobs.PushBack(job)
	}
	return nil
}
