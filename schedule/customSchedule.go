package schedule

import (
	"container/list"
	"log"
	"time"
)

func NewCustomSchedule(name string, start time.Time, months int, interval int, responsibilities []string) ScheduleInterface {
	log.Printf("creating custom schedule starting on %v to last %v months at interval %v days with responsibilities %v", start, months, interval, responsibilities)

	sch := &customSchedule{start: start, months: months, interval: interval, responsibilities: responsibilities}
	sch.name = name
	sch.jobs = list.New()
	sch.populateJobs()
	sch.deletePassedJobs()
	return sch
}

type CustomScheduleInterface interface {
}

type customSchedule struct {
	schedule
	start            time.Time
	months           int
	interval         int
	responsibilities []string
}

// populateJobs uses populates the job list at the interval provided until the <months> months into the future
func (sch *customSchedule) populateJobs() {
	endTime := sch.start.AddDate(0, sch.months, 0)
	currTime := sch.start
	var i int = 0
	for currTime.Before(endTime) {
		job := &job{responsibilities: jobType(sch.responsibilities[i%len(sch.responsibilities)]), date: currTime}
		sch.jobs.PushBack(job)
		currTime = currTime.AddDate(0, 0, sch.interval)
	}
}
