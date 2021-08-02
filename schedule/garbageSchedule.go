package schedule

func NewGarbageSchedule(pathToSchedule string) (ScheduleInterface, error) {
	sch := &schedule{}
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
	var err error
	sch.jobs, err = parseICS("schedule/schedule.ics")
	if err != nil {
		return err
	}
	return nil
}
