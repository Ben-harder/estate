package schedule

type ScheduleInterface interface {
	NextJob() (string, string)
}

type schedule struct {
	jobs []*job
}

func (sch *schedule) NextJob() (responsibilities string, date string) {
	return string(sch.jobs[0].responsibilities), sch.jobs[0].date.String()
}

func (sch *schedule) NextJobResponsibilities() (responsibilities string) {
	return string(sch.jobs[0].responsibilities)
}
