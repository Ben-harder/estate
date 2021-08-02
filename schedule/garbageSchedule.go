package schedule

import "container/list"

func NewGarbageSchedule(pathToSchedule string) (ScheduleInterface, error) {
	sch := &schedule{}
	sch.jobs = list.New()
	err := sch.parseEvents(pathToSchedule)
	if err != nil {
		return nil, err
	}
	return sch, nil
}

type garbageSchedule struct {
	schedule
}

func (sch *schedule) parseEvents(path string) error {
	jobs, err := parseICS("schedule/schedule.ics")
	if err != nil {
		return err
	}
	for _, job := range jobs {
		sch.jobs.PushBack(job)
	}
	return nil
}
